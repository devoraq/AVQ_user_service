package mapping

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

func toNullInt32(val int32) sql.NullInt32 {
	return sql.NullInt32{
		Int32: val,
		Valid: true,
	}
}

func toNullInt64(val int64) sql.NullInt64 {
	return sql.NullInt64{
		Int64: val,
		Valid: true,
	}
}
