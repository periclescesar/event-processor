package event

type Event struct {
	EventType string `json:"eventType"`
}

func (e *Event) SchemaId() string {
	return fmt.Sprintf("file://%s/%s.schema.json", SchemasPath, e.EventType)
}
