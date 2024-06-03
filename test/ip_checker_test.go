package tests

import (
	"github.com/cheef/hw-final-project/internal/domain/models"
	"github.com/cheef/hw-final-project/internal/services"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestIPChecker(t *testing.T) {
	var subnets []models.ExceptionList
	subnets = append(subnets, models.ExceptionList{Type: "blacklist", CIDR: "192.168.0.1/25"})
	subnets = append(subnets, models.ExceptionList{Type: "blacklist", CIDR: "10.0.0.0/24"})

	getCIDRs := func([]models.ExceptionList) chan models.ExceptionList {
		in := make(chan models.ExceptionList)

		go func() {
			defer close(in)

			for _, ex := range subnets {
				in <- ex
			}
		}()

		return in
	}

	t.Run("when IP matches to one of the subnets", func(t *testing.T) {
		ip := "10.0.0.130"

		checker := services.NewIPChecker(getCIDRs(subnets))
		result, exType, err := checker.IsInList(ip)

		require.NoError(t, err)
		require.Equal(t, true, result)
		require.Equal(t, "blacklist", exType)
	})

	t.Run("when IP doesn't matches to any of the subnet", func(t *testing.T) {
		ip := "10.0.5.130"

		checker := services.NewIPChecker(getCIDRs(subnets))
		result, exType, err := checker.IsInList(ip)

		require.NoError(t, err)
		require.Equal(t, false, result)
		require.Equal(t, "", exType)
	})
}
