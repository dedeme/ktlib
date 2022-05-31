// Copyright 26-May-2017 ºDeme
// GNU General Public License - V3 <http://www.gnu.org/licenses/>

// Strings management
package str

import (
	"fmt"
	"strings"
	"unicode/utf16"
)

// Returns the remains elements of 's' after make an 'str.take' operation
func Drop(s string, n int) string {
	switch {
	case n <= 0:
		return s
	case n > len(s):
		return ""
	default:
		return s[n:]
	}
}

// Returns 'true' if 's' ends with 'subs'.
func Ends(s, subs string) bool {
	return strings.HasSuffix(s, subs)
}

// Returns the result of fill 'template' with values of 'values'
//    You can use:
//      '%t': for booleans.
//      '%d': for 'int's.
//      '%f': for floats.
//      '%s': for strings.
//      '%v': for any object.
//    '%f' has the folowing optional constraints:
//      %'width'.'precision'f
// Examples:
//    str.fmt("%s, %d, %f", "abc", 33, 12.5) == "abc, 33, 12.500000"
//    str.fmt("|%12f|%.3f|%5.1f|", 12.5, 12.5, 12.5) ==
//        "|   12.500000|12.500| 12.5|"
func Fmt(template string, values ...any) string {
	return fmt.Sprintf(template, values...)
}

// Returns an UTF8 string from a ISO string.
func FromIso(s string) string {
	var r []byte
	for _, b := range []byte(s) {
		if b < 0x80 {
			r = append(r, b)
		} else {
			r = append(r, 0xc0|(b&0xc0)>>6, 0x80|(b&0x3f))
		}
	}
	return string(r)
}

// Returns a string from UTF16 codepoints.
func FromUtf16(codepoints []uint16) string {
	return string(utf16.Decode(codepoints))
}

// Returns position of the first occurence of 'subs' in 's', counting by bytes,
// or -1 if 'subs' is not in 's'.
func Index(s, subs string) int {
	return strings.Index(s, subs)
}

// Returns position of the first occurence of 'b' in 's', counting by bytes,
// or -1 if 'b' is not in 's'.
func IndexByte(s string, b byte) int {
	return strings.IndexByte(s, b)
}

// Returns position of the first occurence of 'subs' in 's', counting by bytes
// from position 'ix' inclusive, or -1 if 'subs' is not in 's'.
func IndexFrom(s, subs string, ix int) int {
	s2 := Drop(s, ix)
	if s2 == "" {
		return -1
	}
	r := strings.Index(s2, subs)
	if r == -1 {
		return -1
	}
	return r + ix
}

// Returns position of the first occurence of 'b' in 's', counting by bytes
// from position 'ix', or -1 if 'b' is not in 's'.
func IndexByteFrom(s string, b byte, ix int) int {
	s2 := Drop(s, ix)
	if s2 == "" {
		return -1
	}
	r := strings.IndexByte(s2, b)
	if r == -1 {
		return -1
	}
	return r + ix
}

// Returns position of the last occurence of 'subs' in 's', counting by bytes,
// or -1 if 'subs' is not in 's'.
func LastIndex(s, subs string) int {
	return strings.LastIndex(s, subs)
}

// Returns position of the last occurence of 'b' in 's', counting by bytes,
// or -1 if 'b' is not in 's'.
func LastIndexByte(s string, b byte) int {
	return strings.LastIndexByte(s, b)
}

// Returns 's' removing starting spaces (" \n\t\r").
func Ltrim(s string) string {
	return strings.TrimLeft(s, " \n\t\r")
}

// Returns 's' replacing all ocurreces of 'old' by 'new'.
func Replace(s, old, new string) string {
	return strings.ReplaceAll(s, old, new)
}

// Returns 's' removing trailing spaces (" \n\t\r").
func Rtrim(s string) string {
	return strings.TrimRight(s, " \n\t\r")
}

// Returns an array with 's' splitted by 'sep'.
// Examples:
//   assert arr.size(str.split("", "")) == 0;
//   assert arr.join(str.split("", ""), "") == "";
//   assert arr.size(str.split("a", "")) == 1;
//   assert arr.join(str.split("a", ""), "") == "a";
//   assert arr.size(str.split("añ", "")) == 2;
//   assert arr.join(str.split("añ", ""), "") == "añ";
//   assert arr.size(str.split("", ";")) == 1;
//   assert arr.join(str.split("", ";"), ";") == "";
//   assert arr.size(str.split("ab;cd;", ";")) == 3;
//   assert arr.join(str.split("ab;cd;", ";"), ";") == "ab;cd;";
//   assert arr.size(str.split("ab;cd", ";")) == 2;
//   assert arr.join(str.split("ab;cd", ";"), ";") == "ab;cd";
func Split(s, sep string) []string {
	return strings.Split(s, sep)
}

// Equals to split, triming each strings in the resulting array.
func SplitTrim(s, sep string) []string {
	ss := strings.Split(s, sep)
	r := make([]string, len(ss))
	for i, e := range ss {
		r[i] = e
	}
	return r
}

// Returns 'true' if 's' starts with 'subs'.
func Starts(s, subs string) bool {
	return strings.HasPrefix(s, subs)
}

// Returns an array with the 'n' first bytes of 's'.
//   -If 'n <= 0' returns the complete string.
//   -if 'n >= len(s)' returns an empty string.
func Take(s string, n int) string {
	switch {
	case n <= 0:
		return ""
	case n > len(s):
		return s
	default:
		return s[:n]
	}
}

// Returns 's' with all runes in lowercase.
func ToLower(s string) string {
	return strings.ToLower(s)
}

// Returns 's' with all runes in uppercase.
func ToUpper(s string) string {
	return strings.ToUpper(s)
}

// Returns an array with codepoints of 's'
func ToUtf16(s string) []uint16 {
	return utf16.Encode([]rune(s))
}

// Returns 's' removing starting and trailing spaces.
func Trim(s string) string {
	return strings.TrimSpace(s)
}
