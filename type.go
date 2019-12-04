package gogotype

import (
	"database/sql"
	"encoding/json"
)

type NullInt64 struct {
	sql.NullInt64
}

func (this NullInt64) MarshalJSON() ([]byte, error) {
	return json.Marshal([]byte(this))
	/*
	if !ni.Valid {
		return []byte{}, nil
	}
	return json.Marshal(ni.Int64)
	*/
}

func (this *NullInt64) UnmarshalJSON(data []byte) error {
	/*
	err := json.Unmarshal(b, &ni.Int64)
	ni.Valid = (err == nil)
	return err
	*/
	v := new([]byte)
	err := json.Unmarshal(data, v)
	if err != nil {
		return err
	}
	*this = *v
	return nil

}

type NullString struct {
	sql.NullString
}

// MarshalJSON for NullString
func (ns *NullString) MarshalJSON() ([]byte, error) {
	if !ns.Valid {
		return []byte("null"), nil
	}
	return json.Marshal(ns.String)
}

func (ns *NullString) UnmarshalJSON(b []byte) error {
	err := json.Unmarshal(b, &ns.String)
	ns.Valid = (err == nil)
	return err
}
