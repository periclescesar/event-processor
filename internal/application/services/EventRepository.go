package services

import (
	"context"

	"github.com/periclescesar/event-processor/internal/application/event"
)

// EventRepository defines the methods for storing events.
type EventRepository interface {
	// Save stores the given event in the repository.
	// It returns an error if the event could not be saved.
	Save(ctx context.Context, event *event.Event) error
}
