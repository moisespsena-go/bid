package bid

import (
	"bytes"
	"os"
	"testing"
	"time"
)

func TestKey_String(t *testing.T) {
	want := `XrRPlHwdyRQ2ZzYP`
	s := BID{94, 180, 79, 148, 124, 29, 201, 20, 54, 103, 54, 15}.String()
	if s != want {
		t.Errorf("BID.String() = %v, want %v", s, want)
	}
}

func TestKey_Hex(t *testing.T) {
	want := `5eb44f947c1dc9143667360f`
	s := BID{94, 180, 79, 148, 124, 29, 201, 20, 54, 103, 54, 15}.Hex()
	if s != want {
		t.Errorf("BID.MarshalJSON() = %v, want %v", s, want)
	}
}

func TestKey_MarshalJSON(t *testing.T) {
	want := `"XrRPlHwdyRQ2ZzYP"`
	b, _ := BID{94, 180, 79, 148, 124, 29, 201, 20, 54, 103, 54, 15}.MarshalJSON()
	if string(b) != want {
		t.Errorf("BID.MarshalJSON() = %v, want %v", b, want)
	}
}

func TestKey_UnmarshalJSON(t *testing.T) {
	tests := []struct {
		name    string
		arg     string
		want    BID
		wantErr bool
	}{
		{"ok", `"XrRPlHwdyRQ2ZzYP"`, BID{94, 180, 79, 148, 124, 29, 201, 20, 54, 103, 54, 15}, false},
		{"failed", `"XrRPlHwdyRQ2ZzYP`, nil, true},
		{"bad value", `"XrRPlRQ2ZzYP"`, nil, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var this BID
			err := this.UnmarshalJSON([]byte(tt.arg))
			if err != nil {
				if tt.wantErr {
					return
				}
				t.Errorf("BID.UnmarshalJSON() = %v, unexpected error", err)
			} else if tt.wantErr {
				t.Errorf("BID.UnmarshalJSON() = nil, expected error")
			} else if bytes.Compare(this, tt.want) != 0 {
				t.Errorf("BID.UnmarshalJSON() = %v, want %v", this, tt.want)
			}
		})
	}
}

func TestKey_MarshalText(t *testing.T) {
	bid := BID{94, 180, 79, 148, 124, 29, 201, 20, 54, 103, 54, 15}
	b, _ := bid.MarshalText()
	if bytes.Compare(b, bid) != 0 {
		t.Errorf("BID.MarshalText() = %v, want %v", b, bid)
	}
}

func TestKey_UnmarshalText(t *testing.T) {
	b := []byte{94, 180, 79, 148, 124, 29, 201, 20, 54, 103, 54, 15}
	bid := make(BID, 12)
	bid.UnmarshalText(b)
	if bytes.Compare(b, bid) != 0 {
		t.Errorf("BID.UnmarshalText() = %v, want %v", bid, b)
	}
}

func TestKey_Valid(t *testing.T) {
	tests := []struct {
		name string
		this BID
		want bool
	}{
		{"valid", BID{94, 180, 79, 148, 124, 29, 201, 20, 54, 103, 54, 15}, true},
		{"invalid", BID{}, false},
		{"invalid", BID{94, 180, 79, 148, 124, 29, 201, 20, 54, 103, 54}, false},
		{"invalid", BID{94, 180, 79, 148, 124, 29, 201, 20, 54, 103, 54, 15, 17}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.this.Valid(); got != tt.want {
				t.Errorf("BID.Valid() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestKey_Time(t *testing.T) {
	now := time.Unix(time.Now().Unix(), 0).UTC()
	if got := NewWithTime(now).Time(); got != now {
		t.Errorf("BID.Time() = %v, want %v", got, now)
	}
}

func TestKey_Machine(t *testing.T) {
	if got := New().Machine(); bytes.Compare(got, machineId[:]) != 0 {
		t.Errorf("BID.Machine() = %v, want %v", got, machineId)
	}
}

func TestKey_Pid(t *testing.T) {
	pid := uint16(os.Getpid())
	if got := New().Pid(); got != pid {
		t.Errorf("BID.Pid() = %v, want %v", got, pid)
	}
}

func TestKey_Counter(t *testing.T) {
	counter := uint32(678)
	if got := NewArgs(time.Now(), 999, [3]byte{'a', 'b', 'c'}, counter).Counter(); got != counter {
		t.Errorf("BID.Counter() = %v, want %v", got, counter)
	}
}

func TestBID_Generate(t *testing.T) {
	var bid BID
	bid.Generate()
	if len(bid) != 12 {
		t.Errorf("BID.Generate() = %v (%d), want 12 bytes", bid, len(bid))
	}
}