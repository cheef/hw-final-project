package tests

import (
	"github.com/cheef/hw-final-project/internal/storage/repository"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestExceptionListRepository(t *testing.T) {
	t.Run("when invalid CIDR provided", func(t *testing.T) {
		require.Equal(t, false, repository.IsCIDR("test"))
	})

	t.Run("when valid CIDR provided", func(t *testing.T) {
		require.Equal(t, true, repository.IsCIDR("192.168.0.1/25"))
	})

	t.Run("when IP provided", func(t *testing.T) {
		require.Equal(t, false, repository.IsCIDR("192.168.0.1"))
	})
}
