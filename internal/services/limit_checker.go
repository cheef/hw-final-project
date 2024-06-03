package services

import (
	"github.com/cheef/hw-final-project/internal/config"
	"github.com/cheef/hw-final-project/internal/domain/models"
	"log/slog"
)

type LimitChecker struct {
	LoginLimiter    *CredentialLimiter
	PasswordLimiter *CredentialLimiter
	IPLimiter       *CredentialLimiter
	log             *slog.Logger
}

func NewLimitChecker(cfg *config.Config, log *slog.Logger) *LimitChecker {
	return &LimitChecker{
		log: log,
		LoginLimiter: NewCredentialLimiter(
			cfg.BFALimits.Login,
			cfg.BFALimits.Period,
			cfg.BFALimits.BucketLifetime,
			log,
		),
		PasswordLimiter: NewCredentialLimiter(
			cfg.BFALimits.Password,
			cfg.BFALimits.Period,
			cfg.BFALimits.BucketLifetime,
			log,
		),
		IPLimiter: NewCredentialLimiter(
			cfg.BFALimits.IP,
			cfg.BFALimits.Period,
			cfg.BFALimits.BucketLifetime,
			log,
		),
	}
}

func (c *LimitChecker) IsAllowed(el []models.ExceptionList, login, password, ip string) bool {
	in := make(chan models.ExceptionList)

	go func() {
		defer close(in)

		for _, ex := range el {
			in <- ex
		}
	}()

	res, exType := c.testIP(ip, in)

	if res && exType == "whitelist" {
		return true
	}

	if res && exType == "blacklist" {
		return false
	}

	return c.LoginLimiter.IsAllowed(Credential(login)) &&
		c.PasswordLimiter.IsAllowed(Credential(password)) &&
		c.IPLimiter.IsAllowed(Credential(ip))
}

func (c *LimitChecker) testIP(ip string, in <-chan models.ExceptionList) (bool, string) {
	checker := NewIPChecker(in)
	result, exType, err := checker.IsInList(ip)

	if err != nil {
		c.log.Error("LimitChecker failed to test ip", slog.String("error", err.Error()))
		return false, ""
	}

	return result, exType
}
