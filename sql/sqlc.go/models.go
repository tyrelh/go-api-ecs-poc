// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.28.0

package sql

import (
	"database/sql"
)

type Reward struct {
	ID           uint64
	Brand        sql.NullString
	Currency     sql.NullString
	Denomination sql.NullFloat64
}
