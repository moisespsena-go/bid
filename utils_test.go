package keyb

import (
	"reflect"
	"testing"
)

func Test_readMachineId(t *testing.T) {
	tests := []struct {
		name string
		want []byte
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := readMachineId(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("readMachineId() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_readRandomUint32(t *testing.T) {
	tests := []struct {
		name string
		want uint32
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := readRandomUint32(); got != tt.want {
				t.Errorf("readRandomUint32() = %v, want %v", got, tt.want)
			}
		})
	}
}
