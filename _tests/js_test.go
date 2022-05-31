// Copyright 25-May-2020 ºDeme
// GNU General Public License - V3 <http://www.gnu.org/licenses/>

package _tests

import (
	"github.com/dedeme/ktlib/js"
	"github.com/dedeme/ktlib/math"
	"testing"
)

func TestJs(t *testing.T) {
	eq(t, js.Wn(), "null")
	eq(t, js.IsNull("null"), true)

	eq(t, js.Wb(true), "true")
	eq(t, js.Wb(false), "false")
	eq(t, js.Rb("true"), true)
	eq(t, !js.Rb("false"), true)

	eq(t, js.Wi(-12), "-12")
	eq(t, js.Wi(-0), "0")
	eq(t, js.Wi(1234), "1234")
	eq(t, js.Ri("-12"), -12)
	eq(t, js.Ri("-0"), 0)
	eq(t, js.Ri("12"), 12)

	eq(t, js.Wf(-12.23), "-12.23")
	eq(t, js.Wf(-12.23e+4), "-122300")
	eq(t, js.Wf(-0.0), "0")
	eq(t, js.Wf(1234.0), "1234")
	eq(t, js.Wf(1234.0e-2), "12.34")
	eq(t, math.Eq(float64(js.Rf("-12.23")), -12.23, 0.000001), true)
	eq(t, math.Eq(float64(js.Rf("-122300")), -12.23e+4, 0.000001), true)
	eq(t, math.Eq(float64(js.Rf("-0")), -0.0, 0.000001), true)
	eq(t, math.Eq(float64(js.Rf("1234")), 1234.0, 0.000001), true)
	eq(t, math.Eq(float64(js.Rf("12.34")), 1234.0e-2, 0.000001), true)

	eq(t, js.Wd(-12.23), "-12.23")
	eq(t, js.Wd(-12.23e+4), "-122300")
	eq(t, js.Wd(-0.0), "0")
	eq(t, js.Wd(1234.0), "1234")
	eq(t, js.Wd(1234.0e-2), "12.34")
	eq(t, math.Eq(js.Rd("-12.23"), -12.23, 0.00000001), true)
	eq(t, math.Eq(js.Rd("-122300"), -12.23e+4, 0.00000001), true)
	eq(t, math.Eq(js.Rd("-0"), -0.0, 0.00000001), true)
	eq(t, math.Eq(js.Rd("1234"), 1234.0, 0.00000001), true)
	eq(t, math.Eq(js.Rd("12.34"), 1234.0e-2, 0.00000001), true)

	eq(t, js.WdDec(-12.23, 1), "-12.2")
	eq(t, js.WdDec(-12.23e+4, 1), "-122300.0")
	eq(t, js.WdDec(-0.0, 1), "0.0")
	eq(t, js.WdDec(1234.0, 1), "1234.0")
	eq(t, js.WdDec(1235.0e-2, 1), "12.4")

	eq(t, js.Ws(""), "\"\"")
	eq(t, js.Rs("\"\""), "")
	eq(t, js.Ws("¿Qué \"suerte\" cogió?"), "\"¿Qué \\\"suerte\\\" cogió?\"")
	eq(t, js.Rs("\"¿Qué \\\"suerte\\\" cogió?\""), "¿Qué \"suerte\" cogió?")

	eq(t, js.Wa([]string{}), "[]")
	eq(t, len(js.Ra("[]")), 0)
	eq(t, js.Wa([]string{js.Wi(1), js.Wb(true), js.Ws("a")}), "[1,true,\"a\"]")
	ajs := js.Ra("[1,true,\"a\"]")
	eq(t, len(ajs), 3)
	eq(t, js.Ri(ajs[0]), 1)
	eq(t, js.Rb(ajs[1]), true)
	eq(t, js.Rs(ajs[2]), "a")

	eq(t, js.Wo(map[string]string{}), "{}")
	eq(t, len(js.Ro("{}")), 0)
	m0 := js.Ro(js.Wo(map[string]string{
		"one":   js.Wi(-1),
		"two":   js.Wb(false),
		"three": js.Ws(""),
		"four":  js.Wa([]string{js.Wi(1)}),
	}))
	eq(t, js.Ri(m0["one"]), -1)
	eq(t, js.Rb(m0["two"]), false)
	eq(t, js.Rs(m0["three"]), "")
	eq(t, js.Ri(js.Ra(m0["four"])[0]), 1)

}
