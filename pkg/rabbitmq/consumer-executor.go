package rabbitmq

import (
	"fmt"
	"log"

	"github.com/rabbitmq/amqp091-go"
)

func StartConsuming(consumer func(amqp091.Delivery) error) error {
	var ch, err = NewChannel()
	if err != nil {
		return err
	}

	msgs, err := ch.Consume(
		"events",
		"",
		false,
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
				log.Printf("consumming message: %v", errC)
				errR := d.Reject(false)
				if errR != nil {
					log.Printf("reject message: %v", errR)
				}
				continue
			}
			ackErr := d.Ack(false)
			if ackErr != nil {
				log.Printf("ack message: %v", ackErr)
			}
		}
	}()

	log.Printf(" [*] Waiting for messages")
	<-forever
	log.Printf(" [*] Closing channel")
	return nil
}
