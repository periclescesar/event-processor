package rabbitmq

import (
	"fmt"
	"github.com/rabbitmq/amqp091-go"
	"log"
)

func StartConsuming(consumer func(amqp091.Delivery) error) error {
	var ch, err = NewChannel()
	if err != nil {
		return err
	}

	msgs, err := ch.Consume(
		"events",
		"",
		true,
		false,
		false,
		false,
		nil,
	)

	if err != nil {
		return fmt.Errorf("rabbitmq registry consume: %w", err)
	}

	var forever chan struct{}

	go func() {
		for d := range msgs {
			if errC := consumer(d); errC != nil {
				errR := d.Reject(false)
				if errR != nil {
					log.Fatalf("reject message: %v: %v", errR, errC)
				}
				log.Printf("reject message: %v", errC)
			}
		}
	}()

	<-forever
	return nil
}
