package rule

import (
	"database/sql"
	"encoding/json"
)

type NullString struct {
	sql.NullString
}

// MarshalJSON for NullString
func (ns NullString) MarshalJSON() ([]byte, error) {
	if !ns.Valid {
		return []byte(`""`), nil
	}
	return json.Marshal(ns.String)
}

func (v *NullString) UnmarshalJSON(data []byte) error {
	var x *string
	if err := json.Unmarshal(data, &x); err != nil {
		return err
	}
	if x != nil {
		v.String = *x
		v.Valid = true
	} else {
		v.Valid = false
	}
	return nil
}

type NullInt struct {
	sql.NullInt64
}

// MarshalJSON for NullInt64
func (ni NullInt) MarshalJSON() ([]byte, error) {
	if !ni.Valid {
		return []byte(`"0"`), nil
	}
	return json.Marshal(ni.Int64)
}
