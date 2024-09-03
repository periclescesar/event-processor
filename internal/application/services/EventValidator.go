package services

import (
	"context"

	"github.com/periclescesar/event-processor/internal/application/event"
)

type EventValidator interface {
	Validate(ctx context.Context, event *event.Event) error
}
