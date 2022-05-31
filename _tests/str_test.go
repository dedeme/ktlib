// Copyright 26-May-2020 ºDeme
// GNU General Public License - V3 <http://www.gnu.org/licenses/>

package _tests

import (
	"github.com/dedeme/ktlib/arr"
	"github.com/dedeme/ktlib/str"
	"testing"
)

func TestStr(t *testing.T) {
	eq(t, "abc"[0], 'a')
	eq(t, "abc"[2], 'c')

	eq(t, str.Ends("abc", "c"), true)
	eq(t, str.Ends("abc", ""), true)
	eq(t, str.Ends("", ""), true)
	eq(t, !str.Ends("abc", "zc"), true)
	eq(t, !str.Ends("abc", "1abc"), true)
	eq(t, !str.Ends("", "x"), true)

	eq(t, str.Starts("abc", "a"), true)
	eq(t, str.Starts("abc", ""), true)
	eq(t, str.Starts("", ""), true)
	eq(t, !str.Starts("abc", "ac"), true)
	eq(t, !str.Starts("abc", "abc1"), true)
	eq(t, !str.Starts("", "x"), true)

	eq(t, str.IndexByte("abcb", 'b'), 1)
	eq(t, str.Index("abcb", "bc"), 1)
	eq(t, str.Index("abcb", "abcb"), 0)
	eq(t, str.Index("abcb", "abcbc"), -1)
	eq(t, str.Index("a", ""), 0)
	eq(t, str.Index("año", "o"), 3)

	eq(t, str.LastIndexByte("abcb", 'b'), 3)
	eq(t, str.LastIndex("abcb", "bc"), 1)
	eq(t, str.LastIndex("abcb", "abcb"), 0)
	eq(t, str.LastIndex("abcb", "abcbc"), -1)
	eq(t, str.LastIndex("a", ""), 1)
	eq(t, str.LastIndex("año", "o"), 3)

	eq(t, str.IndexByteFrom("abcb", 'b', 3), 3)
	eq(t, str.IndexFrom("abcb", "bc", 3), -1)
	eq(t, str.IndexFrom("abcb", "abcb", 0), 0)
	eq(t, str.IndexFrom("abcb", "abcbc", 0), -1)
	eq(t, str.IndexFrom("a", "", 1), -1)
	eq(t, str.IndexFrom("año", "o", 1), 3)

	eq(t, str.Fmt("%s, %d, %f", "abc", 33, 12.5), "abc, 33, 12.500000")
	eq(t, str.Fmt("|%12f|%.3f|%5.1f|", 12.5, 12.5, 12.5), "|   12.500000|12.500| 12.5|")

	eq(t, string([]rune("")), "")
	eq(t, len([]rune("")), 0)
	eq(t, string([]rune("añña世界")), "añña世界")
	eq(t, len([]rune("añña世界")), 6)

	eq(t, str.FromUtf16(str.ToUtf16("")), "")
	eq(t, len(str.ToUtf16("")), 0)
	eq(t, str.FromUtf16(str.ToUtf16("añña世界")), "añña世界")
	eq(t, len(str.ToUtf16("añña世界")), 6)

	eq(t, str.ToUpper(""), "")
	eq(t, str.ToUpper("año"), "AÑO")
	eq(t, str.ToLower(""), "")
	eq(t, str.ToLower("AÑO"), "año")

	eq(t, str.FromIso(string([]byte{97, 241, 111})), "año")

	eq(t, str.Replace(str.Replace("año", "a", "ca"), "ño", "ñón"), "cañón")

	eq(t, len(str.Split("", "")), 0)
	eq(t, arr.Join(str.Split("", ""), ""), "")
	eq(t, len(str.Split("a", "")), 1)
	eq(t, arr.Join(str.Split("a", ""), ""), "a")
	eq(t, len(str.Split("añ", "")), 2)
	eq(t, arr.Join(str.Split("añ", ""), ""), "añ")
	eq(t, len(str.Split("", ";")), 1)
	eq(t, arr.Join(str.Split("", ";"), ";"), "")
	eq(t, len(str.Split("ab;cd;", ";")), 3)
	eq(t, arr.Join(str.Split("ab;cd;", ";"), ";"), "ab;cd;")
	eq(t, len(str.Split("ab;cd", ";")), 2)
	eq(t, arr.Join(str.Split("ab;cd", ";"), ";"), "ab;cd")
	eq(t, len(str.Split("ab;", ";")), 2)
	eq(t, arr.Join(str.Split("ab;", ";"), ";"), "ab;")
	eq(t, len(str.Split("ab", ";")), 1)
	eq(t, arr.Join(str.Split("ab", ";"), ";"), "ab")
	eq(t, len(str.Split("", "ñ")), 1)
	eq(t, arr.Join(str.Split("", "ñ"), "ñ"), "")
	eq(t, len(str.Split("abñcdñ", "ñ")), 3)
	eq(t, arr.Join(str.Split("abñcdñ", "ñ"), "ñ"), "abñcdñ")
	eq(t, len(str.Split("abñcd", "ñ")), 2)
	eq(t, arr.Join(str.Split("abñcd", "ñ"), "ñ"), "abñcd")
	eq(t, len(str.Split("abñ", "ñ")), 2)
	eq(t, arr.Join(str.Split("abñ", "ñ"), "ñ"), "abñ")
	eq(t, len(str.Split("ab", "ñ")), 1)
	eq(t, arr.Join(str.Split("ab", "ñ"), "ñ"), "ab")
	eq(t, len(str.Split("", "--")), 1)
	eq(t, arr.Join(str.Split("", "--"), "--"), "")
	eq(t, len(str.Split("ab--cd--", "--")), 3)
	eq(t, arr.Join(str.Split("ab--cd--", "--"), "--"), "ab--cd--")
	eq(t, len(str.Split("ab--cd", "--")), 2)
	eq(t, arr.Join(str.Split("ab--cd", "--"), "--"), "ab--cd")
	eq(t, len(str.Split("ab--", "--")), 2)
	eq(t, arr.Join(str.Split("ab--", "--"), "--"), "ab--")
	eq(t, len(str.Split("ab", "--")), 1)
	eq(t, arr.Join(str.Split("ab", "--"), "--"), "ab")

	eq(t, ""[:0], "")
	eq(t, ""[0:], "")
	eq(t, ""[0:0], "")
	eq(t, ""[:], "")
	eq(t, "abc"[:0], "")
	eq(t, "abc"[:1], "a")
	eq(t, "abc"[:3], "abc")
	eq(t, "abc"[0:], "abc")
	eq(t, "abc"[1:], "bc")
	eq(t, "abc"[3:], "")
	eq(t, "abc"[0:2], "ab")
	eq(t, "abc"[1:2], "b")
	eq(t, "abc"[1:1], "")
	eq(t, "abc"[:], "abc")

}
