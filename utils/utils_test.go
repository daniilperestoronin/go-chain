package utils

import (
	"bytes"
	"fmt"
	"testing"
)

func TestIntToHex(t *testing.T) {
	var td = map[int64][]byte{
		123243221: []byte{0, 0, 0, 0, 7, 88, 138, 213},
		12324322:  []byte{0, 0, 0, 0, 0, 188, 13, 226},
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

func TestReverseBytes(t *testing.T) {
	b := []byte{1, 2, 3}
	a := []byte{3, 2, 1}
	if !bytes.Equal(a, ReverseBytes(b)) {
		t.Error("Expected [3 2 1], got", a)
	}
}
