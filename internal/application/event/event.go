package event

// Event represents an event with a type, tenant ID, and raw data.
// The EventType indicates the kind of event, TenantID specifies the tenant to which the event belongs,
// and RawData contains the raw event data in byte format.
type Event struct {
	EventType string `json:"eventType"` // EventType indicates the type of event.
	TenantID  string `json:"tenantId"`  // TenantID specifies the tenant to which the event belongs.
	RawData   []byte // RawData contains the raw event data.
}
