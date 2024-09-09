package rabbitmq

import (
	"context"
	"fmt"
	"github.com/gogf/gf/v2/frame/g"
	amqp "github.com/rabbitmq/amqp091-go"
	"net"
	"reflect"
	"sync"
	"time"
)

type AMQP struct {
	ctx         context.Context
	cfg         Config
	endpoint    string
	conn        *amqp.Connection
	closeNotify chan *amqp.Error
	m           sync.RWMutex // 读写锁
}

func NewAMQP(ctx context.Context, cfg Config) *AMQP {
	customAmqp := &AMQP{
		ctx:      ctx,
		cfg:      cfg,
		endpoint: fmt.Sprintf("amqp://%s:%s@%s:%s/%s", cfg.User, cfg.Password, cfg.Host, cfg.Port, cfg.Vhost),
	}
	return customAmqp
}

// 连接
func (a *AMQP) Conn(ctx context.Context) error {
	a.m.Lock()
	defer func() {
		a.m.Unlock()
	}()
	if a.conn != nil && !a.conn.IsClosed() {
		return nil
	}
	var (
		err    error
		config = amqp.Config{
			Dial: func(network, addr string) (net.Conn, error) {
				return net.DialTimeout(network, addr, a.cfg.MaxDialTimeout)
			},
		}
	)
	if a.conn, err = amqp.DialConfig(a.endpoint, config); err != nil {
		g.Log().Errorf(context.TODO(), "Conn fail, reason: %v", err)
	}
	return err
}

// 校验并且验证
func (a *AMQP) CheckAndGetConn(ctx context.Context, indefinitely bool, times uint) {
	if a.conn == nil {
		a.ReConn(ctx, indefinitely, times)
	} else {
		if a.conn.IsClosed() {
			a.ReConn(ctx, indefinitely, times)
		}
	}
}

// 关闭连接
func (a *AMQP) Close() error {
	a.m.Lock()
	defer func() {
		a.m.Unlock()
	}()
	return a.conn.Close()
}

// 重新连接
func (a *AMQP) ReConn(ctx context.Context, indefinitely bool, times uint) {
	//是否一直重试
	if indefinitely {
		for {
			err := a.Conn(ctx)
			if err == nil {
				break
			}
			g.Log().Errorf(context.TODO(), "connection closed, reason: %v", err)
			time.Sleep(3 * time.Second)
			continue
		}
	} else {
		for i := uint(0); i <= times; i++ {
			err := a.Conn(ctx)
			if err == nil {
				break
			}
			g.Log().Errorf(context.TODO(), "connection closed, reason: %v", err)
			time.Sleep(3 * time.Second)
		}
	}
}

// 获取channel
func (a *AMQP) Channel(ctx context.Context) (ch *amqp.Channel, err error) {
	a.CheckAndGetConn(ctx, false, 3)
	ch, err = a.conn.Channel()
	return ch, err
}

func ParsePushMessage(message interface{}) []byte {
	var (
		t   = reflect.TypeOf(message)
		msg []byte
		err error
	)
	switch t.Kind() {
	case reflect.String:
		msg = []byte(fmt.Sprintf("%s", message))
	case reflect.Map:
		fallthrough
	case reflect.Slice:
		fallthrough
	case reflect.Array:
		fallthrough
	case reflect.Struct:
		if msg, err = until.Marshal(message); err != nil {
			panic(err)
		}
	default:
		panic("不支持的消息类型！")
	}
	return msg
}
