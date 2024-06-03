package bfaprotection

import (
	"context"
	"github.com/cheef/hw-final-project/internal/domain/models"
	"github.com/cheef/hw-final-project/internal/services"
	"github.com/cheef/hw-final-project/internal/storage/repository"
	sqlstorage "github.com/cheef/hw-final-project/internal/storage/sql"
	"log/slog"
)

type BfaProtection struct {
	log                   *slog.Logger
	limitChecker          *services.LimitChecker
	exceptionListProvider repository.ExceptionListProvider
}

func New(log *slog.Logger, s *sqlstorage.Storage, checker *services.LimitChecker) *BfaProtection {
	return &BfaProtection{
		log:                   log,
		limitChecker:          checker,
		exceptionListProvider: repository.ExceptionListProvider{Storage: s},
	}
}

func (a *BfaProtection) Authorize(_ context.Context, login, password, ip string) (bool, error) {
	a.log.Debug(
		"grpc Authorize",
		slog.String("login", login),
		slog.String("password", password),
		slog.String("ip", ip),
	)

	exceptionsList, err := a.exceptionListProvider.ShowExceptionLists()

	if err != nil {
		return false, err
	}

	a.log.Debug("Search for a black/white lists", slog.Int("found records", len(exceptionsList)))

	return a.limitChecker.IsAllowed(exceptionsList, login, password, ip), nil
}

func (a *BfaProtection) FlushBucket(_ context.Context, login, ip string) error {
	a.log.Debug(
		"grpc FlushBucket",
		slog.String("login", login),
		slog.String("ip", ip),
	)

	a.limitChecker.LoginLimiter.RemoveBucket(services.Credential(login))
	a.limitChecker.IPLimiter.RemoveBucket(services.Credential(ip))

	return nil
}

func (a *BfaProtection) BlacklistAdd(_ context.Context, cidr string) error {
	a.log.Debug(
		"grpc BlacklistAdd",
		slog.String("cidr", cidr),
	)

	el := models.ExceptionList{Type: "blacklist", CIDR: cidr}

	if _, err := a.exceptionListProvider.CreateExceptionList(el); err != nil {
		return err
	}

	return nil
}

func (a *BfaProtection) BlacklistRemove(_ context.Context, cidr string) error {
	a.log.Debug(
		"grpc BlacklistRemove",
		slog.String("cidr", cidr),
	)

	el := models.ExceptionList{Type: "blacklist", CIDR: cidr}

	if err := a.exceptionListProvider.DeleteExceptionList(el); err != nil {
		return err
	}

	return nil
}

func (a *BfaProtection) WhitelistAdd(_ context.Context, cidr string) error {
	a.log.Debug(
		"grpc WhitelistAdd",
		slog.String("cidr", cidr),
	)

	el := models.ExceptionList{Type: "whitelist", CIDR: cidr}

	if _, err := a.exceptionListProvider.CreateExceptionList(el); err != nil {
		return err
	}

	return nil
}

func (a *BfaProtection) WhitelistRemove(_ context.Context, cidr string) error {
	a.log.Debug(
		"grpc WhitelistRemove",
		slog.String("cidr", cidr),
	)

	el := models.ExceptionList{Type: "whitelist", CIDR: cidr}

	if err := a.exceptionListProvider.DeleteExceptionList(el); err != nil {
		return err
	}

	return nil
}
