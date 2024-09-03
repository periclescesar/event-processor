package services

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/periclescesar/event-processor/internal/application/event"
)

type EventService struct {
	validator EventValidator
	repo      EventRepository
}

func NewEventService(validator EventValidator, repo EventRepository) *EventService {
	return &EventService{validator: validator, repo: repo}
}

func (es *EventService) Save(ctx context.Context, rawEvent []byte) error {
	ev := &event.Event{RawData: rawEvent}
	err := json.Unmarshal(rawEvent, ev)
	if err != nil {
		return fmt.Errorf("event decode: %w", err)
	}

	errValid := es.validator.Validate(ctx, ev)
	if errValid != nil {
		return errValid
	}

	return es.repo.Save(ctx, ev)
}
