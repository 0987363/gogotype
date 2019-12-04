package gogotype

import (
	"database/sql"
	"encoding/json"
)

type NullInt64 struct {
	sql.NullInt64
}

func (ni NullInt64) Size() int {
	return 8
}

func (ni NullInt64) Marshal() ([]byte, error) {
	if !ni.Valid {
		return []byte{}, nil
	}
	return json.Marshal(ni.Int64)
}

func (ni *NullInt64) Unmarshal(data []byte) error {
	err := json.Unmarshal(data, &ni.Int64)
	ni.Valid = (err == nil)
	return err
}

type NullString struct {
	sql.NullString
}

func (ni NullString) Size() int {
	return len(ns.String)
}

func (ns *NullString) Marshal() ([]byte, error) {
	if !ns.Valid {
		return []byte{}, nil
	}
	return json.Marshal(ns.String)
}

func (ns *NullString) Unmarshal(b []byte) error {
	err := json.Unmarshal(b, &ns.String)
	ns.Valid = (err == nil)
	return err
}
