package dao

import (
	"database/sql"
)

type PrivateData struct {
	UserID      sql.NullString `db:"user_id" goqu:"omitempty"`
	FirstName   sql.NullString `db:"first_name" goqu:"omitempty"`
	LastName    sql.NullString `db:"last_name" goqu:"omitempty"`
	MiddleName  sql.NullString `db:"middle_name" goqu:"omitempty"`
	DateOfBirth sql.NullTime   `db:"date_of_birth" goqu:"omitempty"`
	Gender      sql.NullString `db:"gender" goqu:"omitempty"`
}
