package keyb

import (
	"reflect"
	"testing"
	"time"
)

func TestNew(t *testing.T) {
	tests := []struct {
		name string
		want Key
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := New(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("New() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestKeyHex(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want Key
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := KeyHex(tt.args.s); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("KeyHex() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsKeyHex(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsKeyHex(tt.args.s); got != tt.want {
				t.Errorf("IsKeyHex() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewKeyWithTime(t *testing.T) {
	type args struct {
		t time.Time
	}
	tests := []struct {
		name string
		args args
		want Key
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewKeyWithTime(tt.args.t); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewKeyWithTime() = %v, want %v", got, tt.want)
			}
		})
	}
}
