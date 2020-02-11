package bid

import (
	"crypto/md5"
	"os"
	"time"
)

// readMachineId generates and returns a machine id.
// If this function fails to get the hostname it will cause a runtime error.
func readMachineId() []byte {
	var sum [3]byte
	id := sum[:]
	hostname, err1 := os.Hostname()
	if err1 != nil {
		n := uint32(time.Now().UnixNano())
		sum[0] = byte(n >> 0)
		sum[1] = byte(n >> 8)
		sum[2] = byte(n >> 16)
		return id
	}
	hw := md5.New()
	hw.Write([]byte(hostname))
	copy(id, hw.Sum(nil))
	return id
}
// readMachineId generates and returns a machine id.
// If this function fails to get the hostname it will cause a runtime error.
func MachineId(hostname string) []byte {
	var sum [3]byte
	id := sum[:]
	hw := md5.New()
	hw.Write([]byte(hostname))
	copy(id, hw.Sum(nil))
	return id
}

// readRandomUint32 returns a random objectIdCounter.
func readRandomUint32() uint32 {
	// We've found systems hanging in this function due to lack of entropy.
	// The randomness of these bytes is just preventing nearby clashes, so
	// just look at the time.
	return uint32(time.Now().UnixNano())
}
