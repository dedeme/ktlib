// Copyright 25-May-2020 ºDeme
// GNU General Public License - V3 <http://www.gnu.org/licenses/>

package _tests

import (
	"github.com/dedeme/ktlib/cryp"
	"testing"
)

func TestCryp(t *testing.T) {
	eq(t, len(cryp.GenK(5)), 5)
	k := cryp.Key("abc", 8)
	eq(t, k, "C8vYu4C/")
	code := cryp.Encode(k, "El cañón disparó")
	eq(t, code == "El cañón disparó", false)
	eq(t, cryp.Decode(k, code), "El cañón disparó")

	eq(t, cryp.Encode(k, ""), "")
	eq(t, cryp.Decode(k, ""), "")
}
