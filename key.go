package keyb

import (
	"bytes"
	"database/sql/driver"
	"encoding/binary"
	"encoding/hex"
	"errors"
	"fmt"
	"time"
)

// Key is a unique ID identifying a BSON value. It must be exactly 12 bytes
// long. MongoDB objects by default have such a property set in their "_id"
// property.
//
// http://www.mongodb.org/display/DOCS/Object+IDs
type Key string

func (this *Key) Scan(src interface{}) error {
	*this = ""
	if src == nil {
		return nil
	}
	switch t := src.(type) {
	case []byte:
		switch len(t) {
		case 12:
			*this = Key(string(t))
		default:
			return this.UnmarshalText(t)
		}
		return nil
	case string:
		switch len(t) {
		case 12:
			*this = Key(t)
		default:
			return this.UnmarshalText([]byte(t))
		}
		return nil
	default:
		return errors.New("bad source type")
	}
}

func (this Key) Value() (driver.Value, error) {
	return []byte(this), nil
}

// String returns a hex string representation of the id.
// Example: KeyHex("4d88e15b60f486e428412dc9").
func (this Key) String() string {
	return this.Hex()
}

// Hex returns a hex representation of the Key.
func (this Key) Hex() string {
	return hex.EncodeToString([]byte(this))
}

// MarshalJSON turns a bson.Key into a json.Marshaller.
func (this Key) MarshalJSON() ([]byte, error) {
	return []byte(`"` + this.Hex() + `"`), nil
}

var nullBytes = []byte("null")

// UnmarshalJSON turns *bson.Key into a json.Unmarshaller.
func (this *Key) UnmarshalJSON(data []byte) error {
	if len(data) == 2 && data[0] == '"' && data[1] == '"' || bytes.Equal(data, nullBytes) {
		*this = ""
		return nil
	}
	if len(data) != 26 || data[0] != '"' || data[25] != '"' {
		return errors.New(fmt.Sprintf("invalid Key in JSON: %s", string(data)))
	}
	return this.UnmarshalText(data[1:25])
}

// MarshalText turns bson.Key into an encoding.TextMarshaler.
func (this Key) MarshalText() ([]byte, error) {
	return []byte(string(this)), nil
}

// UnmarshalText turns *bson.Key into an encoding.TextUnmarshaler.
func (this *Key) UnmarshalText(data []byte) error {
	if len(data) == 1 && data[0] == ' ' || len(data) == 0 {
		*this = ""
		return nil
	}
	if len(data) != 24 {
		return fmt.Errorf("invalid Key: %s", data)
	}
	var buf [12]byte
	_, err := hex.Decode(buf[:], data[:])
	if err != nil {
		return fmt.Errorf("invalid Key: %s (%s)", data, err)
	}
	*this = Key(string(buf[:]))
	return nil
}

// Valid returns true if id is valid. A valid id must contain exactly 12 bytes.
func (this Key) Valid() bool {
	return len(this) == 12
}

// byteSlice returns byte slice of id from start to end.
// Calling this function with an invalid id will cause a runtime panic.
func (this Key) byteSlice(start, end int) []byte {
	if len(this) != 12 {
		panic(fmt.Sprintf("invalid Key: %q", string(this)))
	}
	return []byte(string(this)[start:end])
}

// Time returns the timestamp part of the id.
// It's a runtime error to call this method with an invalid id.
func (this Key) Time() time.Time {
	// First 4 bytes of Key is 32-bit big-endian seconds from epoch.
	secs := int64(binary.BigEndian.Uint32(this.byteSlice(0, 4)))
	return time.Unix(secs, 0)
}

// Machine returns the 3-byte machine id part of the id.
// It's a runtime error to call this method with an invalid id.
func (this Key) Machine() []byte {
	return this.byteSlice(4, 7)
}

// Pid returns the process id part of the id.
// It's a runtime error to call this method with an invalid id.
func (this Key) Pid() uint16 {
	return binary.BigEndian.Uint16(this.byteSlice(7, 9))
}

// Counter returns the incrementing value part of the id.
// It's a runtime error to call this method with an invalid id.
func (this Key) Counter() int32 {
	b := this.byteSlice(9, 12)
	// Counter is stored as big-endian 3-byte value
	return int32(uint32(b[0])<<16 | uint32(b[1])<<8 | uint32(b[2]))
}
