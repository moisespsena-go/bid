package bid

import (
	"bytes"
	"database/sql/driver"
	"encoding/base64"
	"encoding/binary"
	"encoding/hex"
	"fmt"
	"time"

	"github.com/moisespsena-go/logging"
	path_helpers "github.com/moisespsena-go/path-helpers"

	"github.com/pkg/errors"
)

var (
	log    = logging.GetOrCreateLogger(path_helpers.GetCalledDir())
	b64enc = base64.RawURLEncoding
	zero   = make([]byte, 12)
)

// BID is a unique BID identifying. It must be exactly 12 bytes
// long.
type BID []byte

func (this BID) Reset() {
	this = nil
}

func (this BID) Eq(other BID) bool {
	return bytes.Compare(this, other) == 0
}

func (this *BID) Generate() {
	*this = New()
}

func (this *BID) Scan(src interface{}) error {
	*this = []byte{}
	if src == nil {
		return nil
	}
	switch t := src.(type) {
	case []byte:
		return this.ParseBytes(t)
	case string:
		return this.ParseString(t)
	default:
		return errors.New("bad source type")
	}
}

func (this BID) IsZero() bool {
	if this == nil || len(this) == 0 {
		return true
	}
	return bytes.Compare(this[:], zero) == 0
}

func (this BID) Value() (driver.Value, error) {
	if this.IsZero() {
		return "", nil
	}
	return this[:], nil
}

// String returns a base64 string representation of the id.
func (this BID) String() string {
	if this.IsZero() {
		return ""
	}
	return this.B64()
}

// Hex returns a hex representation of the BID.
func (this BID) Hex() string {
	return hex.EncodeToString(this[:])
}

// B64 returns a base 64 representation of the BID.
func (this BID) B64() string {
	return b64enc.EncodeToString(this[:])
}

// AsBytes returns a byte slice.
func (this BID) AsBytes() []byte {
	return this[:]
}

// Bytes returns a byte slice copy.
func (this BID) Bytes() (b []byte) {
	b = make([]byte, 12)
	copy(b, this[:])
	return
}

// MarshalJSON turns a bson.BID into a json.Marshaller.
func (this BID) MarshalJSON() ([]byte, error) {
	return []byte(`"` + this.B64() + `"`), nil
}

var nullBytes = []byte("null")

// UnmarshalJSON turns *bson.BID into a json.Unmarshaller.
func (this *BID) UnmarshalJSON(data []byte) error {
	if len(data) == 2 && data[0] == '"' && data[1] == '"' || bytes.Equal(data, nullBytes) {
		*this = nil
		return nil
	}
	if data[0] != '"' || data[len(data)-1] != '"' {
		return errors.New(fmt.Sprintf("invalid BID in JSON: %s", string(data)))
	}
	return this.ParseString(string(data[1 : len(data)-1]))
}

// MarshalText turns bson.BID into an encoding.TextMarshaler.
func (this BID) MarshalText() ([]byte, error) {
	return this, nil
}

// UnmarshalText turns *BID into an encoding.TextUnmarshaler.
func (this *BID) UnmarshalText(data []byte) error {
	*this = data
	return nil
}

// ParseString turns *BID from string.
func (this *BID) ParseString(s string) error {
	if s == "" {
		*this = nil
		return nil
	}
	if b, err := b64enc.DecodeString(s); err != nil {
		return errors.Wrapf(err, "invalid BID %q in base64", s)
	} else if len(b) != 12 {
		return fmt.Errorf("BID %q in base64 is not 12 bytes", s)
	} else {
		*this = b
		return nil
	}
}

// ParseBytes turns *BID from byte slice.
func (this *BID) ParseBytes(data []byte) error {
	if len(data) == 12 {
		*this = data
		return nil
	}
	if len(data) == 1 && data[0] == ' ' || len(data) == 0 || bytes.Compare(data, zero) == 0 {
		*this = []byte{}
		return nil
	}
	if data[0] == '0' {
		if b, err := b64enc.DecodeString(string(data[1:])); err != nil {
			return errors.Wrapf(err, "invalid BID %q in base64", string(data))
		} else if len(b) != 12 {
			return errors.Wrapf(err, "BID %q in base64 is not 12 bytes", string(data))
		} else {
			*this = b
			return nil
		}
	}
	if len(data) != 24 {
		return fmt.Errorf("invalid BID: %s", data)
	}
	var buf [12]byte
	_, err := hex.Decode(buf[:], data[:])
	if err != nil {
		return fmt.Errorf("invalid BID: %s (%s)", data, err)
	}
	*this = buf[:]
	return nil
}

// Valid returns true if id is valid. A valid id must contain exactly 12 bytes.
func (this BID) Valid() bool {
	return len(this) == 12
}

// Time returns the timestamp part of the id.
// It's a runtime error to call this method with an invalid id.
func (this BID) Time() time.Time {
	// First 4 bytes of BID is 32-bit big-endian seconds from epoch.
	secs := int64(binary.BigEndian.Uint32(this[:][0:4]))
	return time.Unix(secs, 0).UTC()
}

// Machine returns the 3-byte machine id part of the id.
// It's a runtime error to call this method with an invalid id.
func (this BID) Machine() []byte {
	return this[:][4:7]
}

// Pid returns the process id part of the id.
// It's a runtime error to call this method with an invalid id.
func (this BID) Pid() uint16 {
	return binary.BigEndian.Uint16(this[:][7:9])
}

// Counter returns the incrementing value part of the id.
// It's a runtime error to call this method with an invalid id.
func (this BID) Counter() uint32 {
	b := this[:][9:12]
	// Counter is stored as big-endian 3-byte value
	return uint32(b[0])<<16 | uint32(b[1])<<8 | uint32(b[2])
}
