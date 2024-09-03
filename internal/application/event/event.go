package event

import (
	"encoding/json"
	"fmt"
)

// Event represents an event with a type, tenant ID, and raw data.
// The EventType indicates the kind of event, TenantID specifies the tenant to which the event belongs,
// and RawData contains the raw event data in byte format.
type Event struct {
	EventType string `json:"eventType"` // EventType indicates the type of event.
	TenantID  string `json:"tenantId"`  // TenantID specifies the tenant to which the event belongs.
	RawData   []byte // RawData contains the raw event data.
}

func NewEventFromBytes(rawData []byte) (*Event, error) {
	ev := &Event{RawData: rawData}
	err := json.Unmarshal(rawData, &ev)
	if err != nil {
		return nil, fmt.Errorf("event decode: %w", err)
	}

	return ev, nil
}

func (e *Event) ToMap() (map[string]interface{}, error) {
	fullEvent := make(map[string]interface{})
	err := json.Unmarshal(e.RawData, &fullEvent)
	if err != nil {
		return nil, fmt.Errorf("map conversion: %w", err)
	}

	return fullEvent, nil
}
