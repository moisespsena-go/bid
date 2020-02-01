package keyb

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

// keyCounter is atomically incremented when generating a new Key
// using NewObjectId() function. It's used as a counter part of an id.
var keyCounter = readRandomUint32()

func New() Key {
	return NewArgs(time.Now(), processId, machineId, atomic.AddUint32(&keyCounter, 1))
}

// KeyHex returns an Key from the provided hex representation.
// Calling this function with an invalid hex representation will
// cause a runtime panic. See the IsObjectIdHex function.
func KeyHex(s string) Key {
	d, err := hex.DecodeString(s)
	if err != nil || len(d) != 12 {
		panic(fmt.Sprintf("invalid input to KeyHex: %q", s))
	}
	return Key(d)
}

// IsObjectIdHex returns whether s is a valid hex representation of
// an Key. See the KeyHex function.
func IsKeyHex(s string) bool {
	if len(s) != 24 {
		return false
	}
	_, err := hex.DecodeString(s)
	return err == nil
}

// NewKeyWithTime returns a dummy Key with the timestamp part filled
// with the provided number of seconds from epoch UTC, and all other parts
// filled with zeroes. It's not safe to insert a document with an id generated
// by this method, it is useful only for queries to find documents with ids
// generated before or after the specified timestamp.
func NewKeyWithTime(t time.Time) Key {
	var b [12]byte
	binary.BigEndian.PutUint32(b[:4], uint32(t.Unix()))
	return Key(string(b[:]))
}

// NewArgs create the Key object with args
func NewArgs(now time.Time, processId int, machineId []byte, counter uint32) Key {
	var b [12]byte
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

	return Key(b[:])
}

func Parse(hex string) (key Key, err error) {
	if len(hex) == 12 {
		key = Key(hex)
		return
	}
	err = key.UnmarshalText([]byte(hex))
	return
}

func MustParse(hex string) (key Key) {
	if len(hex) == 12 {
		return Key(hex)
	}
	if err := key.UnmarshalText([]byte(hex)); err != nil {
		panic(fmt.Errorf("Parse Key failed: %v", err.Error()))
	}
	return
}
