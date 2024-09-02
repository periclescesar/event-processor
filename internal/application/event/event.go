package event

type Event struct {
	EventType string `json:"eventType"`
	TenantID  string `json:"tenantId"`
	RawData   []byte
}
