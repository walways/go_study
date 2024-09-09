package rabbitmq

import (
	"context"
	"fmt"
	"github.com/gogf/gf/v2/net/gtrace"
	amqp "github.com/rabbitmq/amqp091-go"
	"reflect"
	"sync"
	"time"
)

type ConsumerHandler func(ctx context.Context, message *ReceiveMsgPackage)

type Consumer struct {
	ConsumerKey string
	Exchange    string
	RoutingKey  string
	Queue       string
	AutoAck     bool
	ConsumerNum int // 消费者数量
	Handle      ConsumerHandler
}

type ReceiveMsgPackage struct {
	RequestId  string
	MsgId      string
	SendTime   time.Time
	MsgContent string
	originMsg  []byte
	resType    reflect.Type
}

func (a *AMQP) QueueDeclare(ctx context.Context, channel *amqp.Channel, consumer *Consumer) (err error) {
	_, err = channel.QueueDeclare(consumer.Queue, true, false, false, false, nil)
	return err
}

func (a *AMQP) QueueBind(ctx context.Context, channel *amqp.Channel, consumer *Consumer) (err error) {
	err = channel.QueueBind(
		consumer.Queue,
		consumer.RoutingKey,
		consumer.Exchange,
		false,
		nil,
	)
	return err
}
func (a *AMQP) Running(ctx context.Context, consumer *Consumer) (err error) {
	channel, err := a.Channel(ctx)
	if err != nil {
		return fmt.Errorf("failed to create channel: %w", err)
	}
	defer func() {
		if closeErr := channel.Close(); closeErr != nil {
			//a.logger.Errorf(ctx, "Failed to close channel: %v", closeErr)
		}
	}()

	if err := a.setupQueue(ctx, channel, consumer); err != nil {
		return err
	}

	messages, err := a.consumeMessages(ctx, channel, consumer)
	if err != nil {
		return err
	}

	return a.processMessages(ctx, messages, consumer)
}

func (a *AMQP) setupQueue(ctx context.Context, channel *amqp.Channel, consumer *Consumer) error {
	if err := a.QueueDeclare(ctx, channel, consumer); err != nil {
		return fmt.Errorf("failed to declare queue: %w", err)
	}
	if err := a.QueueBind(ctx, channel, consumer); err != nil {
		return fmt.Errorf("failed to bind queue: %w", err)
	}
	return nil
}

func (a *AMQP) consumeMessages(ctx context.Context, channel *amqp.Channel, consumer *Consumer) (<-chan amqp.Delivery, error) {
	messages, err := channel.Consume(
		consumer.Queue,
		consumer.ConsumerKey,
		consumer.AutoAck,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		return nil, fmt.Errorf("failed to consume messages: %w", err)
	}
	return messages, nil
}

func (a *AMQP) processMessages(ctx context.Context, messages <-chan amqp.Delivery, consumer *Consumer) error {
	var wg sync.WaitGroup
	wg.Add(consumer.ConsumerNum)

	for i := 0; i < consumer.ConsumerNum; i++ {
		go func() {
			defer wg.Done()
			a.messageWorker(ctx, messages, consumer)
		}()
	}

	wg.Wait()
	return nil
}

func (a *AMQP) messageWorker(ctx context.Context, messages <-chan amqp.Delivery, consumer *Consumer) {
	for {
		select {
		case message, ok := <-messages:
			if !ok {
				//a.logger.Error(ctx, "Message channel closed")
				return
			}
			a.handleMessage(ctx, message, consumer)
		case closeNotify := <-a.closeNotify:
			_ = fmt.Errorf("Connection closed, reason: %v", closeNotify.Reason)
			a.ReConn(ctx, true, 0)
			return
		case <-ctx.Done():
			_ = fmt.Errorf("Context canceled")
			return
		}
	}
}

func (a *AMQP) handleMessage(ctx context.Context, message amqp.Delivery, consumer *Consumer) {
	defer func() {
		if r := recover(); r != nil {
			//	a.logger.Errorf(ctx, "Recovered from panic while processing message: %v", r)
		}
	}()

	receive := &ReceiveMsgPackage{
		originMsg:  message.Body,
		MsgId:      message.MessageId,
		MsgContent: string(message.Body),
		SendTime:   message.Timestamp,
	}

	traceCtx, _ := gtrace.WithTraceID(ctx, receive.RequestId)
	consumer.Handle(traceCtx, receive)
}
