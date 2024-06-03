package main

import (
	"context"
	"fmt"
	"github.com/cheef/hw-final-project/internal/config"
	pb "github.com/cheef/hw-final-project/pkg/server/grpc/api/grpc"
	"github.com/spf13/cobra"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"net"
	"os"
	"strconv"
)

func main() {
	const op = "CLI"

	ctx := context.Background()
	cfg, err := config.Load()

	if err != nil {
		fmt.Println(fmt.Errorf("%s: %w", op, err))
		os.Exit(1)
	}

	client, err := NewGRPCClient(cfg.GRPC.Port)

	if err != nil {
		fmt.Println(fmt.Errorf("%s: %w", op, err))
		os.Exit(1)
	}

	var cmdBlacklist = &cobra.Command{
		Use: "blacklist",
	}

	var cmdWhitelist = &cobra.Command{
		Use: "whitelist",
	}

	var cmdBlacklistAdd = &cobra.Command{
		Use:   "add [CIDR here]",
		Short: "Add CIDR to the blacklist",
		Args:  cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			for _, cidr := range args {
				message := pb.CIDRRequest{Cidr: cidr}
				_, err := client.BlacklistAdd(ctx, &message)

				if err != nil {
					fmt.Println(err)
				}
			}
		},
	}

	var cmdBlacklistRemove = &cobra.Command{
		Use:   "remove [CIDR here]",
		Short: "Remove CIDR to the blacklist",
		Args:  cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			for _, cidr := range args {
				message := pb.CIDRRequest{Cidr: cidr}
				_, err := client.BlacklistRemove(ctx, &message)

				if err != nil {
					fmt.Println(err)
				}
			}
		},
	}

	var cmdWhitelistAdd = &cobra.Command{
		Use:   "add [CIDR here]",
		Short: "Add CIDR to the whitelist",
		Args:  cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			for _, cidr := range args {
				message := pb.CIDRRequest{Cidr: cidr}
				_, err := client.WhitelistAdd(ctx, &message)

				if err != nil {
					fmt.Println(err)
				}
			}
		},
	}

	var cmdWhitelistRemove = &cobra.Command{
		Use:   "remove [CIDR here]",
		Short: "Remove CIDR to the whitelist",
		Args:  cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			for _, cidr := range args {
				message := pb.CIDRRequest{Cidr: cidr}
				_, err := client.WhitelistRemove(ctx, &message)

				if err != nil {
					fmt.Println(err)
				}
			}
		},
	}

	var cmdFlushBucket = &cobra.Command{
		Use:   "flush-bucket [login here] [ip here]",
		Short: "Flush bucket for provided login and IP-address",
		Args:  cobra.MinimumNArgs(2),
		Run: func(cmd *cobra.Command, args []string) {
			message := pb.FlushBucketRequest{Login: args[0], Ip: args[1]}
			_, err := client.FlushBucket(ctx, &message)

			if err != nil {
				fmt.Println(err)
			}
		},
	}

	var rootCmd = &cobra.Command{Use: "app"}

	rootCmd.AddCommand(cmdBlacklist, cmdWhitelist, cmdFlushBucket)
	cmdBlacklist.AddCommand(cmdBlacklistAdd, cmdBlacklistRemove)
	cmdWhitelist.AddCommand(cmdWhitelistAdd, cmdWhitelistRemove)
	err = rootCmd.Execute()

	if err != nil {
		fmt.Println(fmt.Errorf("%s: %w", op, err))
		os.Exit(1)
	}

	os.Exit(0)
}

func NewGRPCClient(port int) (pb.BFAProtectionClient, error) {
	var opts []grpc.DialOption
	opts = append(opts, grpc.WithTransportCredentials(insecure.NewCredentials()))

	grpcAddress := net.JoinHostPort("localhost", strconv.Itoa(port))
	cc, err := grpc.NewClient(grpcAddress, opts...)

	if err != nil {
		return nil, err
	}

	return pb.NewBFAProtectionClient(cc), nil
}
