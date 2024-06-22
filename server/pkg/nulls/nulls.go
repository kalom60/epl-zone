package nulls

import (
	"database/sql"
	"encoding/json"
)

type NullString struct {
	sql.NullString
}

func (ns NullString) MarshalJSON() ([]byte, error) {
	if ns.Valid {
		return json.Marshal(ns.String)
	}

	return json.Marshal(nil)
}

type NullFloat64 struct {
	sql.NullFloat64
}

func (ns NullFloat64) MarshalJSON() ([]byte, error) {
	if ns.Valid {
		return json.Marshal(ns.Float64)
	}

	return json.Marshal(nil)
}

type NullInt64 struct {
	sql.NullInt64
}

func (ns NullInt64) MarshalJSON() ([]byte, error) {
	if ns.Valid {
		return json.Marshal(ns.Int64)
	}

	return json.Marshal(nil)
}
