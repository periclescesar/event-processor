package main

import (
	"flag"
	"github.com/periclescesar/event-processor/configs"
	"github.com/periclescesar/event-processor/pkg/rabbitmq"
	"log"
	"os"
	"path/filepath"
	"time"
)

func main() {
	configs.InitConfigs()

	if err := rabbitmq.Connect(configs.Rabbitmq.Uri); err != nil {
		log.Fatalf("connection failure: %v", err)
	}

	single := flag.Bool("single-message", false, "a bool")
	eventPath := flag.String("event-path", "test/mocked-events/user-created.json", "path to event, json file")

	if *single {
		publishSingleMessage(*eventPath)
		return
	}

	eventsPath := flag.String("all-event-on-path", "test/mocked-events", "path to all events")
	loadEventsFromPath(*eventsPath)
}

func loadEventsFromPath(path string) {
	for {
		files, err := os.ReadDir(path)
		if err != nil {
			log.Fatalf("retriving files on %s: %v", path, err)
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
			time.Sleep(5 * time.Second)
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
