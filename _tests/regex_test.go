// Copyright 30-May-2020 ºDeme
// GNU General Public License - V3 <http://www.gnu.org/licenses/>

package _tests

import (
	"github.com/dedeme/ktlib/arr"
	"github.com/dedeme/ktlib/regex"
	"testing"
)

func TestRegex(t *testing.T) {
	eq(t, regex.Replace("-ab-axxb-", "a(x*)b", "T"), "-T-T-")
	eq(t, regex.Replace("paranormal", "(a)(.)", "$1"), "paaorma")
	eq(t, regex.Replace("paranormal", "(a)(.)", "$2"), "prnorml")
	eq(t, regex.Replace("paranormal", "(a)(.)", "$3"), "porm")
	eq(t, regex.Replace("paranormal", "(a)(.)", "$1x"), "porm")
	eq(t, regex.Replace("paranormal", "(a)(.)", "${1}x"), "paxaxormax")
	eq(t, regex.Replace("paranormal", "(?P<one>a)(.)", "$one"), "paaorma")
	eq(t, regex.Replace("-ab-axxb-", "a(x*)b", "$1AB"), "---")
	eq(t, regex.Replace("-ab-axxb-", "a(x*)b", "${1}AB"), "-AB-xxAB-")

	eq(t, arr.ToStr(regex.Matches("paranormal", "a.")), "[ar, an, al]")
	eq(t, arr.ToStr(regex.Matches("paranormal", "xx")), "[]")
}
