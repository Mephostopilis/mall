package dao

import (
	"context"
	"time"

	"edu/service/im/internal/conf"
	"edu/service/im/internal/model"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-redis/redis/v8"
	"github.com/google/wire"
	"github.com/streadway/amqp"
)

var ProviderSet = wire.NewSet(New)

type Dao interface {
	Close()
	Ping(ctx context.Context) (err error)
	PingRedis(ctx context.Context) (err error)
	// bts: -nullcache=&model.Article{ID:-1} -check_null_code=$!=nil&&$.ID==-1
	// Article(c context.Context, id int64) (*model.Article, error)

	PushMsg(c context.Context, op int32, server string, keys []string, msg []byte) (err error)
	BroadcastRoomMsg(c context.Context, op int32, room string, msg []byte) (err error)
	BroadcastMsg(c context.Context, op, speed int32, msg []byte) (err error)

	AddMapping(c context.Context, mid int64, key, server string) (err error)
	ExpireMapping(c context.Context, mid int64, key string) (has bool, err error)
	DelMapping(c context.Context, mid int64, key, server string) (has bool, err error)
	ServersByKeys(c context.Context, keys []string) (res []string, err error)
	KeysByMids(c context.Context, mids []int64) (ress map[string]string, olMids []int64, err error)
	AddServerOnline(c context.Context, server string, online *model.Online) (err error)
	ServerOnline(c context.Context, server string) (online *model.Online, err error)
	DelServerOnline(c context.Context, server string) (err error)
}

// Dao dao.
type dao struct {
	c           *conf.Data
	conn        *amqp.Connection
	channel     *amqp.Channel
	redis       *redis.Client
	redisExpire time.Duration
	log         *log.Helper
}

// New new a dao and return.
func New(c *conf.Data, logger log.Logger) Dao {
	d, _, err := newDao(c, logger)
	if err != nil {
		panic(err)
	}
	return d
}

func newDao(c *conf.Data, logger log.Logger) (d *dao, cf func(), err error) {
	log := log.NewHelper(log.With(logger, "module", "dao"))
	conn, err := amqp.Dial("amqp://admin:123456@127.0.0.1:5672/my_vhost")
	if err != nil {
		log.Errorf("err = %v", err)
		return
	}
	ch, err := conn.Channel()
	if err != nil {
		log.Errorf("err = %v", err)
		return
	}
	re := newRedis(c.Redis)
	d = &dao{
		c:           c,
		conn:        conn,
		channel:     ch,
		redis:       re,
		redisExpire: 10 * time.Second,
	}
	return d, nil, nil
}

func newRedis(c *conf.Data_Redis) *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Addr: ":6379",
	})
	// return &redis.Pool{
	// 	MaxIdle:     time.Minute,
	// 	MaxActive:   1,
	// 	IdleTimeout: time.Duration(c.IdleTimeout),
	// 	Dial: func() (redis.Conn, error) {
	// 		conn, err := redis.Dial(c.Network, c.Addr,
	// 			redis.DialConnectTimeout(c.DialTimeout.AsDuration()),
	// 			redis.DialReadTimeout(c.ReadTimeout.AsDuration()),
	// 			redis.DialWriteTimeout(c.WriteTimeout.AsDuration()),
	// 			redis.DialPassword(c.Auth),
	// 		)
	// 		if err != nil {
	// 			return nil, err
	// 		}
	// 		return conn, nil
	// 	},
	// }
	return rdb
}

func (d *dao) PingRedis(ctx context.Context) (err error) {
	if err = d.redis.Ping(ctx).Err(); err != nil {
		d.log.Errorf("conn.Set(PING) error(%v)", err)
	}
	return
}

// Close close the resource.
func (d *dao) Close() {
	d.redis.Close()
}

// Ping dao ping.
func (d *dao) Ping(c context.Context) error {
	return d.PingRedis(c)
}
