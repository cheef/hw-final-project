package tests

import (
	services "github.com/cheef/hw-final-project/internal/services"
	"github.com/stretchr/testify/require"
	"log/slog"
	"testing"
	"time"
)

func TestTokenBucket(t *testing.T) {
	log := slog.Default()
	bucket := services.NewTokenBucket(10, 1000, log)

	t.Run("use all tokens", func(t *testing.T) {
		for i := 1; i <= 10; i++ {
			require.Equal(t, true, bucket.UseToken())
		}

		for i := 1; i <= 10; i++ {
			require.Equal(t, false, bucket.UseToken())
		}

		time.Sleep(time.Duration(1200) * time.Millisecond)

		require.Equal(t, true, bucket.UseToken())
	})
}
