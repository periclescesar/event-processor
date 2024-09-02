package services

import (
	"context"
	"github.com/periclescesar/event-processor/internal/application/event"
)

//go:generate mockery --name EventValidator
type EventValidator interface {
	Validate(ctx context.Context, event *event.Event) error
}
