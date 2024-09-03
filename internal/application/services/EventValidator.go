package services

import (
	"context"

	"github.com/periclescesar/event-processor/internal/application/event"
)

// EventValidator defines the method for validating events.
type EventValidator interface {
	// Validate checks the given event for validity.
	// It returns an error if the event is not valid.
	Validate(ctx context.Context, event *event.Event) error
}
