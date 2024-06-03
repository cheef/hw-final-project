package tests

import (
	services "github.com/cheef/hw-final-project/internal/services"
	"github.com/stretchr/testify/require"
	"log/slog"
	"testing"
	"time"
)

func TestCredentialLimiter(t *testing.T) {
	log := slog.Default()

	t.Run("splits credentials into separate buckets", func(t *testing.T) {
		limiter := services.NewCredentialLimiter(10, 1000, 4000, log)

		for i := 1; i <= 10; i++ {
			require.Equal(t, true, limiter.IsAllowed("admin"))
		}

		require.Equal(t, false, limiter.IsAllowed("admin"))
		require.Equal(t, true, limiter.IsAllowed("user"))

		time.Sleep(time.Duration(1200) * time.Millisecond)

		require.Equal(t, true, limiter.IsAllowed("admin"))
	})

	t.Run("sweeps old buckets", func(t *testing.T) {
		limiter := services.NewCredentialLimiter(10, 6000, 1000, log)

		for i := 1; i <= 10; i++ {
			require.Equal(t, true, limiter.IsAllowed("admin"))
		}

		require.Equal(t, false, limiter.IsAllowed("admin"))

		time.Sleep(time.Duration(1100) * time.Millisecond)

		require.Equal(t, true, limiter.IsAllowed("admin"))
	})
}
