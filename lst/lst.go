// Copyright 17-Jul-2022 ºDeme
// GNU General Public License - V3 <http://www.gnu.org/licenses/>

// Lazy list.
package lst

import (
	"fmt"
	"strings"
)

type T[E any] struct {
	head E
	tail func() *T[E]
}

// Returns 'true' if every element of 'l' is equals to 'e'.
func All[E comparable](l *T[E], e E) bool {
	for {
		if l.IsEmpty() {
			return true
		}
		if l.head != e {
			return false
		}
		l = l.tail()
	}
}

// Returns 'true' if 'fn' returns true with every element of 'l'.
func (l *T[E]) Allf(fn func(E) bool) bool {
	for {
		if l.IsEmpty() {
			return true
		}
		if !fn(l.head) {
			return false
		}
		l = l.tail()
	}
}

// Returns 'true' if at least an element of 'l' is equals to 'e'.
func Any[E comparable](l *T[E], e E) bool {
	for {
		if l.IsEmpty() {
			return false
		}
		if l.head == e {
			return true
		}
		l = l.tail()
	}
}

// Returns 'true' if 'fn' returns true with at least one element of 'l'.
func (l *T[E]) Anyf(fn func(E) bool) bool {
	for {
		if l.IsEmpty() {
			return false
		}
		if fn(l.head) {
			return true
		}
		l = l.tail()
	}
}

// Appends l2 to l1.
func (l *T[E]) Cat(l2 *T[E]) *T[E] {
	l = l.Reverse()
	for {
		if l.IsEmpty() {
			break
		}
		l2 = l2.Cons(l.head)
		l = l.tail()
	}
	return l2
}

// Adds an element at the head of 'l'.
func (l *T[E]) Cons(e E) *T[E] {
	return New(e, func() *T[E] {
		return l
	})
}

// Returns the elements number of 'l'.
func (l *T[E]) Count() int {
	r := 0
	for {
		if l.IsEmpty() {
			break
		}
		r++
		l = l.tail()
	}
	return r
}

// Returns the remains elements of 'l' after make a 'lst.Take' operation
func (l *T[E]) Drop(n int) *T[E] {
	_, r := l.TakeDrop(n)
	return r
}

// Returns the remains elements of 'a' after make a 'lst.TakeWhile' operation
func (l *T[E]) DropWhile(fn func(E) bool) *T[E] {
	_, r := l.TakeDropWhile(fn)
	return r
}

//  Returns two new lists:
//    'els', wiht elements of 'l' witout duplicates.
//    'dup', with duplicates of 'l'. It can be several copies of each duplicate.
func Duplicates[E comparable](l *T[E]) (els, dup *T[E]) {
	els = NewEmpty[E]()
	dup = NewEmpty[E]()
	for {
		if l.IsEmpty() {
			break
		}
		if Any(els, l.head) {
			dup = dup.Cons(l.head)
		} else {
			els = els.Cons(l.head)
		}
		l = l.tail()
	}
	return
}

//  Returns two new lists:
//    'els', wiht elements of 'l' witout duplicates.
//    'dup', with duplicates of 'l'. It can be several copies of each duplicate.
//  'fn' returns 'true' when its two elements are equals.
func (l *T[E]) Duplicatesf(fn func(e1, e2 E) bool) (els, dup *T[E]) {
	els = NewEmpty[E]()
	dup = NewEmpty[E]()
	for {
		if l.IsEmpty() {
			break
		}
		if els.Anyf(func(e E) bool {
			return fn(l.head, e)
		}) {
			dup = dup.Cons(l.head)
		} else {
			els = els.Cons(l.head)
		}
		l = l.tail()
	}
	return
}

// Executes 'fn' with each element of 'l'.
func (l *T[E]) Each(fn func(E)) {
	for {
		if l.IsEmpty() {
			break
		}
		fn(l.head)
		l = l.tail()
	}
}

// Executes 'fn' with each element of 'l' and its index.
func (l *T[E]) EachIx(fn func(E, int)) {
	i := 0
	for {
		if l.IsEmpty() {
			break
		}
		fn(l.head, i)
		l = l.tail()
		i++
	}
}

// Returns true if l1 == l2
func Eq[E comparable](l1, l2 *T[E]) bool {
	ok := false
	for {
		if l1.IsEmpty() {
			ok = l2.IsEmpty()
			break
		}
		if l2.IsEmpty() {
			break
		}
		if l1.head != l2.head {
			break
		}
		l1 = l1.tail()
		l2 = l2.tail()
	}
	return ok
}

