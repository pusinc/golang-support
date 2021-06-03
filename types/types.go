package types

import (
	"database/sql"
	"database/sql/driver"
	"encoding/json"
	"strconv"
	"time"
)

type (
	IntArray       []int
	Int8Array      []int8
	Int16Array     []int16
	Int32Array     []int32
	Int64Array     []int64
	UintArray      []uint
	Uint8Array     []uint8
	Uint16Array    []uint16
	Uint32Array    []uint32
	Uint64Array    []uint64
	Float32Array   []float32
	Float64Array   []float64
	StringArray    []string
	InterfaceArray []interface{}

	MapStringInt       map[string]int
	MapStringInt8      map[string]int8
	MapStringInt16     map[string]int16
	MapStringInt32     map[string]int32
	MapStringInt64     map[string]int64
	MapStringUint      map[string]uint
	MapStringUint8     map[string]uint8
	MapStringUint16    map[string]uint16
	MapStringUint32    map[string]uint32
	MapStringUint64    map[string]uint64
	MapStringFloat32   map[string]float32
	MapStringFloat64   map[string]float64
	MapStringString    map[string]string
	MapStringInterface map[string]interface{}

	NullTime sql.NullTime
	NullDate struct {
		Time  time.Time
		Valid bool
	}
	NullBool    sql.NullBool
	NullFloat64 sql.NullFloat64
	NullInt32   sql.NullInt32
	NullInt64   sql.NullInt64
	NullString  sql.NullString
)

func (i IntArray) Value() (driver.Value, error) {
	return jsonArrayString(i)
}

func (i *IntArray) Scan(v interface{}) error {
	return json.Unmarshal(v.([]byte), i)
}

func (i Int8Array) Value() (driver.Value, error) {
	return jsonArrayString(i)
}

func (i *Int8Array) Scan(v interface{}) error {
	return json.Unmarshal(v.([]byte), i)
}

func (i Int16Array) Value() (driver.Value, error) {
	return jsonArrayString(i)
}

func (i *Int16Array) Scan(v interface{}) error {
	return json.Unmarshal(v.([]byte), i)
}

func (i Int32Array) Value() (driver.Value, error) {
	return jsonArrayString(i)
}

func (i *Int32Array) Scan(v interface{}) error {
	return json.Unmarshal(v.([]byte), i)
}

func (i Int64Array) Value() (driver.Value, error) {
	return jsonArrayString(i)
}

func (i *Int64Array) Scan(v interface{}) error {
	return json.Unmarshal(v.([]byte), i)
}

func (u UintArray) Value() (driver.Value, error) {
	return jsonArrayString(u)
}

func (u *UintArray) Scan(v interface{}) error {
	return json.Unmarshal(v.([]byte), u)
}

func (u Uint8Array) Value() (driver.Value, error) {
	return jsonArrayString(u)
}

func (u *Uint8Array) Scan(v interface{}) error {
	return json.Unmarshal(v.([]byte), u)
}

func (u Uint16Array) Value() (driver.Value, error) {
	return jsonArrayString(u)
}

func (u *Uint16Array) Scan(v interface{}) error {
	return json.Unmarshal(v.([]byte), u)
}

func (u Uint32Array) Value() (driver.Value, error) {
	return jsonArrayString(u)
}

func (u *Uint32Array) Scan(v interface{}) error {
	return json.Unmarshal(v.([]byte), u)
}

func (u Uint64Array) Value() (driver.Value, error) {
	return jsonArrayString(u)
}

func (u *Uint64Array) Scan(v interface{}) error {
	return json.Unmarshal(v.([]byte), u)
}

func (f Float32Array) Value() (driver.Value, error) {
	return jsonArrayString(f)
}

func (f *Float32Array) Scan(v interface{}) error {
	return json.Unmarshal(v.([]byte), f)
}

func (f Float64Array) Value() (driver.Value, error) {
	return jsonArrayString(f)
}

func (f *Float64Array) Scan(v interface{}) error {
	return json.Unmarshal(v.([]byte), f)
}

func (s StringArray) Value() (driver.Value, error) {
	return jsonArrayString(s)
}

func (s *StringArray) Scan(v interface{}) error {
	return json.Unmarshal(v.([]byte), s)
}

func (i InterfaceArray) Value() (driver.Value, error) {
	return jsonArrayString(i)
}

func (i *InterfaceArray) Scan(v interface{}) error {
	return json.Unmarshal(v.([]byte), i)
}

func (m MapStringInt) Value() (driver.Value, error) {
	return jsonArrayString(m)
}

