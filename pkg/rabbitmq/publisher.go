package rabbitmq

import (
	"fmt"
	amqp "github.com/AsidStorm/go-amqp-reconnect/rabbitmq"
)

func Publish(exchange string, body string) error {
	ch, err := NewChannel()
	if err != nil {
		return err
	}

	err = ch.Publish(
		exchange,
		"",
		false,
		false,
		amqp.Publishing{
			Body: []byte(body),
		},
	)

	if err != nil {
		return fmt.Errorf("publisher: %w", err)
	}

	return nil
}
