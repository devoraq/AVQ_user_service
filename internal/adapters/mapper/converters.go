package mapper

import (
	"database/sql"
	"time"
)

func toNullString(val string) sql.NullString {
	return sql.NullString{
		String: val,
		Valid:  val != "",
	}
}

func toNullTimestamp(t time.Time) sql.NullTime {
	if t.IsZero() {
		return sql.NullTime{Valid: false}
	}
	return sql.NullTime{
		Time:  t,
		Valid: true,
	}
}

func toNullBool(val bool) sql.NullBool {
	return sql.NullBool{
		Bool:  val,
		Valid: true,
	}
}

func toNullInt16(val int16) sql.NullInt16 {
	return sql.NullInt16{
		Int16: val,
		Valid: true,
	}
}
