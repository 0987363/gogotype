package gogotype

import (
	"database/sql"
	"encoding/json"
)

type NullInt64 struct {
	sql.NullInt64
}

func (ni NullInt64) Size() int {
	d, _ := json.Marshal(&ni)
	return len(d)
}

func (ni NullInt64) Marshal() ([]byte, error) {
	return json.Marshal(&ni)
}

func (ni *NullInt64) Unmarshal(data []byte) error {
	return json.Unmarshal(data, &ni)
}

func (ni NullInt64) String() string {
	d, _ := json.Marshal(&ni)
	return string(d)
}

type NullString struct {
	sql.NullString
}

func (ns NullString) Size() int {
	return len(ns.String)
}

func (ns NullString) Marshal() ([]byte, error) {
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
