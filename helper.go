package bid

import (
	"encoding/binary"
	"encoding/hex"
	"fmt"
	"os"
	"sync/atomic"
	"time"
)

// machineId stores machine id generated once and used in subsequent calls
// to NewObjectId function.
var machineId = readMachineId()
var processId = os.Getpid()

// keyCounter is atomically incremented when generating a new BID
// using New() function. It's used as a counter part of an id.
var keyCounter = readRandomUint32()

func New() BID {
	return NewWithTime(time.Now())
}

// Hex returns an BID from the provided hex representation.
// Calling this function with an invalid hex representation will
// cause a runtime panic. See the IsHex function.
func Hex(s string) (bid BID) {
	d, err := hex.DecodeString(s)
	if err != nil || len(d) != 12 {
		panic(fmt.Sprintf("invalid input to Hex: %q", s))
	}
	copy(bid[:], d)
	return
}

// IsHex returns whether s is a valid hex representation of
// an BID. See the Hex function.
func IsHex(s string) bool {
	if len(s) != 24 {
		return false
	}
	_, err := hex.DecodeString(s)
	return err == nil
}

// B64 returns an BID from the provided base64 representation.
// Calling this function with an invalid hex representation will
// cause a runtime panic. See the IsObjectIdHex function.
func B64(s string) (bid BID) {
	d, err := b64enc.DecodeString(s)
	if err != nil || len(d) != 12 {
		panic(fmt.Sprintf("invalid input to Hex: %q", s))
	}
	copy(bid[:], d)
	return
}

// IsB64 returns whether s is a valid base64 representation of
// an BID. See the Hex function.
func IsB64(s string) bool {
	if s[0] != '0' {
		return false
	}
	_, err := b64enc.DecodeString(s)
	return err == nil
}

// NewKeyWithTime returns a dummy BID with the timestamp part filled
// with the provided number of seconds from epoch UTC, and all other parts
// filled with zeroes. It's not safe to insert a document with an id generated
// by this method, it is useful only for queries to find documents with ids
// generated before or after the specified timestamp.
func NewKeyWithTime(t time.Time) (bid BID) {
	binary.BigEndian.PutUint32(bid[:4], uint32(t.Unix()))
	return
}

// NewWithTime returns a BID with the timestamp part filled
// with the provided number of seconds from epoch UTC.
func NewWithTime(t time.Time) BID {
	return NewArgs(t, processId, machineId, atomic.AddUint32(&keyCounter, 1))
}

// NewWithTimeArgs returns a BID with the timestamp args (year, month, day, hour, minute, second).
func NewWithTimeArgs(Y, M, D, h, m, s int) BID {
	return NewWithTime(time.Date(Y, time.Month(M), D, h, m, s, 0, time.UTC))
}

// NewArgs create the BID object with args
func NewArgs(now time.Time, processId int, machineId [3]byte, counter uint32) (b BID) {
	b = make(BID, 12, 12)
	// Timestamp, 4 bytes, big endian
	binary.BigEndian.PutUint32(b[:], uint32(now.UTC().Unix()))
	// Machine, first 3 bytes of md5(hostname)
	b[4] = machineId[0]
	b[5] = machineId[1]
	b[6] = machineId[2]
	// Pid, 2 bytes, specs don't specify endianness, but we use big endian.
	b[7] = byte(processId >> 8)
	b[8] = byte(processId)
	// Increment, 3 bytes, big endian
	b[9] = byte(counter >> 16)
	b[10] = byte(counter >> 8)
	b[11] = byte(counter)
	return
}

func Parse(b []byte) (bid BID, err error) {
	err = bid.ParseBytes(b)
	return
}

func MustParse(b []byte) (bid BID) {
	if err := bid.ParseBytes(b); err != nil {
		log.Fatal(err)
	}
	return
}

func ParseString(s string) (bid BID, err error) {
	err = bid.ParseBytes([]byte(s))
	return
}

func MustParseString(s string) (bid BID) {
	if err := bid.ParseString(s); err != nil {
		log.Fatal(err)
	}
	return
}

func From(value interface{}) (bid BID) {
	if value == nil {
		return
	}
	switch t := value.(type) {
	case BID:
		return t
	case string:
		if err := bid.ParseBytes([]byte(t)); err != nil {
			log.Fatal("From %v failed: %s", value, err.Error())
		}
		return
	case []byte:
		if err := bid.ParseBytes(t); err != nil {
			log.Fatal("From %v failed: %s", value, err.Error())
		}
		return
	case interface{ GetBID() BID }:
		return t.GetBID()
	default:
		log.Fatal("From %v: bad value type", value)
		return
	}
}