func (m *MapStringInt) Scan(v interface{}) error {
	return jsonParseForMap(v.([]byte), m)
}

func (m MapStringInt8) Value() (driver.Value, error) {
	return jsonArrayString(m)
}

func (m *MapStringInt8) Scan(v interface{}) error {
	return jsonParseForMap(v.([]byte), m)
}

func (m MapStringInt16) Value() (driver.Value, error) {
	return jsonArrayString(m)
}

func (m *MapStringInt16) Scan(v interface{}) error {
	return jsonParseForMap(v.([]byte), m)
}

func (m MapStringInt32) Value() (driver.Value, error) {
	return jsonArrayString(m)
}

func (m *MapStringInt32) Scan(v interface{}) error {
	return jsonParseForMap(v.([]byte), m)
}

func (m MapStringInt64) Value() (driver.Value, error) {
	return jsonArrayString(m)
}

func (m *MapStringInt64) Scan(v interface{}) error {
	return jsonParseForMap(v.([]byte), m)
}

func (m MapStringUint) Value() (driver.Value, error) {
	return jsonArrayString(m)
}

func (m *MapStringUint) Scan(v interface{}) error {
	return jsonParseForMap(v.([]byte), m)
}

func (m MapStringUint8) Value() (driver.Value, error) {
	return jsonArrayString(m)
}

func (m *MapStringUint8) Scan(v interface{}) error {
	return jsonParseForMap(v.([]byte), m)
}

func (m MapStringUint16) Value() (driver.Value, error) {
	return jsonArrayString(m)
}

func (m *MapStringUint16) Scan(v interface{}) error {
	return jsonParseForMap(v.([]byte), m)
}

func (m MapStringUint32) Value() (driver.Value, error) {
	return jsonArrayString(m)
}

func (m *MapStringUint32) Scan(v interface{}) error {
	return jsonParseForMap(v.([]byte), m)
}

func (m MapStringUint64) Value() (driver.Value, error) {
	return jsonArrayString(m)
}

func (m *MapStringUint64) Scan(v interface{}) error {
	return jsonParseForMap(v.([]byte), m)
}

func (m MapStringFloat32) Value() (driver.Value, error) {
	return jsonArrayString(m)
}

func (m *MapStringFloat32) Scan(v interface{}) error {
	return jsonParseForMap(v.([]byte), m)
}

func (m MapStringFloat64) Value() (driver.Value, error) {
	return jsonArrayString(m)
}

func (m *MapStringFloat64) Scan(v interface{}) error {
	return jsonParseForMap(v.([]byte), m)
}

func (m MapStringString) Value() (driver.Value, error) {
	return jsonArrayString(m)
}

func (m *MapStringString) Scan(v interface{}) error {
	return jsonParseForMap(v.([]byte), m)
}

func (m MapStringInterface) Value() (driver.Value, error) {
	return jsonArrayString(m)
}

func (m *MapStringInterface) Scan(v interface{}) error {
	return jsonParseForMap(v.([]byte), m)
}

func (t NullTime) Value() (driver.Value, error) {
	if !t.Valid {
		return nil, nil
	}
	return t.Time, nil
}

func (t *NullTime) Scan(v interface{}) error {
	return (*sql.NullTime)(t).Scan(v)
}

func (t *NullTime) UnmarshalJSON(v []byte) error {
	strV := string(v)
	if strV == "null" || strV == "" {
		t.Valid = false
		return nil
	}
	parsedTime, err := time.ParseInLocation(`"2006-01-02 15:04:05"`, strV, time.Local)
	if err != nil {
		t.Valid = false
		return err
	}
	t.Time, t.Valid = parsedTime, true
	return err
}

func (t NullTime) MarshalJSON() ([]byte, error) {
	if !t.Valid {
		return json.Marshal(nil)
	}
	return json.Marshal(t.Time)
}

func (d NullDate) Value() (driver.Value, error) {
	if !d.Valid {
		return nil, nil
	}
	return d.Time, nil
}

func (d *NullDate) Scan(v interface{}) error {
	return (*sql.NullTime)(d).Scan(v)
}

func (d *NullDate) UnmarshalJSON(v []byte) error {
	strV := string(v)
	if strV == "null" || strV == "" {
		d.Valid = false
		return nil
	}
	parsedTime, err := time.ParseInLocation(`"2006-01-02"`, strV, time.Local)
	if err != nil {
		d.Valid = false
		return err
	}
	d.Time, d.Valid = parsedTime, true
	return err
}

