// Package crypto represents algorithms for encrypting and decrypting data
package crypto

import (
	"bytes"
	"testing"
)

// TestBase58_Encode test for structure Encode function
func TestBase58_Encode(t *testing.T) {
	var td = map[string][]byte{
		"":                              []byte(""),
		"0":                             []byte("1"),
		"32":                            []byte("Z"),
		"64":                            []byte("27"),
		"000":                           []byte("111"),
		"512":                           []byte("9q"),
		"1024":                          []byte("Jf"),
		"16777216":                      []byte("2UzHM"),
		"00068719476736":                []byte("1112ohWHHR"),
		"18446744073709551616":          []byte("jpXCZedGfVR"),
		"79228162514264337593543950336": []byte("5qCHTcgbQwpvYZQ9d"),
	}
	bs68 := BitcoinBase58

	for dat, res := range td {
		dat := dat
		res := res
		t.Run(dat, func(t *testing.T) {
			t.Parallel()
			tres, err := bs68.Encode([]byte(dat))
			if err != nil {
				t.Errorf("Error occurred while decoding %s.", dat)
			}
			if !bytes.Equal(res, tres) {
				t.Error("Expected ", res, ", got ", tres)
			}
		})
	}
}

//TestBase58_Decode test for Base58 structure Decode function
func TestBase58_Decode(t *testing.T) {
	ts := [][]byte{
		[]byte(""),
		[]byte("0"),
		[]byte("32"),
		[]byte("64"),
		[]byte("000"),
		[]byte("512"),
		[]byte("1024"),
		[]byte("16777216"),
		[]byte("00068719476736"),
		[]byte("18446744073709551616"),
		[]byte("79228162514264337593543950336"),
	}
	bs68 := BitcoinBase58

	for _, dat := range ts {
		dat := dat
		t.Run(string(dat), func(t *testing.T) {
			t.Parallel()
			ets, eerr := bs68.Encode(dat)
			if eerr != nil {
				t.Errorf("Error occurred while encoding %s.", dat)
			}
			dets, derr := bs68.Decode(ets)
			if derr != nil {
				t.Errorf("Error occurred while decoding %s.", ets)
			}
			if !bytes.Equal(dat, dets) {
				t.Error("Expected", dat, " got ", dets)
			}
		})
	}
}
