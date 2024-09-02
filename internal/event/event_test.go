package event

import (
	"fmt"
	"testing"
)

func TestEvent_SchemaId(t *testing.T) {
	type fields struct {
		EventType string
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{name: "event type a", fields: fields{EventType: "a"}, want: fmt.Sprintf("file://%s/a.schema.json", SchemasPath)},
		{name: "event type b", fields: fields{EventType: "b"}, want: fmt.Sprintf("file://%s/b.schema.json", SchemasPath)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := &Event{
				EventType: tt.fields.EventType,
			}
			if got := e.SchemaId(); got != tt.want {
				t.Errorf("SchemaId() = %v, want %v", got, tt.want)
			}
		})
	}
}
