// Copyright 25-May-2022 ºDeme
// GNU General Public License - V3 <http://www.gnu.org/licenses/>

package _tests

import (
	"github.com/dedeme/ktlib/b64"
	"testing"
)

func TestB64(t *testing.T) {
	eq(t, b64.Decode(b64.Encode("")), "")
	eq(t, len(b64.DecodeBytes(b64.EncodeBytes([]byte{}))), 0)

	eq(t, b64.Decode(b64.Encode("¿Vió un cañón?")), "¿Vió un cañón?")
	eq(t, string(b64.DecodeBytes(b64.EncodeBytes([]byte{1, 255, 0, 11}))),
		string([]byte{1, 255, 0, 11}))
}