// Returns true if l1 == l2 comparing with 'fn'
//    'fn' returns 'true' when its two elements are equals.
func (l *T[E]) Eqf(l2 *T[E], fnEq func(E, E) bool) bool {
	ok := false
	for {
		if l.IsEmpty() {
			ok = l2.IsEmpty()
			break
		}
		if l2.IsEmpty() {
			break
		}
		if !fnEq(l.head, l2.head) {
			break
		}
		l = l.tail()
		l2 = l2.tail()
	}
	return ok
}

// Returns the elements which produce 'true' with 'fn'.
func (l *T[E]) Filter(fn func(E) bool) *T[E] {
	l = l.Reverse()
	l2 := NewEmpty[E]()
	for {
		if l.IsEmpty() {
			break
		}
		if fn(l.head) {
			l2 = l2.Cons(l.head)
		}
		l = l.tail()
	}
	return l2
}

// Returns the elements which produce 'true' ('yes') and which produce 'false'
// ('not') with 'fn'.
func (l *T[E]) FilterSplit(fn func(E) bool) (yes, not *T[E]) {
	l = l.Reverse()
	yes = NewEmpty[E]()
	not = NewEmpty[E]()
	for {
		if l.IsEmpty() {
			break
		}
		if fn(l.head) {
			yes = yes.Cons(l.head)
		} else {
			not = not.Cons(l.head)
		}
		l = l.tail()
	}
	return
}

// Returns the first element which produces 'true' with 'fn' or 'ok=false'
func (l *T[E]) Find(fn func(E) bool) (e E, ok bool) {
	for {
		if l.IsEmpty() {
			return
		}
		if fn(l.head) {
			e = l.head
			ok = true
			return
		}
		l = l.tail()
	}
}

// Returns the first element of 'l'.
//  If 'l' is empty produce 'panic'.
func (l *T[E]) Head() E {
	if l.tail == nil {
		panic("List is empty")
	}
	return l.head
}

// Returns the index of the first element equals to 'e',
// or -1 if such element does not exist.
func Index[E comparable](l *T[E], e E) int {
	i := 0
	for {
		if l.IsEmpty() {
			return -1
		}
		if l.head == e {
			return i
		}
		l = l.tail()
		i++
	}
}

// Returns the index of the first element which produce 'true' with 'fn',
// or -1 if such element does not exist.
func (l *T[E]) Indexf(fn func(E) bool) int {
	i := 0
	for {
		if l.IsEmpty() {
			return -1
		}
		if fn(l.head) {
			return i
		}
		l = l.tail()
		i++
	}
}

// Returns 'true' if 'l' is empty.
func (l *T[E]) IsEmpty() bool {
	return l.tail == nil
}

// Returns a string with elements of 'l' joined with 'sep'.
func Join(l *T[string], sep string) string {
	return strings.Join(l.ToArr(), sep)
}

// Returns a new list retrieved by appliyng 'fn' to each element of 'a'.
func Map[E, U any](l *T[E], fn func(E) U) *T[U] {
	if l.IsEmpty() {
		return NewEmpty[U]()
	}
	return New(fn(l.head), func() *T[U] {
		return Map(l.tail(), fn)
	})
}

// Basic constructor.
func New[E any](head E, tail func() *T[E]) *T[E] {
	return &T[E]{head, tail}
}

// Returns an empty list.
func NewEmpty[E any]() *T[E] {
	var e E
	return New(e, nil)
}

// Returns a list with elements of 'a' in the same order.
func NewFromArr[E any](a []E) *T[E] {
	if len(a) == 0 {
		return NewEmpty[E]()
	}
	return New(a[0], func() *T[E] {
		return NewFromArr[E](a[1:])
	})
}

// Returns a list with 's' splitted by 'sep'.
// Examples:
//   assert lst.Count(lst.NewSplit("", "")) == 0
//   assert lst.Join(lst.NewSplit("", ""), "") == ""
//   assert lst.Count(lst.NewSplit("a", "")) == 1
//   assert lst.Join(lst.NewSplit("a", ""), "") == "a"
//   assert lst.Count(lst.NewSplit("añ", "")) == 2
//   assert lst.Join(lst.NewSplit("añ", ""), "") == "añ"
//   assert lst.Count(lst.NewSplit("", ";")) == 1
//   assert lst.Join(lst.NewSplit("", ";"), ";") == ""
//   assert lst.Count(lst.NewSplit("ab;cd;", ";")) == 3
//   assert lst.Join(lst.NewSplit("ab;cd;", ";"), ";") == "ab;cd;"
//   assert lst.Count(lst.NewSplit("ab;cd", ";")) == 2
//   assert lst.Join(lst.NewSplit("ab;cd", ";"), ";") == "ab;cd"
func NewSplit(s string, sep string) *T[string] {
	return NewFromArr(strings.Split(s, sep))
}

