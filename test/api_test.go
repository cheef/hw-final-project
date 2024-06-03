package tests

import (
	grpc "github.com/cheef/hw-final-project/pkg/server/grpc/api/grpc"
	suite "github.com/cheef/hw-final-project/test/suite"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestAPI(t *testing.T) {
	ctx, st := suite.New(t)

	t.Run("when tried authorize too many times", func(t *testing.T) {
		m1 := grpc.AuthorizeRequest{
			Login:    "test",
			Password: "test",
			Ip:       "192.168.0.1",
		}

		m2 := grpc.FlushBucketRequest{
			Login: "test",
			Ip:    "192.168.0.1",
		}

		for i := 1; i <= 10; i++ {
			resp, err := st.Client.Authorize(ctx, &m1)

			require.NoError(t, err)
			require.Equal(t, true, resp.GetOk())
		}

		resp, err := st.Client.Authorize(ctx, &m1)
		require.NoError(t, err)
		require.Equal(t, false, resp.GetOk())

		_, err = st.Client.FlushBucket(ctx, &m2)
		require.NoError(t, err)
	})

	t.Run("when item added to the whitelist", func(t *testing.T) {
		m1 := grpc.AuthorizeRequest{
			Login:    "test",
			Password: "test",
			Ip:       "192.168.0.5",
		}

		m2 := grpc.CIDRRequest{
			Cidr: "192.168.0.1/25",
		}

		m3 := grpc.FlushBucketRequest{
			Login: "test",
			Ip:    "192.168.0.5",
		}

		for i := 1; i <= 10; i++ {
			resp, err := st.Client.Authorize(ctx, &m1)
			require.NoError(t, err)
			require.Equal(t, true, resp.GetOk())
		}

		resp, err := st.Client.Authorize(ctx, &m1)
		require.NoError(t, err)
		require.Equal(t, false, resp.GetOk())

		_, err = st.Client.WhitelistAdd(ctx, &m2)
		require.NoError(t, err)

		resp, err = st.Client.Authorize(ctx, &m1)
		require.NoError(t, err)
		require.Equal(t, true, resp.GetOk())

		_, err = st.Client.WhitelistRemove(ctx, &m2)
		require.NoError(t, err)

		resp, err = st.Client.Authorize(ctx, &m1)
		require.NoError(t, err)
		require.Equal(t, false, resp.GetOk())

		_, err = st.Client.FlushBucket(ctx, &m3)
		require.NoError(t, err)
	})

	t.Run("when item added to the blacklist", func(t *testing.T) {
		m1 := grpc.AuthorizeRequest{
			Login:    "test",
			Password: "test",
			Ip:       "192.168.0.5",
		}

		m2 := grpc.CIDRRequest{
			Cidr: "192.168.0.1/25",
		}

		m3 := grpc.FlushBucketRequest{
			Login: "test",
			Ip:    "192.168.0.5",
		}

		resp, err := st.Client.Authorize(ctx, &m1)
		require.NoError(t, err)
		require.Equal(t, true, resp.GetOk())

		_, err = st.Client.BlacklistAdd(ctx, &m2)
		require.NoError(t, err)

		resp, err = st.Client.Authorize(ctx, &m1)
		require.NoError(t, err)
		require.Equal(t, false, resp.GetOk())

		_, err = st.Client.BlacklistRemove(ctx, &m2)
		require.NoError(t, err)

		resp, err = st.Client.Authorize(ctx, &m1)
		require.NoError(t, err)
		require.Equal(t, true, resp.GetOk())

		_, err = st.Client.FlushBucket(ctx, &m3)
		require.NoError(t, err)
	})
}
