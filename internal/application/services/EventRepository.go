package services

import (
	"context"

	"github.com/periclescesar/event-processor/internal/application/event"
)

type EventRepository interface {
	Save(ctx context.Context, event *event.Event) error
}
