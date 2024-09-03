package services

import (
	"context"
	"fmt"

	"github.com/periclescesar/event-processor/internal/application/event"
)

// EventSaver defines the method for saving raw event data.
type EventSaver interface {
	// Save stores the given raw event data.
	// It returns an error if the data could not be saved.
	Save(ctx context.Context, data []byte) error
}

// EventService provides methods for handling events, including validation and saving.
type EventService struct {
	validator EventValidator  // Validator to check the validity of events.
	repo      EventRepository // Repository to store events.
}

// NewEventService creates a new instance of EventService with the provided validator and repository.
func NewEventService(validator EventValidator, repo EventRepository) *EventService {
	return &EventService{validator: validator, repo: repo}
}

// Save processes and saves a raw event.
// It decodes the raw event data into an Event, validates it, and then stores it in the repository.
// It returns an error if any of these steps fail.
func (es *EventService) Save(ctx context.Context, rawEvent []byte) error {
	ev, err := event.NewEventFromBytes(rawEvent)
	if err != nil {
		return fmt.Errorf("parse event: %w", err)
	}

	errValid := es.validator.Validate(ctx, ev)
	if errValid != nil {
		return fmt.Errorf("event validate: %w", errValid)
	}

	errSave := es.repo.Save(ctx, ev)
	if errSave != nil {
		return fmt.Errorf("save event: %w", errSave)
	}

	return nil
}
