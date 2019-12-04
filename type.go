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
	d, _ := json.Marshal(&ns)
	return string(d)
}
