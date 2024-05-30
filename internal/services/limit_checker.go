package services

import (
	"github.com/cheef/hw-final-project/internal/config"
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

func (c *LimitChecker) IsAllowed(login, password, ip string) bool {
	return c.LoginLimiter.IsAllowed(Credential(login)) &&
		c.PasswordLimiter.IsAllowed(Credential(password)) &&
		c.IPLimiter.IsAllowed(Credential(ip))
}
