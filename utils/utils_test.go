package utils

import (
	"bytes"
	"testing"
)

func TestIntToHex(t *testing.T) {
	a := IntToHex(123243221)
	b := []byte{0, 0, 0, 0, 7, 88, 138, 213}
	if !bytes.Equal(a, b) {
		t.Error("Expected [0 0 0 0 7 88 138 213], got", a)
	}
}

func TestReverseBytes(t *testing.T) {
	b := []byte{1, 2, 3}
	a := []byte{3, 2, 1}
	if !bytes.Equal(a, ReverseBytes(b)) {
		t.Error("Expected [3 2 1], got", a)
	}
}
