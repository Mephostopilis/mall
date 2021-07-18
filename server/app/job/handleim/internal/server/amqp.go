package server

import (
	"context"
	"fmt"
	"runtime"
	"sync"
	"time"

	pb "edu/api/logic/grpc"
	"edu/job/handleim/internal/biz"
	"edu/job/handleim/internal/conf"
	"edu/pkg/tools"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/streadway/amqp"
	"google.golang.org/protobuf/proto"
)

type AmqpConsumer struct {
	conn    *amqp.Connection
	channel *amqp.Channel
	log     *log.Helper
	uc      *biz.Job
	ctx     context.Context
	cf      context.CancelFunc
	wg      sync.WaitGroup
}

func NewAmqpConsumer(c *conf.Server, logger log.Logger, uc *biz.Job) (consumer *AmqpConsumer, err error) {
	log := log.NewHelper(log.With(logger, "module", "amqp/consumer"))
	connection, err := amqp.Dial("amqp://admin:123456@127.0.0.1:5672/my_vhost")
	if err != nil {
		log.Errorf("rabbit刷新连接失败, %v", err)
		return
	}
	ch, err := connection.Channel()
	if err != nil {
		log.Errorf("rabbit刷新连接失败, %v", err)
		return
	}
	if err != nil {
		log.Errorf("Failed to new consumer: %s", err)
		return
	}
	ctx, cf := context.WithCancel(context.Background())
	consumer = &AmqpConsumer{
		conn:    connection,
		channel: ch,
		log:     log,
		ctx:     ctx,
		cf:      cf,
		uc:      uc,
	}
	return
}

func (s *AmqpConsumer) Start(context.Context) error {
	s.wg.Add(1)
	go s.run(s.ctx, "test.rabbit")
	s.log.Infof("start amqp consumer.")
	return nil
}

func (s *AmqpConsumer) Stop(context.Context) error {
	s.cf()
	s.wg.Wait()
	s.uc.Close()
	s.log.Infof("stop amqp consumer.")
	return nil
}

func (s *AmqpConsumer) run(ctx context.Context, queueName string) {
	gid, _ := tools.GetGID()
	defer func() {
		s.wg.Done()
		s.log.Infof("grab done, co(%d)", gid)
		if err := recover(); err != nil {
			const size = 64 << 10
			buf := make([]byte, size)
			buf = buf[:runtime.Stack(buf, false)]
			pl := fmt.Sprintf("grab call panic: %v\n%s\n", err, buf)
			s.log.Errorf("%s", pl)
		}
	}()

	// 获取消费通道
	s.channel.Qos(1, 0, true) // 确保rabbitmq会一个一个发消息

retry:
	s.log.Infof("channel consume queue[%s]", queueName)
	msgs, err := s.channel.Consume(
		queueName, // queue
		"",        // consumer
		false,     // auto-ack
		false,     // exclusive
		false,     // no-local
		false,     // no-wait
		nil,       // args
	)
	if nil != err {
		s.log.Errorf("获取队列[%s]的消费通道失败: %v", queueName, err)
		time.Sleep(1 * time.Second)
		goto retry
	}
	s.log.Infof("开始处理[%s]消息", queueName)

	for {
		select {
		case <-ctx.Done():
			s.log.Info("结束消费队列[%s], exit.", queueName)
			return
		case msg := <-msgs:
			// 当接收者消息处理失败的时候，
			// 比如网络问题导致的数据库连接失败，redis连接失败等等这种
			// 通过重试可以成功的操作，那么这个时候是需要重试的
			// 直到数据处理成功后再返回，然后才会回复rabbitmq ack
			// receiver.OnReceive(&msg)
			s.handleMsg(ctx, &msg)

		default:
			// log.Warn("休眠3s")
			time.Sleep(3 * time.Second)
		}
	}
}

func (s *AmqpConsumer) handleMsg(ctx context.Context, msg *amqp.Delivery) error {
	pushMsg := new(pb.PushMsg)
	if err := proto.Unmarshal(msg.Body, pushMsg); err != nil {
		s.log.Errorf("proto.Unmarshal(%v) error(%v)", msg, err)
		return err
	}
	if err := s.uc.Push(ctx, pushMsg); err != nil {
		s.log.Errorf("j.push(%v) error(%v)", pushMsg, err)
		return err
	}
	s.log.Infof("consume: %s/\t%+v", msg.Exchange, pushMsg)
	return nil
}
