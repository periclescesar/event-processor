package main

import (
	"flag"
	"log"
	"os"
	"path/filepath"
	"time"

	"github.com/periclescesar/event-processor/configs"
	"github.com/periclescesar/event-processor/pkg/rabbitmq"
)

const defaultDelay = 3

func main() {
	configs.InitConfigs()

	if err := rabbitmq.Connect(configs.Rabbitmq.URI); err != nil {
		log.Fatalf("connection failure: %v", err)
	}

	single := flag.Bool("single-message", false, "a bool")
	eventPath := flag.String("event-path", "test/mocked-events/user-created.json", "path to event, json file")
	eventsPath := flag.String("all-event-on-path", "test/mocked-events", "path to all events")
	d := flag.Duration("delay", defaultDelay, "delay in milliseconds between messages")

	flag.Parse()

	if *single {
		publishSingleMessage(*eventPath)
		return
	}

	loadEventsFromPath(*eventsPath, *d)
}

func loadEventsFromPath(path string, delay time.Duration) {
	for {
		files, err := os.ReadDir(path)
		if err != nil {
			log.Fatalf("retrieving files on %s: %v", path, err)
		}

		for _, file := range files {
			if file.IsDir() {
				continue
			}

			fullPath := filepath.Join(path, file.Name())
			byteFile, errR := os.ReadFile(fullPath)
			if errR != nil {
				log.Fatalf("read file %s: %v", fullPath, errR)
			}

			publish(string(byteFile))
			time.Sleep(delay)
		}
	}
}

func publishSingleMessage(path string) {
	file, errR := os.ReadFile(path)
	if errR != nil {
		log.Fatalf("read event file: %v", errR)
	}

	publish(string(file))
}

func publish(message string) {
	err := rabbitmq.Publish("events.exchange", message)
	if err != nil {
		log.Fatalf("publish failure: %v", err)
	}

	log.Printf("message sent: %v", message)

	errC := rabbitmq.Close()
	if errC != nil {
		log.Fatalf("rabbitmq graceful shutdown: %v", errC)
	}
}
