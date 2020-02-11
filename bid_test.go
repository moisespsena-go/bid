package bid

import (
	"bytes"
	"reflect"
	"testing"
	"time"
)

func TestKey_String(t *testing.T) {
	tests := []struct {
		name string
		this BID
		want string
	}{
		{"t1", MustParse("5cd9a8767c1dc9687b139fd7"), "5cd9a8767c1dc9687b139fd7"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.this.String(); got != tt.want {
				t.Errorf("BID.String() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestKey_Hex(t *testing.T) {
	tests := []struct {
		name string
		this BID
		want string
	}{
		{"t1", MustParse("5cd9a8767c1dc9687b139fd7"), "5cd9a8767c1dc9687b139fd7"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.this.Hex(); got != tt.want {
				t.Errorf("BID.Hex() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestKey_MarshalJSON(t *testing.T) {
	tests := []struct {
		name    string
		this    BID
		want    []byte
		wantErr bool
	}{
		{"t1", MustParse("5cd9a8767c1dc9687b139fd7"), []byte(`"5cd9a8767c1dc9687b139fd7"`), false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.this.MarshalJSON()
			if (err != nil) != tt.wantErr {
				t.Errorf("BID.MarshalJSON() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("BID.MarshalJSON() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestKey_UnmarshalJSON(t *testing.T) {
	type args struct {
		data []byte
	}
	tests := []struct {
		name    string
		this    *BID
		args    args
		wantErr bool
	}{
		{"t1", new(BID), args{[]byte(`"5cd9a8767c1dc9687b139fd7"`)}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.this.UnmarshalJSON(tt.args.data); (err != nil) != tt.wantErr {
				t.Errorf("BID.UnmarshalJSON() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got, _ := tt.this.MarshalJSON(); bytes.Compare(got, tt.args.data) != 0 {
				t.Errorf("BID.MarshalJSON() = %v, want %v", got, tt.args.data)
				return
			}
		})
	}
}

func TestKey_MarshalText(t *testing.T) {
	tests := []struct {
		name    string
		this    BID
		want    []byte
		wantErr bool
	}{
		{"t1", MustParse("5cd9a8767c1dc9687b139fd7"), []byte("5cd9a8767c1dc9687b139fd7"), false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.this.MarshalText()
			if (err != nil) != tt.wantErr {
				t.Errorf("BID.MarshalText() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("BID.MarshalText() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestKey_UnmarshalText(t *testing.T) {
	type args struct {
		data []byte
	}
	tests := []struct {
		name    string
		this    *BID
		args    args
		wantErr bool
	}{
		{"t1", new(BID), args{[]byte("5cd9a8767c1dc9687b139fd7")}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.this.ParseBytes(tt.args.data); (err != nil) != tt.wantErr {
				t.Errorf("BID.UnmarshalText() error = %v, wantErr %v", err, tt.wantErr)
			}
			if got := []byte(tt.this.Hex()); bytes.Compare(got, tt.args.data) != 0 {
				t.Errorf("BID.UnmarshalText() = %v, want %v", got, tt.args.data)
				return
			}
		})
	}
}

func TestKey_Valid(t *testing.T) {
	tests := []struct {
		name string
		this BID
		want bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.this.Valid(); got != tt.want {
				t.Errorf("BID.Valid() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestKey_byteSlice(t *testing.T) {
	type args struct {
		start int
		end   int
	}
	tests := []struct {
		name string
		this BID
		args args
		want []byte
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.this.byteSlice(tt.args.start, tt.args.end); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("BID.byteSlice() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestKey_Time(t *testing.T) {
	tests := []struct {
		name string
		this BID
		want time.Time
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.this.Time(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("BID.Time() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestKey_Machine(t *testing.T) {
	tests := []struct {
		name string
		this BID
		want []byte
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.this.Machine(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("BID.Machine() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestKey_Pid(t *testing.T) {
	tests := []struct {
		name string
		this BID
		want uint16
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.this.Pid(); got != tt.want {
				t.Errorf("BID.Pid() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestKey_Counter(t *testing.T) {
	tests := []struct {
		name string
		this BID
		want int32
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.this.Counter(); got != tt.want {
				t.Errorf("BID.Counter() = %v, want %v", got, tt.want)
			}
		})
	}
}
