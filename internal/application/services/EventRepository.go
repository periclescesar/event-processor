package services

import (
	"context"

	"github.com/periclescesar/event-processor/internal/application/event"
)

//go:generate mockery --name EventRepository
type EventRepository interface {
	Save(ctx context.Context, event *event.Event) error
}
