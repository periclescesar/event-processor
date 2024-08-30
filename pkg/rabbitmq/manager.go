package rabbitmq

import (
	"fmt"
	amqp "github.com/rabbitmq/amqp091-go"
)

type Manager struct {
	conn *amqp.Connection
	ch   *amqp.Channel
}

var manager *Manager

func Connect(uri string) error {
	conn, err := amqp.Dial(uri)
	if err != nil {
		return fmt.Errorf("rabbitmq connection: %w", err)
	}

	manager = &Manager{conn, nil}
	return nil
}

func NewChannel() (*amqp.Channel, error) {
	if manager.ch != nil {
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
		return errCh
	}

	errConn := manager.conn.Close()
	if errConn != nil {
		return errConn
	}
	return nil
}
