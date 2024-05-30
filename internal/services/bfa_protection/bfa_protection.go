package bfaprotection

import (
	"context"
	"github.com/cheef/hw-final-project/internal/services"
	"log/slog"
)

type BfaProtection struct {
	log          *slog.Logger
	limitChecker *services.LimitChecker
}

func New(log *slog.Logger, checker *services.LimitChecker) *BfaProtection {
	return &BfaProtection{
		log:          log,
		limitChecker: checker,
	}
}

func (a *BfaProtection) Authorize(_ context.Context, login, password, ip string) (bool, error) {
	a.log.Info(
		"grpc Authorize called with",
		slog.String("login", login),
		slog.String("password", password),
		slog.String("ip", ip),
	)

	return a.limitChecker.IsAllowed(login, password, ip), nil
}

func (a *BfaProtection) FlushBucket(_ context.Context, login, ip string) error {
	a.limitChecker.LoginLimiter.RemoveBucket(services.Credential(login))
	a.limitChecker.IPLimiter.RemoveBucket(services.Credential(ip))

	return nil
}

func (a *BfaProtection) BlacklistAdd(_ context.Context, ip string) error {
	return nil
}

func (a *BfaProtection) BlacklistRemove(_ context.Context, ip string) error {
	return nil
}

func (a *BfaProtection) WhitelistAdd(_ context.Context, ip string) error {
	return nil
}

func (a *BfaProtection) WhitelistRemove(_ context.Context, ip string) error {
	return nil
}
