package rabbitmq

import (
	"fmt"

	amqp "github.com/AsidStorm/go-amqp-reconnect/rabbitmq"
)

type Manager struct {
	uri  string
	conn *amqp.Connection
	ch   *amqp.Channel
}

var manager *Manager

func Connect(uri string) error {
	conn, err := amqp.Dial(uri)
	if err != nil {
		return fmt.Errorf("rabbitmq connection: %w", err)
	}

	manager = &Manager{uri, conn, nil}

	return nil
}

func NewChannel() (*amqp.Channel, error) {
	if manager.conn.IsClosed() {
		err := Connect(manager.uri)
		if err != nil {
			return nil, err
		}
	}

	if manager.ch != nil && !manager.ch.IsClosed() {
		return manager.ch, nil
	}

	ch, err := manager.conn.Channel()
	if err != nil {
		return nil, fmt.Errorf("rabbitmq channel: %w", err)
	}

	manager.ch = ch

	return ch, nil
}

func Close() error {
	errCh := manager.ch.Close()
	if errCh != nil {
		return fmt.Errorf("rabbitmq channel close: %w", errCh)
	}

	errConn := manager.conn.Close()
	if errConn != nil {
		return fmt.Errorf("rabbitmq connection close: %w", errConn)
	}

	return nil
}
