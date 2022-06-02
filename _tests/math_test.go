// Copyright 29-May-2022 ºDeme
// GNU General Public License - V3 <http://www.gnu.org/licenses/>

package _tests

import (
	"github.com/dedeme/ktlib/math"
	"github.com/dedeme/ktlib/sys"
	"testing"
)

func TestMath(t *testing.T) {
	eqm := func(f1, f2 float64) bool {
		return math.Eq(f1, f2, 0.0000001)
	}

	eq(t, eqm(math.FromEn("12,520.11"), 12520.11), true)
	eq(t, eqm(math.FromIso("-12.520,11"), -12520.11), true)
	eq(t, eqm(math.FromStr("-12520.11"), -12520.11), true)

	eq(t, math.ToInt("-12"), -12)
	eq(t, math.ToInt64("-12"), -12)

	eq(t, eqm(math.Abs(3.2), 3.2), true)
	eq(t, eqm(math.Abs(-3.2), 3.2), true)

	eq(t, eqm(math.Acos(0.5), 1.0471975511966), true)
	eq(t, eqm(math.Acosh(1.5), 0.962423650119207), true)
	eq(t, eqm(math.Asin(0.5), 0.523598775598299), true)
	eq(t, eqm(math.Asinh(1.5), 1.19476321728711), true)
	eq(t, eqm(math.Atan(0.5), 0.463647609000806), true)
	eq(t, eqm(math.Atanh(0.5), 0.549306144334055), true)

	eq(t, eqm(math.Ceil(1.2), 2.0), true)
	eq(t, eqm(math.Ceil(-1.2), -1.0), true)
	eq(t, eqm(math.Floor(1.2), 1.0), true)
	eq(t, eqm(math.Floor(-1.2), -2.0), true)
	eq(t, eqm(math.Trunc(1.2), 1.0), true)
	eq(t, eqm(math.Trunc(-1.2), -1.0), true)
	eq(t, eqm(math.Round(1.49999, 0), 1.0), true)
	eq(t, eqm(math.Round(1.5, 0), 2.0), true)
	eq(t, eqm(math.Round(1.6149999, 2), 1.61), true)
	eq(t, eqm(math.Round(1.615, 2), 1.62), true)

	eq(t, eqm(math.Cos(1.5), 0.070737201667703), true)
	eq(t, eqm(math.Cosh(1.5), 2.35240961524325), true)
	eq(t, eqm(math.Sin(1.5), 0.997494986604055), true)
	eq(t, eqm(math.Sinh(1.5), 2.12927945509482), true)
	eq(t, eqm(math.Tan(1.5), 14.1014199471717), true)
	eq(t, eqm(math.Tanh(1.5), 0.905148253644866), true)

	eq(t, eqm(math.Exp(1.5), 4.48168907033806), true)
	eq(t, eqm(math.Exp2(1.5), 2.82842712474619), true)
	eq(t, eqm(math.Pow(2.0, 1.5), 2.82842712474619), true)
	eq(t, eqm(math.Pow10(2), 100.0), true)
	eq(t, eqm(math.Sqrt(9.0), 3.0), true)

	eq(t, eqm(math.Log(1.5), 0.405465108108164), true)
	eq(t, eqm(math.Log10(1.5), 0.176091259055681), true)
	eq(t, eqm(math.Log2(1.5), 0.584962500721), true)

	eq(t, eqm(math.Max(1.5, -1.5), 1.5), true)
	eq(t, eqm(math.Min(1.5, -1.5), -1.5), true)

	sys.Rand()
	//t.Fatal(math.Rnd())
	//t.Fatal(math.Rndi(6))
	//t.Fatal(math.Rndi64(6))
}
