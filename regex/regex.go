// Copyright 30-May-2020 ºDeme
// GNU General Public License - V3 <http://www.gnu.org/licenses/>

// Regular expressions.
package regex

import "regexp"

// Returns an array with matches of 'rg' in 's'".
//    s : String for searching.
//    rg: Regular expression.
// Examples:
//    arr.ToStr(regex.Matches("paranormal", "a.")) ==> "[ar, an, al]"
//    arr.ToStr(regex.Matches("paranormal", "xx")) ==> "[]"
func Matches(s, rg string) []string {
	re := regexp.MustCompile(rg)
	rs := re.FindAllString(s, -1)
	if rs != nil {
		return rs
	} else {
		return []string{}
	}
}

// Returns the result of replace "rg" by "repl" in "s".
//    s   : Original string.
//    rg  : Regular expression.
//    repl: Replacement. It allows to use templates with '$'.
// Examples:
//    regex.Replace("-ab-axxb-", "a(x*)b", "T") ==> "-T-T-";
//    regex.Replace("paranormal", "(a)(.)", "$1") ==> "paaorma";
//    regex.Replace("paranormal", "(a)(.)", "$2") ==> "prnorml";
//    regex.Replace("paranormal", "(a)(.)", "$3") ==> "porm";
//    regex.Replace("paranormal", "(a)(.)", "$1x") ==> "porm";
//    regex.Replace("paranormal", "(a)(.)", "${1}x") ==> "paxaxormax";
//    regex.Replace("paranormal", "(?P<one>a)(.)", "$one") ==> "paaorma";
//    regex.Replace("-ab-axxb-", "a(x*)b", "$1AB") ==> "---";
//    regex.Replace("-ab-axxb-", "a(x*)b", "${1}AB") ==> "-AB-xxAB-";
func Replace(s, rg, repl string) string {
	re := regexp.MustCompile(rg)
	return re.ReplaceAllString(s, repl)
}
