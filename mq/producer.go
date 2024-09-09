package rabbitmq

import (
	"context"
	"github.com/gogf/gf/v2/frame/g"
	amqp "github.com/rabbitmq/amqp091-go"
	uuid "github.com/satori/go.uuid"
	"time"
)

type Producer struct {
	ProducerKey  string
	Exchange     string
	RoutingKey   string
	ProducerType string
}

func (a *AMQP) DeclareProducer(ctx context.Context, channel *amqp.Channel, p *Producer) (err error) {
	err = channel.ExchangeDeclare(p.Exchange, p.ProducerType, true, true, false, false, nil)
	if err != nil {
		return err
	}
	return nil
}

func (a *AMQP) PushWithChannel(ctx context.Context, channel *amqp.Channel, p *Producer, message any) (err error) {
	var (
		messageId = uuid.NewV4().String()
	)
	err = channel.PublishWithContext(
		ctx,
		p.Exchange,
		p.RoutingKey,
		false,
		false,
		amqp.Publishing{
			ContentType:     "application/json",
			ContentEncoding: "UTF-8",
			Timestamp:       time.Now(),
			MessageId:       messageId,
			Body:            ParsePushMessage(message),
		},
	)
	if err != nil {
		g.Log().Errorf(ctx, "PublishWithContext Err: %v", err)
		return
	}
	return
}

// confirm 模式的
func (a *AMQP) PushWithChannelAndConfirm(ctx context.Context, p *Producer, message any) (err error) {
	channel, err := a.Channel(ctx)
	if err != nil {
		g.Log().Errorf(ctx, "get channel err:", err)
		return err
	}
	confirmChan := channel.NotifyPublish(make(chan amqp.Confirmation, 1))
	if err != nil {
		g.Log().Errorf(ctx, "Producer Confirm err:", err)
		return err
	}
	err = a.PushWithChannel(ctx, channel, p, message)
	if err != nil {
		g.Log().Errorf(ctx, "Producer Confirm err:", err)
		return err
	}
	//异步处理确认
	select {
	case confirm := <-confirmChan:
		if !confirm.Ack {
			g.Log().Errorf(ctx, "Message with delivery tag %d not acknowledged", confirm.DeliveryTag)
			return err
		}
		return nil
	case <-time.After(5 * time.Second): // 设置超时时间
		g.Log().Errorf(ctx, "Waiting for confirmation timed out")
		return err
	}
}