func (d NullDate) MarshalJSON() ([]byte, error) {
	if !d.Valid {
		return json.Marshal(nil)
	}
	return json.Marshal(d.Time)
}

func (n NullBool) Value() (driver.Value, error) {
	if !n.Valid {
		return nil, nil
	}
	return n.Bool, nil
}

func (n *NullBool) Scan(v interface{}) error {
	return (*sql.NullBool)(n).Scan(v)
}

func (n *NullBool) UnmarshalJSON(v []byte) error {
	strV := string(v)
	if strV == "null" {
		n.Valid = false
		return nil
	}
	parsedBool, err := strconv.ParseBool(strV)
	if err != nil {
		n.Valid = false
		return err
	}
	n.Bool, n.Valid = parsedBool, true
	return nil
}

func (n NullBool) MarshalJSON() ([]byte, error) {
	if !n.Valid {
		return json.Marshal(nil)
	}
	return json.Marshal(n.Bool)
}

func (n NullFloat64) Value() (driver.Value, error) {
	if !n.Valid {
		return nil, nil
	}
	return n.Float64, nil
}

func (n *NullFloat64) Scan(v interface{}) error {
	return (*sql.NullFloat64)(n).Scan(v)
}

func (n *NullFloat64) UnmarshalJSON(v []byte) error {
	strV := string(v)
	if strV == "null" {
		n.Valid = false
		return nil
	}
	parsedFloat, err := strconv.ParseFloat(strV, 32)
	if err != nil {
		n.Valid = false
		return err
	}
	n.Float64, n.Valid = parsedFloat, true
	return nil
}

func (n NullFloat64) MarshalJSON() ([]byte, error) {
	if !n.Valid {
		return json.Marshal(nil)
	}
	return json.Marshal(n.Float64)
}

func (n NullInt32) Value() (driver.Value, error) {
	if !n.Valid {
		return nil, nil
	}
	return n.Int32, nil
}

func (n *NullInt32) Scan(v interface{}) error {
	return (*sql.NullInt32)(n).Scan(v)
}

func (n *NullInt32) UnmarshalJSON(v []byte) error {
	strV := string(v)
	if strV == "null" {
		n.Valid = false
		return nil
	}
	parsedInt, err := strconv.ParseInt(strV, 10, 32)
	if err != nil {
		n.Valid = false
		return err
	}
	n.Int32, n.Valid = int32(parsedInt), true
	return nil
}

func (n NullInt32) MarshalJSON() ([]byte, error) {
	if !n.Valid {
		return json.Marshal(nil)
	}
	return json.Marshal(n.Int32)
}

func (n NullInt64) Value() (driver.Value, error) {
	if !n.Valid {
		return nil, nil
	}
	return n.Int64, nil
}

func (n *NullInt64) Scan(v interface{}) error {
	return (*sql.NullInt64)(n).Scan(v)
}

func (n *NullInt64) UnmarshalJSON(v []byte) error {
	strV := string(v)
	if strV == "null" {
		n.Valid = false
		return nil
	}
	parsedInt, err := strconv.ParseInt(strV, 10, 64)
	if err != nil {
		n.Valid = false
		return err
	}
	n.Int64, n.Valid = parsedInt, true
	return nil
}

func (n NullInt64) MarshalJSON() ([]byte, error) {
	if !n.Valid {
		return json.Marshal(nil)
	}
	return json.Marshal(n.Int64)
}

func (n NullString) Value() (driver.Value, error) {
	if !n.Valid {
		return nil, nil
	}
	return n.String, nil
}

func (n *NullString) Scan(v interface{}) error {
	return (*sql.NullString)(n).Scan(v)
}

func (n *NullString) UnmarshalJSON(v []byte) error {
	strV := string(v)
	if strV == "null" {
		n.Valid = false
		return nil
	}
	n.String, n.Valid = strV, true
	return nil
}

func (n NullString) MarshalJSON() ([]byte, error) {
	if !n.Valid {
		return json.Marshal(nil)
	}
	return json.Marshal(n.String)
}

func jsonArrayString(v interface{}) (driver.Value, error) {
	r, err := json.Marshal(v)
	str := string(r)
	if err != nil || str == "null" || str == "{}" {
		r = []byte("[]")
	}
	return r, nil
}

func jsonParseForMap(v []byte, dst interface{}) error {
	strV := string(v)
	if strV == "[]" || strV == "{}" || strV == "null" {
		return nil
	}
	return json.Unmarshal(v, dst)
}
