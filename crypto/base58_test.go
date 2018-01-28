package crypto

import (
	"bytes"
	"testing"
)

func TestBase58Encode(t *testing.T) {
	var td = map[string][]byte{
		"Test string":   []byte{49, 77, 118, 113, 76, 110, 90, 85, 71, 85, 103, 78, 98, 68, 120, 50},
		"Test string 2": []byte{49, 56, 50, 105, 80, 55, 57, 71, 84, 122, 102, 67, 65, 118, 71, 118, 51, 57, 51},
	}

	for dat, res := range td {
		dat := dat
		res := res
		t.Run(dat, func(t *testing.T) {
			t.Parallel()
			tres := Base58Encode([]byte(dat))
			if !bytes.Equal(res, tres) {
				t.Error("Expected ", res, ", got ", tres)
			}
		})
	}
}

func TestBase58Decode(t *testing.T) {
	ts := [][]byte{
		[]byte("Test string"),
		[]byte("Test string 2"),
		[]byte("Test string 3"),
	}

	for _, dat := range ts {
		dat := dat
		t.Run(string(dat), func(t *testing.T) {
			t.Parallel()
			ets := Base58Encode(dat)
			dets := Base58Decode(ets)
			if !bytes.Equal(dat, dets) {
				t.Error("Expected", dat, " got ", dets)
			}
		})
	}
}