// Equals to split, triming each string in the resulting list.
func NewSplitTrim(s string, sep string) *T[string] {
	ss := strings.Split(s, sep)
	r := make([]string, len(ss))
	for i, e := range ss {
		r[i] = strings.TrimSpace(e)
	}
	return NewFromArr(r)
}

// Returns a list with integers between begin (inclusive) and end (exclusive)
func NewRange(begin, end int) *T[int] {
	if begin >= end {
		return NewEmpty[int]()
	}
	return New(begin, func() *T[int] {
		return NewRange(begin+1, end)
	})
}

// Equals to NewRange(0, end)
func NewRange0(end int) *T[int] {
	return NewRange(0, end)
}

// Returns a list with integers from '0' (inclusive) to infinity.
func NewRangeInf() *T[int] {
	var new func(i int) *T[int]
	new = func(i int) *T[int] {
		return New(i, func() *T[int] {
			return new(i + 1)
		})
	}
	return new(0)
}

// Executes 'fn' with each element of 'l'.
func Reduce[E, U any](l *T[E], seed U, fn func(U, E) U) U {
	for {
		if l.IsEmpty() {
			break
		}
		seed = fn(seed, l.head)
		l = l.tail()
	}
	return seed
}

// Returns a list with elements of 'l' reversed.
func (l *T[E]) Reverse() *T[E] {
	r := NewEmpty[E]()
	for {
		if l.IsEmpty() {
			break
		}
		r = r.Cons(l.head)
		l = l.tail()
	}
	return r
}

//  Sorts elements of 'l' from less to greater.
//    'less' is a function which returns 'true' if the first paramenter is
//    less than the second one.
//    NOTE: If 'less' returns 'true' when the first parameter is greater than the
//      second one, then the order is from greater to less.
func (l *T[E]) Sort(less func(E, E) bool) *T[E] {
	if l.IsEmpty() {
		return l
	}
	pv := l.head
	lt, gt := l.tail().FilterSplit(func(e E) bool {
		return less(e, pv)
	})
	return lt.Sort(less).Cat(gt.Sort(less).Cons(pv))
}

// Returns a representation of 'l'
func (l *T[E]) String() string {
	l2 := Map(l, func(e E) string {
		return fmt.Sprintf("%v", e)
	})
	return "[:" + Join(l2, ", ") + ":]"
}

// Returns a the first 'n' elements of 'l'.
//    -If 'n >= l.Count()' returns the complete list.
//    -if 'n <= 0' returns an empty list.
func (l *T[E]) Take(n int) *T[E] {
	r, _ := l.TakeDrop(n)
	return r
}

// Returns the first elements of 'l' which produce 'true' with 'fn'.
func (l *T[E]) TakeWhile(fn func(E) bool) *T[E] {
	r, _ := l.TakeDropWhile(fn)
	return r
}

// Returns l.Take(n) [left] and l.Drop(n) [right]
func (l *T[E]) TakeDrop(n int) (left, right *T[E]) {
	left = NewEmpty[E]()
	for i := 0; i < n; i++ {
		if l.IsEmpty() {
			break
		}
		left = left.Cons(l.head)
		l = l.tail()
	}
	left = left.Reverse()
	right = l
	return
}

// Returns l.TakeWhile(fn) [left] and l.DropWhile(fn) [right]
func (l *T[E]) TakeDropWhile(fn func(E) bool) (left, right *T[E]) {
	left = NewEmpty[E]()
	for {
		if l.IsEmpty() || !fn(l.head) {
			break
		}
		left = left.Cons(l.head)
		l = l.tail()
	}
	left = left.Reverse()
	right = l
	return
}

// Returns a list with elements of 'l' removing 'l.Head()'
//  If 'l' is empty produce 'panic'.
func (l *T[E]) Tail() *T[E] {
	if l.tail == nil {
		panic("List is empty")
	}
	return l.tail()
}

// Returns an array with elements of 'l' en the same order.
func (l *T[E]) ToArr() []E {
	var r []E
	for {
		if l.IsEmpty() {
			break
		}
		r = append(r, l.head)
		l = l.tail()
	}
	return r
}
