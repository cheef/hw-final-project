package repository

import (
	"context"
	"errors"
	"github.com/cheef/hw-final-project/internal/domain/models"
	sqlstorage "github.com/cheef/hw-final-project/internal/storage/sql"
	"regexp"
)

var (
	ErrExceptionFieldBlank      = errors.New("field is blank")
	ErrExceptionListUnknownType = errors.New("unknown type")
	ErrExceptionListInvalidCIDR = errors.New("invalid CIDR")
	CIDRPattern                 = regexp.MustCompile(`^(([1-9]{0,1}[0-9]{0,2}|2[0-4][0-9]|25[0-5])\.){3}([1-9]{0,1}[0-9]{0,2}|2[0-4][0-9]|25[0-5])/([1-2][0-9]|3[0-1])$`)
)

type ExceptionListProvider struct {
	Ctx     context.Context
	Storage *sqlstorage.Storage
}

func (elp *ExceptionListProvider) CreateExceptionList(el models.ExceptionList) (int64, error) {
	if el.Type == "" || el.CIDR == "" {
		return 0, ErrExceptionFieldBlank
	}

	if el.Type != "whitelist" && el.Type != "blacklist" {
		return 0, ErrExceptionListUnknownType
	}

	if !IsCIDR(el.CIDR) {
		return 0, ErrExceptionListInvalidCIDR
	}

	return elp.Storage.CreateExceptionList(elp.Ctx, el.Type, el.CIDR)
}

func (elp *ExceptionListProvider) DeleteExceptionList(el models.ExceptionList) error {
	return elp.Storage.DeleteExceptionList(elp.Ctx, el.Type, el.CIDR)
}

func (elp *ExceptionListProvider) ShowExceptionLists() ([]models.ExceptionList, error) {
	return elp.Storage.ShowExceptionLists(elp.Ctx)
}

func IsCIDR(s string) bool {
	return CIDRPattern.Match([]byte(s))
}
