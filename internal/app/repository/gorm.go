package repository

import (
	"errors"

	"github.com/jackc/pgconn"
	"gorm.io/gorm"
)

type abstractError error

var (
	errRecordNotFound      = abstractError(errors.New("record not found"))
	errUniqueViolation     = abstractError(errors.New("unique violation"))
	errForeignKeyViolation = abstractError(errors.New("foreign key violation"))
)

// IsErrorGorm provides a helper for checking/comparing an error
// returning from GORM with an abstract errors.
func isErrorGorm(err error, expected abstractError) bool {
	switch expected {
	case errRecordNotFound:
		return errors.Is(err, gorm.ErrRecordNotFound)
	case errUniqueViolation:
		var pgerr *pgconn.PgError
		if errors.As(err, &pgerr) && pgerr.Code == "23505" /* unique_violation */ {
			return true
		}
	case errForeignKeyViolation:
		var pgerr *pgconn.PgError
		if errors.As(err, &pgerr) && pgerr.Code == "23503" /* foreign_key_violation */ {
			return true
		}
	}
	return false
}
