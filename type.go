package gogotype

import (
	"database/sql"
	"encoding/json"
	"strconv"
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
	if ns.Valid {
		return strconv.FormatInt(ni.Int64, 10)
	}
	return "0"
}

type NullString struct {
	sql.NullString
}

func (ns NullString) Size() int {
	d, _ := json.Marshal(&ns)
	return len(d)
}

func (ns NullString) Marshal() ([]byte, error) {
	return json.Marshal(&ns)
}

func (ns *NullString) Unmarshal(data []byte) error {
	return json.Unmarshal(data, &ns)
}

func (ns NullString) String() string {
	if ns.Valid {
		return ns.String
	}
	return ""
}
