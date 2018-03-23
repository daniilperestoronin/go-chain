// Package crypto represents algorithms for encrypting and decrypting data
package crypto

import (
	"fmt"
	"math/big"

	"github.com/daniilperestoronin/go-chain/utils"
)

// Base58 structeure represent base 58 cryptographic algorithm
type Base58 struct {
	alphabet  [58]byte
	decodeMap [256]int64
}

// base58 create Base58 struct with received alphabet
func base58(alphabet []byte) *Base58 {
	bs := &Base58{}
	copy(bs.alphabet[:], alphabet[:])
	for i := range bs.decodeMap {
		bs.decodeMap[i] = -1
	}
	for i, b := range bs.alphabet {
		bs.decodeMap[b] = int64(i)
	}
	return bs
}

// BitcoinBase58 initialize Base58 sructure for Bitcoin with received alphabet
var BitcoinBase58 = base58([]byte("123456789ABCDEFGHJKLMNPQRSTUVWXYZabcdefghijkmnopqrstuvwxyz"))

// radix for base 58
var radix = big.NewInt(58)

// Encode encrypting src using base 58 algorithm
func (base58 *Base58) Encode(src []byte) ([]byte, error) {
	if len(src) == 0 {
		return []byte{}, nil
	}
	n, ok := new(big.Int).SetString(string(src), 10)
	if !ok {
		return nil, fmt.Errorf("expecting a number but got %q", src)
	}
	bytes := make([]byte, 0, len(src))
	for _, c := range src {
		if c == '0' {
			bytes = append(bytes, base58.alphabet[0])
		} else {
			break
		}
	}
	zerocnt := len(bytes)
	mod := new(big.Int)
	zero := big.NewInt(0)
	for {
		switch n.Cmp(zero) {
		case 1:
			n.DivMod(n, radix, mod)
			bytes = append(bytes, base58.alphabet[mod.Int64()])
		case 0:
			utils.Reverse(bytes[zerocnt:])
			return bytes, nil
		default:
			return nil, fmt.Errorf("expecting a positive number in base58 encoding but got %q", n)
		}
	}
}

// Decode decrypting encrypted base 58 src 
func (base58 *Base58) Decode(src []byte) ([]byte, error) {
	if len(src) == 0 {
		return []byte{}, nil
	}
	var zeros []byte
	for i, c := range src {
		if c == base58.alphabet[0] && i < len(src)-1 {
			zeros = append(zeros, '0')
		} else {
			break
		}
	}
	n := new(big.Int)
	var i int64
	for _, c := range src {
		if i = base58.decodeMap[c]; i < 0 {
			return nil, fmt.Errorf("invalid character '%c' in decoding a base58 string \"%s\"", c, src)
		}
		n.Add(n.Mul(n, radix), big.NewInt(i))
	}
	return n.Append(zeros, 10), nil
}
