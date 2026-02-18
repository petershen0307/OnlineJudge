package addbinary

import (
	"slices"
)

func addBinary(a string, b string) string {
	aa := toBigEndian(a)
	bb := toBigEndian(b)
	var cc []byte
	if len(aa) > len(bb) {
		cc = add(bb, aa)
	} else {
		cc = add(aa, bb)
	}
	// to little endian
	slices.Reverse(cc)
	// byte to rune
	r := []rune{}
	for _, c := range cc {
		if c == byte(1) {
			r = append(r, '1')
		} else {
			r = append(r, '0')
		}
	}
	return string(r)
}

func toBigEndian(s string) []byte {
	r := []byte{}
	ts := []rune(s)
	slices.Reverse(ts)
	for _, t := range ts {
		if t == '1' {
			r = append(r, 1)
		} else {
			r = append(r, 0)
		}
	}
	return r
}

func add(aa, bb []byte) []byte {
	// aa length always less than bb
	cc := []byte{}
	carry := byte(0)
	for i := 0; i < len(bb); i++ {
		c := bb[i] ^ carry
		if i < len(aa) {
			c = aa[i] ^ c
			carry = (aa[i] + bb[i] + carry) / 2
		} else {
			carry = (bb[i] + carry) / 2
		}
		cc = append(cc, c)
	}
	if carry == byte(1) {
		cc = append(cc, 1)
	}
	return cc
}
