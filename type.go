package gogotype

import (
	"bytes"
	"database/sql"
	"database/sql/driver"
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
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

func (ni *NullInt64) String() string {
	if ni.Valid {
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

func (ns *NullString) String() string {
	if ns.Valid {
		return ns.NullString.String
	}
	return ""
}

type StringArray []string

func (a StringArray) Size() int {
	var i int
	for _, s := range a {
		i += len(s)
	}
	return i
}

func (a StringArray) Marshal() ([]byte, error) {
	return json.Marshal(&a)
}

func (a *StringArray) Unmarshal(data []byte) error {
	return json.Unmarshal(data, &a)
}

func (a *StringArray) String() string {
	return fmt.Sprintf("%v", *a)
}

// Scan implements the sql.Scanner interface.
func (a *StringArray) Scan(src interface{}) error {
	switch src := src.(type) {
	case []byte:
		return a.scanBytes(src)
	case string:
		return a.scanBytes([]byte(src))
	case nil:
		*a = nil
		return nil
	}

	return fmt.Errorf("pq: cannot convert %T to StringArray", src)
}

func (a *StringArray) scanBytes(src []byte) error {
	elems, err := scanLinearArray(src, []byte{','}, "StringArray")
	if err != nil {
		return err
	}
	if *a != nil && len(elems) == 0 {
		*a = (*a)[:0]
	} else {
		b := make(StringArray, len(elems))
		for i, v := range elems {
			if b[i] = string(v); v == nil {
				return fmt.Errorf("pq: parsing array element index %d: cannot convert nil to string", i)
			}
		}
		*a = b
	}
	return nil
}

func scanLinearArray(src, del []byte, typ string) (elems [][]byte, err error) {
	dims, elems, err := parseArray(src, del)
	if err != nil {
		return nil, err
	}
	if len(dims) > 1 {
		return nil, fmt.Errorf("pq: cannot convert ARRAY%s to %s", strings.Replace(fmt.Sprint(dims), " ", "][", -1), typ)
	}
	return elems, err
}

// Value implements the driver.Valuer interface.
func (a StringArray) Value() (driver.Value, error) {
	if a == nil {
		return nil, nil
	}

	if n := len(a); n > 0 {
		// There will be at least two curly brackets, 2*N bytes of quotes,
		// and N-1 bytes of delimiters.
		b := make([]byte, 1, 1+3*n)
		b[0] = '{'

		b = appendArrayQuotedBytes(b, []byte(a[0]))
		for i := 1; i < n; i++ {
			b = append(b, ',')
			b = appendArrayQuotedBytes(b, []byte(a[i]))
		}

		return string(append(b, '}')), nil
	}

	return "{}", nil
}
func appendArrayQuotedBytes(b, v []byte) []byte {
	b = append(b, '"')
	for {
		i := bytes.IndexAny(v, `"\`)
		if i < 0 {
			b = append(b, v...)
			break
		}
		if i > 0 {
			b = append(b, v[:i]...)
		}
		b = append(b, '\\', v[i])
		v = v[i+1:]
	}
	return append(b, '"')
}
