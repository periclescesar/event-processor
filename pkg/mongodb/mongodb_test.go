package mongodb_test

import (
	"context"
	"testing"

	"github.com/periclescesar/event-processor/pkg/mongodb"
	"github.com/stretchr/testify/require"
)

func TestConnect(t *testing.T) {
	ctx := context.TODO()

	t.Run("fails to connect to MongoDB", func(t *testing.T) {
		err := mongodb.Connect(ctx, "invalid_uri", "testdb")
		require.Error(t, err)
	})
}
