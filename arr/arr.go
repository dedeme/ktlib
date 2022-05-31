// Copyright 25-May-2020 ºDeme
// GNU General Public License - V3 <http://www.gnu.org/licenses/>

// Utilities for managing slices.
package arr

import (
	"fmt"
	"math/rand"
	"sort"
	"strings"
)

// Returns 'true' if 'fn' returns true with every element of 'a'.
func All[T any](a []T, fn func(T) bool) bool {
	for _, e := range a {
		if !fn(e) {
			return false
		}
	}
	return true
}

// Returns 'true' if 'fn' returns true with at least one element of 'a'.
func Any[T any](a []T, fn func(T) bool) bool {
	for _, e := range a {
		if fn(e) {
			return true
		}
	}
	return false
}

// Removes in place every element of 'a'.
func Clear[T any](a *[]T) {
	*a = []T{}
}

// Returns a shallow copy of 'a'.
func Copy[T any](a []T) []T {
	r := make([]T, len(a))
	for i, e := range a {
		r[i] = e
	}
	return r
}

// Appends a2 to a1 in place.
func Cat[T any](a1 *[]T, a2 []T) {
	*a1 = append(*a1, a2...)
}

// Returns a new slice with the remains elements of 'a' after make an
// 'arr.Take' operation
func Drop[T any](a []T, n int) []T {
	l := len(a)
	if n < 0 {
		n = 0
	} else if n > l {
		return []T{}
	}
	return Copy(a[n:l])
}

// Returns a new slice with the remains elements of 'a' after make an
// 'arr.TakeWhile' operation
func DropWhile[T any](a []T, fn func(T) bool) []T {
	ix := -1
	for i, e := range a {
		if !fn(e) {
			ix = i
			break
		}
	}
	if ix == -1 {
		return []T{}
	}
	return Copy(a[ix:])
}

//  Returns two new slices:
//    'els', wiht elements of 'a' witout duplicates.
//    'dup', with duplicates of 'a'. There is only one copy of each duplicate.
func Duplicates[T comparable](a []T) (els, dup []T) {
	inEls := func(el T) bool {
		for _, e := range els {
			if e == el {
				return true
			}
		}
		return false
	}
	inDup := func(el T) bool {
		for _, e := range dup {
			if e == el {
				return true
			}
		}
		return false
	}

	for _, e := range a {
		if inEls(e) {
			if !inDup(e) {
				dup = append(dup, e)
			}
		} else {
			els = append(els, e)
		}
	}
	return
}

//  Returns two new slices:
//    'els', wiht elements of 'a' witout duplicates.
//    'dup', with duplicates of 'a'. There is only one copy of each duplicate.
//  'fn' returns 'true' when its two elements are equals.
func Duplicatesf[T any](a []T, fn func(e1, e2 T) bool) (els, dup []T) {
	inEls := func(el T) bool {
		for _, e := range els {
			if fn(e, el) {
				return true
			}
		}
		return false
	}
	inDup := func(el T) bool {
		for _, e := range dup {
			if fn(e, el) {
				return true
			}
		}
		return false
	}

	for _, e := range a {
		if inEls(e) {
			if !inDup(e) {
				dup = append(dup, e)
			}
		} else {
			els = append(els, e)
		}
	}
	return
}

// Executes 'fn' with each element of 'a'.
func Each[T any](a []T, fn func(T)) {
	for _, e := range a {
		fn(e)
	}
}

// Executes 'fn' with each element of 'a'.
func EachIx[T any](a []T, fn func(T, int)) {
	for i, e := range a {
		fn(e, i)
	}
}

// Returns true if a1 == a2
func Eq[T comparable](a1, a2 []T) bool {
	l := len(a1)
	if l != len(a2) {
		return false
	}
	for i := 0; i < l; i++ {
		if a1[i] != a2[i] {
			return false
		}
	}
	return true
}

// Returns true if a1 == a2 comparing with 'fn'
//    'fn' returns 'true' when its two elements are equals.
func Eqf[T any](a1, a2 []T, fn func(T, T) bool) bool {
	l := len(a1)
	if l != len(a2) {
		return false
	}
	for i := 0; i < l; i++ {
		if !fn(a1[i], a2[i]) {
			return false
		}
	}
	return true
}

// Returns 'true' if 'A' has no element.
func Empty[T any](a []T) bool {
	return len(a) == 0
}

// Returns an new slice with elements which produce 'true' with 'fn'.
func Filter[T any](a []T, fn func(T) bool) []T {
	var r []T
	for _, e := range a {
		if fn(e) {
			r = append(r, e)
		}
	}
	return r
}

// Filter 'a' in place, removing elements which produce 'false' with 'fn'
func FilterIn[T any](a *[]T, fn func(T) bool) {
	*a = Filter(*a, fn)
}

