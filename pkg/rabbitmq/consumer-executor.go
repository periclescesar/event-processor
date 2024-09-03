package rabbitmq

import (
	"fmt"

	log "github.com/sirupsen/logrus"

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
				log.Debugf("consuming message: %v", errC)
				errR := d.Reject(false)
				if errR != nil {
					log.Errorf("reject message: %v", errR)
				}

				continue
			}
			ackErr := d.Ack(false)
			if ackErr != nil {
				log.Errorf("ack message: %v", ackErr)
			}
		}
	}()

	log.Info("[*] Waiting for messages")
	<-forever
	log.Info("[*] Closing channel")

	return nil
}
