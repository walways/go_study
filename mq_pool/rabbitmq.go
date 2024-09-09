// connection_pool.go

package rabbitmq

import (
	"context"
	amqp "github.com/rabbitmq/amqp091-go"
	"sync"
	"time"
)

type ConnectionPool interface {
	Get(ctx context.Context) (*amqp.Connection, error)
	Put(*amqp.Connection)
	Close()
}

type AMQPConnectionPool struct {
	mu          sync.Mutex
	connections []*amqp.Connection
	maxSize     int
	currentSize int
	url         string
	maxLifetime time.Duration
	maxIdleTime time.Duration
}

func NewAMQPConnectionPool(url string, maxSize int, maxLifetime, maxIdleTime time.Duration) *AMQPConnectionPool {
	return &AMQPConnectionPool{
		maxSize:     maxSize,
		url:         url,
		maxLifetime: maxLifetime,
		maxIdleTime: maxIdleTime,
		connections: make([]*amqp.Connection, 0, maxSize),
	}
}

func (p *AMQPConnectionPool) Get(ctx context.Context) (*amqp.Connection, error) {
	p.mu.Lock()
	defer p.mu.Unlock()

	if len(p.connections) > 0 {
		conn := p.connections[len(p.connections)-1]
		p.connections = p.connections[:len(p.connections)-1]
		return conn, nil
	}

	if p.currentSize >= p.maxSize {
		//return nil, ErrConnectionPoolFull
	}

	conn, err := amqp.Dial(p.url)
	if err != nil {
		return nil, err
	}

	p.currentSize++
	go p.manageConnection(conn)

	return conn, nil
}

func (p *AMQPConnectionPool) Put(conn *amqp.Connection) {
	p.mu.Lock()
	defer p.mu.Unlock()

	if conn.IsClosed() {
		p.currentSize--
		return
	}

	p.connections = append(p.connections, conn)
}

func (p *AMQPConnectionPool) Close() {
	p.mu.Lock()
	defer p.mu.Unlock()

	for _, conn := range p.connections {
		conn.Close()
	}
	p.connections = nil
	p.currentSize = 0
}

func (p *AMQPConnectionPool) manageConnection(conn *amqp.Connection) {
	timer := time.NewTimer(p.maxLifetime)
	idleTimer := time.NewTimer(p.maxIdleTime)

	for {
		select {
		case <-timer.C:
			p.closeConnection(conn)
			return
		case <-idleTimer.C:
			p.closeConnection(conn)
			return
		case <-conn.NotifyClose(make(chan *amqp.Error)):
			p.removeConnection(conn)
			return
		}
	}
}

func (p *AMQPConnectionPool) closeConnection(conn *amqp.Connection) {
	p.mu.Lock()
	defer p.mu.Unlock()

	conn.Close()
	p.removeConnectionLocked(conn)
}

func (p *AMQPConnectionPool) removeConnection(conn *amqp.Connection) {
	p.mu.Lock()
	defer p.mu.Unlock()

	p.removeConnectionLocked(conn)
}

func (p *AMQPConnectionPool) removeConnectionLocked(conn *amqp.Connection) {
	for i, c := range p.connections {
		if c == conn {
			p.connections = append(p.connections[:i], p.connections[i+1:]...)
			break
		}
	}
	p.currentSize--
}
