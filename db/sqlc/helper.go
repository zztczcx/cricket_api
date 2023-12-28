package db

import (
	"database/sql"
	"strconv"
)

func ToNullInt64(s string) sql.NullInt64 {
	i, err := strconv.Atoi(s)
	return sql.NullInt64{Int64: int64(i), Valid: err == nil}
}

func ToNullFloat64(s string) sql.NullFloat64 {
	f, err := strconv.ParseFloat(s, 64)
	return sql.NullFloat64{Float64: f, Valid: err == nil}
}
