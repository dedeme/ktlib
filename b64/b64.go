// Copyright 25-May-2020 ºDeme
// GNU General Public License - V3 <http://www.gnu.org/licenses/>

// Base 64 functions.
package b64

import (
	"encoding/base64"
)

// Returns the b64 representation of 's'
func Encode(s string) string {
	return base64.StdEncoding.EncodeToString([]byte(s))
}

// Returns the b64 representation of 'bs'
func EncodeBytes(bs []byte) string {
	return base64.StdEncoding.EncodeToString(bs)
}

// Restore the string encoded with Encode.
func Decode(s string) string {
	bs, err := base64.StdEncoding.DecodeString(s)
	if err != nil {
		panic(err)
	}
	return string(bs)
}

// Restore the 'bytes' encoded with EncodeBytes
func DecodeBytes(s string) []byte {
	bs, err := base64.StdEncoding.DecodeString(s)
	if err != nil {
		panic(err)
	}
	return bs
}
