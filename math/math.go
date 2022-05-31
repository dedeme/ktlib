// Copyright 29-May-2022 ºDeme
// GNU General Public License - V3 <http://www.gnu.org/licenses/>

// Mathematical utilities.
package math

import (
	"github.com/dedeme/ktlib/str"
	gmath "math"
	"math/rand"
	"strconv"
)

// Returns the absolute value of n.
func Abs(n float64) float64 {
	return gmath.Abs(n)
}

// Returns arccosine of n.
func Acos(n float64) float64 {
	return gmath.Acos(n)
}

// Returns arc-hyperbolic cosine of n.
func Acosh(n float64) float64 {
	return gmath.Acosh(n)
}

// Returns arcsine of n.
func Asin(n float64) float64 {
	return gmath.Asin(n)
}

// Returns arc-hyperbolic sine of n.
func Asinh(n float64) float64 {
	return gmath.Asinh(n)
}

// Returns arctangent of n.
func Atan(n float64) float64 {
	return gmath.Atan(n)
}

// Returns arc-hyperbolic tangent of n.
func Atanh(n float64) float64 {
	return gmath.Atanh(n)
}

// Returns the next greater integer of n.
func Ceil(n float64) float64 {
	return gmath.Ceil(n)
}

// Returns cosine of n.
func Cos(n float64) float64 {
	return gmath.Cos(n)
}

// Returns hyperbolic cosine of n.
func Cosh(n float64) float64 {
	return gmath.Cosh(n)
}

// Returns if n1 is equals to n2 +- gap
func Eq(n1, n2, gap float64) bool {
	return gmath.Abs(n1-n2) <= gmath.Abs(gap)
}

// Returns 'pow(e, n)'.
func Exp(n float64) float64 {
	return gmath.Exp(n)
}

// Returns 'pow(2, n)'.
func Exp2(n float64) float64 {
	return gmath.Exp2(n)
}

// Returns the next less integer of n.
func Floor(n float64) float64 {
	return gmath.Floor(n)
}

// Returns the float value of 's', which contains a number in English format.
func FromEn(s string) float64 {
	r, err := strconv.ParseFloat(str.Replace(s, ",", ""), 64)
	if err != nil {
		panic(err)
	}
	return r
}

// Returns the float value of 's', which contains a number in ISO format.
func FromIso(s string) float64 {
	r, err := strconv.ParseFloat(str.Replace(str.Replace(s, ".", ""), ",", "."), 64)
	if err != nil {
		panic(err)
	}
	return r
}

// Returns the float value of 's', which contains a number in standard format.
func FromStr(s string) float64 {
	r, err := strconv.ParseFloat(s, 64)
	if err != nil {
		panic(err)
	}
	return r
}

// Returns napierian logarithm of n.
func Log(n float64) float64 {
	return gmath.Log(n)
}

// Returns base 10 logarithm of n.
func Log10(n float64) float64 {
	return gmath.Log10(n)
}

// Returns base 2 logarithm of n.
func Log2(n float64) float64 {
	return gmath.Log2(n)
}

// Returns the greater of n1 and n2.
func Max(n1, n2 float64) float64 {
	return gmath.Max(n1, n2)
}

// Returns the less of n1 and n2.
func Min(n1, n2 float64) float64 {
	return gmath.Min(n1, n2)
}

// Returns 'base' raised to the power of 'ex'.
func Pow(base, ex float64) float64 {
	return gmath.Pow(base, ex)
}

// Returns 10 raised to the power of 'ex'.
func Pow10(n int) float64 {
	return gmath.Pow10(n)
}

// Returns a random number between 0 (inclusive) and 1 (exclusive).
//
// 'sys.Rand' should be called previously.
func Rnd() float64 {
	return rand.Float64()
}

// Returns a random integer between 0 (inclusive) and top (exclusive).
//
// 'sys.Rand' should be called previously.
func Rndi(top int) int {
	return rand.Intn(top)
}

// Returns a random integer between 0 (inclusive) and top (exclusive).
//
// 'sys.Rand' should be called previously.
func Rndi64(top int64) int64 {
	return rand.Int63n(top)
}

// Rounds 'n' with 'dec' decimal (between 0 and 9, both inclusive).
func Round(n float64, dec int) float64 {
	switch {
	case dec <= 0:
		return gmath.Round(n)
	case dec == 1:
		return gmath.Round(n*10.0) / 10.0
	case dec == 2:
		return gmath.Round(n*100.0) / 100.0
	case dec == 3:
		return gmath.Round(n*1000.0) / 1000.0
	case dec == 4:
		return gmath.Round(n*10000.0) / 10000.0
	case dec == 5:
		return gmath.Round(n*100000.0) / 100000.0
	case dec == 6:
		return gmath.Round(n*1000000.0) / 1000000.0
	case dec == 7:
		return gmath.Round(n*10000000.0) / 10000000.0
	case dec == 8:
		return gmath.Round(n*100000000.0) / 100000000.0
	default:
		return gmath.Round(n*1000000000.0) / 1000000000.0
	}
}

// Returns sine of n.
func Sin(n float64) float64 {
	return gmath.Sin(n)
}

// Returns hyperbolic sine of n.
func Sinh(n float64) float64 {
	return gmath.Sinh(n)
}

// Returns the square root of n.
func Sqrt(n float64) float64 {
	return gmath.Sqrt(n)
}

// Returns tan of n.
func Tan(n float64) float64 {
	return gmath.Tan(n)
}

// Returns hyperbolic sine of n.
func Tanh(n float64) float64 {
	return gmath.Tanh(n)
}

// Returns the int value of 's'
func ToInt(s string) int {
	r, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return r
}

// Returns the int64 value of 's'
func ToInt64(s string) int64 {
	r, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		panic(err)
	}
	return r
}

// Returns the integer part of n.
func Trunc(n float64) float64 {
	return gmath.Trunc(n)
}
