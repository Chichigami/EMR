package handlers

import (
	"database/sql"
	"strconv"
)

func NullString(s string) sql.NullString {
	return sql.NullString{
		String: s,
		Valid:  s != "",
	}
}

func NullInt32(i int32) sql.NullInt32 {
	return sql.NullInt32{
		Int32: i,
		Valid: i != 0,
	}
}

func ConvertStringToInt32(s string) (int32, error) {
	d, err := strconv.Atoi(s)
	if err != nil {
		return -1, err
	}
	return int32(d), nil
}
