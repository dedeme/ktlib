// Copyright 21-Jul-2022 ºDeme
// GNU General Public License - V3 <http://www.gnu.org/licenses/>

package _tests

import (
	"github.com/dedeme/ktlib/sbf"
	"testing"
)

func TestSbf(t *testing.T) {
	eq(t, sbf.New().Len(), 0)
	eq(t, sbf.New().String(), "")
	eq(t, sbf.NewStr("abc").Len(), 3)
	eq(t, sbf.NewStr("abc").String(), "abc")
	sb := sbf.New()
	sb.AddByte('r')
	eq(t, sb.String(), "r")
	sb.Add("st")
	eq(t, sb.String(), "rst")
	sb.Reset()
	eq(t, sb.String(), "")
}
