// Copyright 25-May-2017 ºDeme
// GNU General Public License - V3 <http://www.gnu.org/licenses/>

// Cryptographic utilities
package cryp

import (
	"crypto/rand"
	"encoding/base64"
)

var b64 = base64.StdEncoding

// GenK generates a B64 randon key of length 'lg'
func GenK(lg int) string {
	arr := make([]byte, lg)
	_, err := rand.Read(arr)
	if err != nil {
		panic(err)
	}
	return b64.EncodeToString(arr)[:lg]
}

// Returns 'key' codified in irreversible way, using 'lg' B64 digits.
//   key   : String to codify
//   lg    : Length of result
//   return: 'lg' B64 digits
func Key(key string, lg int) string {
	k := []byte(
		key + "codified in irreversibleDeme is good, very good!\n\r8@@")

	lenk := len(k)
	sum := 0
	for i := 0; i < lenk; i++ {
		sum += int(k[i])
	}

	lg2 := lg + lenk
	r := make([]byte, lg2)
	r1 := make([]byte, lg2)
	r2 := make([]byte, lg2)
	ik := 0
	for i := 0; i < lg2; i++ {
		v1 := int(k[ik])
		v2 := v1 + int(k[v1%lenk])
		v3 := v2 + int(k[v2%lenk])
		v4 := v3 + int(k[v3%lenk])
		sum = (sum + i + v4) & 255
		r1[i] = byte(sum)
		r2[i] = byte(sum)
		ik++
		if ik == lenk {
			ik = 0
		}
	}

	for i := 0; i < lg2; i++ {
		v1 := int(r2[i])
		v2 := v1 + int(r2[v1%lg2])
		v3 := v2 + int(r2[v2%lg2])
		v4 := v3 + int(r2[v3%lg2])
		sum = (sum + v4) & 255
		r2[i] = byte(sum)
		r[i] = byte((sum + int(r1[i])) & 255)
	}

	return b64.EncodeToString(r)[:lg]
}

// Encodes 'msg' with key 'key'.
//   key   : Key for encoding
//   msg   : Message to encode
//   return: 'm' codified in B64 digits.
func Encode(key, msg string) string {
	m := b64.EncodeToString([]byte(msg))
	lg := len(m)
	k := Key(key, lg)
	mb := []byte(m)
	kb := []byte(k)
	r := make([]byte, lg)
	for i := 0; i < lg; i++ {
		r[i] = mb[i] + kb[i]
	}
	return b64.EncodeToString(r)
}

// Decodes 'c' using key 'key'. 'c' was codified with Encode().
//   key   : Key for decoding
//   c     : Text codified with Encode()
//   return: 'c' decoded.
func Decode(key, c string) string {
	mb, err := b64.DecodeString(c)
	if err != nil {
		panic(err)
	}
	lg := len(mb)
	k := Key(key, lg)
	kb := []byte(k)
	r := make([]byte, lg)
	for i := 0; i < lg; i++ {
		r[i] = mb[i] - kb[i]
	}
	mb, err = b64.DecodeString(string(r))
	if err != nil {
		panic(err)
	}
	return string(mb)
}
