package dao

import (
	"context"

	pb "edu/api/logic/grpc"

	"github.com/golang/protobuf/proto"
	"github.com/streadway/amqp"
)

// PushMsg push a message to databus.
func (d *dao) PushMsg(c context.Context, op int32, server string, keys []string, msg []byte) (err error) {
	pushMsg := &pb.PushMsg{
		Type:      pb.PushMsg_PUSH,
		Operation: op,
		Server:    server,
		Keys:      keys,
		Msg:       msg,
	}
	b, err := proto.Marshal(pushMsg)
	if err != nil {
		return
	}
	header := make(map[string]interface{})
	// header["x-delay"] = 0 // 设置0.1秒延时
	// header["x-delay"] = delay // 设置0.1秒延时r
	// 将消息发送到延时队列上
	err = d.channel.Publish(
		"amq.direct", // exchange 这里为空则不选择 exchange
		"im",         // routing key
		false,        // mandatory
		false,        // immediate
		amqp.Publishing{
			Headers:      header,
			ContentType:  "text/plain",
			Body:         b,
			DeliveryMode: amqp.Persistent,
		})
	if err != nil {
		d.log.Error("err = %v", err)
		return
	}
	d.log.Info(" PushWitMsg === %s", b)
	return
}

// BroadcastRoomMsg push a message to databus.
func (d *dao) BroadcastRoomMsg(c context.Context, op int32, room string, msg []byte) (err error) {
	pushMsg := &pb.PushMsg{
		Type:      pb.PushMsg_ROOM,
		Operation: op,
		Room:      room,
		Msg:       msg,
	}
	b, err := proto.Marshal(pushMsg)
	if err != nil {
		return
	}
	header := make(map[string]interface{})
	// header["x-delay"] = 0 // 设置0.1秒延时
	// header["x-delay"] = delay // 设置0.1秒延时r
	// 将消息发送到延时队列上
	err = d.channel.Publish(
		"amq.direct", // exchange 这里为空则不选择 exchange
		"im",         // routing key
		false,        // mandatory
		false,        // immediate
		amqp.Publishing{
			Headers:      header,
			ContentType:  "text/plain",
			Body:         b,
			DeliveryMode: amqp.Persistent,
		})
	if err != nil {
		d.log.Error("err = %v", err)
		d.log.Errorf("PushMsg.send(broadcast_room pushMsg:%v) error(%v)", pushMsg, err)
		return
	}
	return
}

// BroadcastMsg push a message to databus.
func (d *dao) BroadcastMsg(c context.Context, op, speed int32, msg []byte) (err error) {
	pushMsg := &pb.PushMsg{
		Type:      pb.PushMsg_BROADCAST,
		Operation: op,
		Speed:     speed,
		Msg:       msg,
	}
	b, err := proto.Marshal(pushMsg)
	if err != nil {
		return
	}
	header := make(map[string]interface{})
	// header["x-delay"] = 0 // 设置0.1秒延时
	// header["x-delay"] = delay // 设置0.1秒延时r
	// 将消息发送到延时队列上
	err = d.channel.Publish(
		"amq.direct", // exchange 这里为空则不选择 exchange
		"im",         // routing key
		false,        // mandatory
		false,        // immediate
		amqp.Publishing{
			Headers:      header,
			ContentType:  "text/plain",
			Body:         b,
			DeliveryMode: amqp.Persistent,
		})
	if err != nil {
		d.log.Error("err = %v", err)
		d.log.Errorf("PushMsg.send(broadcast pushMsg:%v) error(%v)", pushMsg, err)
		return
	}
	return
}
