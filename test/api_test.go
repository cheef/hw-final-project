package tests

import (
	grpc "github.com/cheef/hw-final-project/pkg/server/grpc/api/grpc"
	suite "github.com/cheef/hw-final-project/test/suite"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestAPI(t *testing.T) {
	ctx, st := suite.New(t)

	t.Run("Authorize", func(t *testing.T) {
		message := grpc.AuthorizeRequest{
			Login:    "test",
			Password: "test",
			Ip:       "192.168.0.1/25",
		}

		resp, err := st.Client.Authorize(ctx, &message)

		require.NoError(t, err)
		require.Equal(t, true, resp.GetOk())
	})

	t.Run("FlushBucket", func(t *testing.T) {
		message := grpc.FlushBucketRequest{
			Login: "test",
			Ip:    "192.168.0.1/25",
		}

		_, err := st.Client.FlushBucket(ctx, &message)

		require.NoError(t, err)
	})

	t.Run("BlacklistAdd", func(t *testing.T) {
		message := grpc.IPRequest{
			Ip: "192.168.0.1/25",
		}

		_, err := st.Client.BlacklistAdd(ctx, &message)

		require.NoError(t, err)
	})

	t.Run("BlacklistRemove", func(t *testing.T) {
		message := grpc.IPRequest{
			Ip: "192.168.0.1/25",
		}

		_, err := st.Client.BlacklistRemove(ctx, &message)

		require.NoError(t, err)
	})

	t.Run("WhitelistAdd", func(t *testing.T) {
		message := grpc.IPRequest{
			Ip: "192.168.0.1/25",
		}

		_, err := st.Client.WhitelistAdd(ctx, &message)

		require.NoError(t, err)
	})

	t.Run("WhitelistRemove", func(t *testing.T) {
		message := grpc.IPRequest{
			Ip: "192.168.0.1/25",
		}

		_, err := st.Client.WhitelistRemove(ctx, &message)

		require.NoError(t, err)
	})
}
