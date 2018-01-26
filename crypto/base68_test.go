package crypto

import (
	"bytes"
	"testing"
)

func TestBase58Encode(t *testing.T) {
	ts := []byte("Test string")
	tsenc := []byte{49, 77, 118, 113, 76, 110, 90, 85, 71, 85, 103, 78, 98, 68, 120, 50}
	ets := Base58Encode(ts)
	if !bytes.Equal(tsenc, ets) {
		t.Error("Expected: ", tsenc, " got ", ets)
	}
}

func TestBase58Decode(t *testing.T) {
	ts := []byte("Test string")
	ets := Base58Encode(ts)
	dets := Base58Decode(ets)
	if !bytes.Equal(ts, dets) {
		t.Error("Expected", ts, " got ", dets)
	}
}
