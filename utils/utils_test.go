// Package utils contains useful functions
package utils

import (
	"bytes"
	"fmt"
	"testing"
)

// TestIntToHex test for IntToHex
func TestIntToHex(t *testing.T) {
	var td = map[int64][]byte{
		123243221: {0, 0, 0, 0, 7, 88, 138, 213},
		12324322:  {0, 0, 0, 0, 0, 188, 13, 226},
	}
	for dat, res := range td {
		dat := dat
		res := res
		t.Run(fmt.Sprintf("%s%d", "Test for data: ", dat), func(t *testing.T) {
			t.Parallel()
			tres := IntToHex(dat)
			if !bytes.Equal(tres, res) {
				t.Error("Expected ", res, ", got ", tres)
			}
		})
	}
}

// TestReverseBytes test for Reverse
func TestReverseBytes(t *testing.T) {
	type PairByte struct {
		dat, res []byte
	}

	var td = map[string]PairByte{
		"Test1": {[]byte{1, 2, 3}, []byte{3, 2, 1}},
		"Test2": {[]byte{1, 2, 3, 4}, []byte{4, 3, 2, 1}},
		"Test3": {[]byte{1, 2, 3, 4, 5}, []byte{5, 4, 3, 2, 1}},
	}

	for n, dat := range td {
		n := n
		dat := dat
		t.Run(n, func(t *testing.T) {
			t.Parallel()
			tres := Reverse(dat.dat)
			if !bytes.Equal(tres, dat.res) {
				t.Error("Expected ", dat.res, ", got ", tres)
			}
		})
	}
}