// Returns the first element which produces 'true' with 'fn' or 'ok=false'
func Find[T any](a []T, fn func(T) bool) (el T, ok bool) {
	for _, e := range a {
		if fn(e) {
			el = e
			ok = true
			return
		}
	}
	return
}

// Returns the index of the first element which produce 'true' with 'fn',
// or -1 if such element does not exist.
func Index[T any](a []T, fn func(T) bool) int {
	for i, e := range a {
		if fn(e) {
			return i
		}
	}
	return -1
}

// Returns a string with elements of 'a' joined with 'sep'.
func Join(a []string, sep string) string {
	return strings.Join(a, sep)
}

// Returns a new slice retrieved by appliyng 'fn' to each element of 'a'.
func Map[T any, U any](a []T, fn func(T) U) []U {
	l := len(a)
	r := make([]U, l, l)
	for i, e := range a {
		r[i] = fn(e)
	}
	return r
}

// Returns and remove the last element of 'a' or raise 'panic' if 'a' is empty.
func Pop[T any](a *[]T) T {
	a2 := *a
	l1 := len(a2) - 1
	r := a2[l1]
	*a = a2[:l1]
	return r
}

// Returns the last element of 'a' or raise 'panic' if 'a' is empty.
func Peek[T any](a []T) T {
	return a[len(a)-1]
}

// Returna a new array adding an element at the end of 'a'.
func Push[T any](a *[]T, e T) {
	*a = append(*a, e)
}

// Returns the result of calculate 'seed = fn(seed, e)' with every element
// of 'a'.
//    For example:
//    arr.Reduce([]int{1, 2, 3}, 0, fun(seed int, e int) int {
//      return seed + e
//    }) // returns 6
func Reduce[T any, U any](a []T, seed U, fn func(U, T) U) U {
	for _, e := range a {
		seed = fn(seed, e)
	}
	return seed
}

// Removes the element at index 'ix', or raise 'panic' if it does not exist.
func Remove[T any](a *[]T, ix int) {
	a2 := *a
	*a = append(a2[:ix], a2[ix+1:]...)
}

// Removes elements from index 'begin' (inclusive) to index 'end' (exclusive),
// or raise 'panic' if some element does not exist.
//
// If begin >= end this function does nothing.
func RemoveRange[T any](a *[]T, begin, end int) {
	if begin >= end {
		return
	}
	a2 := *a
	*a = append(a2[:begin], a2[end:]...)
}

// Returns a new slice with elements of 'a' reversed.
func Reverse[T any](a []T) []T {
	r := make([]T, len(a))
	ir := 0
	for ia := len(a) - 1; ia >= 0; ia-- {
		r[ir] = a[ia]
		ir++
	}
	return r
}

// Reverses elements of 'a' in place.
func ReverseIn[T any](a []T) {
	right := len(a) - 1
	for left := 0; left < right; left++ {
		a[left], a[right] = a[right], a[left]
		right--
	}
}

// Removes and returns the first element of 'a', or raise 'panic' if 'a' is empty.
func Shift[T any](a *[]T) T {
	a2 := *a
	r := a2[0]
	*a = a2[1:]
	return r
}

// Reorders randomly elements of 'a' in place.
//
// Prevously to call this function, 'sys.Rand' shuld be called.
func Shuffle[T any](a []T) {
	for i := len(a); i > 1; {
		n := rand.Intn(i)
		i--
		a[n], a[i] = a[i], a[n]
	}
}

//  Sorts elements of 'A' from less to greater.
//    'less' is a function which returns 'true' if the first paramenter is
//    less than the second one.
//    NOTE: If 'less' returns 'true' when the first parameter is greater than the
//      second one, then the order is from greater to less.
func Sort[T any](a []T, fn func(T, T) bool) {
	sort.Slice(a, func(i, j int) bool {
		return fn(a[i], a[j])
	})
}

// Returns a copy of the first 'n' elements of 'a'.
//    -If 'n <= 0' returns the complete array.
//    -if 'n >= len(a)' returns an empty array.
func Take[T any](a []T, n int) []T {
	l := len(a)
	if n < 0 {
		return []T{}
	} else if n > l {
		n = l
	}
	return Copy(a[:n])
}

// Returns the first elements of 'a' which produce 'true' with 'fn'.
func TakeWhile[T any](a []T, fn func(T) bool) []T {
	var r []T
	for _, e := range a {
		if fn(e) {
			r = append(r, e)
		} else {
			break
		}
	}
	return r
}

// Returns a representation of 'a'.
func ToStr[T any](a []T) string {
	a2 := Map(a, func(e T) string {
		return fmt.Sprintf("%v", e)
	})
	return "[" + Join(a2, ", ") + "]"
}

// Preppend in place an element at the beginning of 'a'.
func Unshift[T any](a *[]T, el T) {
	a2 := *a
	r := make([]T, len(a2)+1)
	r[0] = el
	for i, e := range a2 {
		r[i+1] = e
	}
	*a = r
}
