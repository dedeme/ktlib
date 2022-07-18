// Copyright 17-Jul-2022 ºDeme
// GNU General Public License - V3 <http://www.gnu.org/licenses/>

package _tests

import (
	"github.com/dedeme/ktlib/arr"
	"github.com/dedeme/ktlib/lst"
	"github.com/dedeme/ktlib/sys"
	"testing"
)

func TestLst(t *testing.T) {
	sys.Rand()

	l0 := lst.NewEmpty[int]()
	eq(t, l0.Count(), 0)
	l1 := l0.Cons(1).Cons(0)
	eq(t, l1.String(), "[:0, 1:]")
	l1 = l1.Cat(lst.NewFromArr([]int{2, 3, 4}))
	eq(t, l1.String(), "[:0, 1, 2, 3, 4:]")
	l1 = l1.Cat(lst.NewEmpty[int]())
	eq(t, l1.String(), "[:0, 1, 2, 3, 4:]")

	lr := l0.Filter(func(i int) bool {
		return i%2 == 0
	})
	eq(t, lr.IsEmpty(), true)
	lr = l1.Filter(func(i int) bool {
		return i%2 == 0
	})
	eq(t, lr.String(), "[:0, 2, 4:]")
	lr = l1.Filter(func(i int) bool {
		return i > 100
	})
	eq(t, lr.IsEmpty(), true)
	lr = l1.Filter(func(i int) bool {
		return i < 100
	})
	eq(t, lr.String(), "[:0, 1, 2, 3, 4:]")

	la0 := lst.NewFromArr([]int{})
	eq(t, la0.Count(), 0)
	eq(t, arr.ToStr(la0.ToArr()), "[]")

	la1 := lst.NewFromArr([]int{0, 1, 2, 3, 4})
	eq(t, la1.Count(), 5)
	eq(t, arr.ToStr(la1.ToArr()), "[0, 1, 2, 3, 4]")

	eq(t, la0.Take(2).String(), "[::]")
	eq(t, la0.Cons(0).Take(2).String(), "[:0:]")
	eq(t, la1.Take(2).String(), "[:0, 1:]")
	eq(t, la1.Take(-2).String(), "[::]")
	eq(t, la0.Drop(2).String(), "[::]")
	eq(t, la0.Cons(0).Drop(2).String(), "[::]")
	eq(t, la1.Drop(2).String(), "[:2, 3, 4:]")
	eq(t, la1.Drop(-2).String(), "[:0, 1, 2, 3, 4:]")

	eq(t, la0.TakeWhile(func(e int) bool {
		return e == 0
	}).String(), "[::]")
	eq(t, la0.Cons(0).TakeWhile(func(e int) bool {
		return e == 0
	}).String(), "[:0:]")
	eq(t, la1.TakeWhile(func(e int) bool {
		return e < 2
	}).String(), "[:0, 1:]")
	eq(t, la1.TakeWhile(func(e int) bool {
		return e > 2
	}).String(), "[::]")
	eq(t, la0.DropWhile(func(e int) bool {
		return e == 0
	}).String(), "[::]")
	eq(t, la0.Cons(0).DropWhile(func(e int) bool {
		return e == 0
	}).String(), "[::]")
	eq(t, la1.DropWhile(func(e int) bool {
		return e < 2
	}).String(), "[:2, 3, 4:]")
	eq(t, la1.DropWhile(func(e int) bool {
		return e > 2
	}).String(), "[:0, 1, 2, 3, 4:]")

	eq(t, la0.Sort(func(e1, e2 int) bool {
		return e1 < e2
	}).String(), "[::]")
	asf := la1.Reverse().ToArr()
	arr.Shuffle(asf)
	eq(t, lst.NewFromArr(asf).Sort(func(e1, e2 int) bool {
		return e1 < e2
	}).String(), "[:0, 1, 2, 3, 4:]")

	eq(t, lst.All(la0, 3), true)
	eq(t, lst.All(la1, 3), false)
	eq(t, lst.All(la1.Take(1), 0), true)
	eq(t, la0.Allf(func(e int) bool {
		return e < 3
	}), true)
	eq(t, la0.Allf(func(e int) bool {
		return e < 32
	}), true)
	eq(t, la1.Allf(func(e int) bool {
		return e < 3
	}), false)
	eq(t, lst.Any(la0, 3), false)
	eq(t, lst.Any(la1, 3), true)
	eq(t, lst.Any(la1, 33), false)
	eq(t, la0.Anyf(func(e int) bool {
		return e < 3
	}), false)
	eq(t, la1.Anyf(func(e int) bool {
		return e < 32
	}), true)
	eq(t, la1.Anyf(func(e int) bool {
		return e < 3
	}), true)
	eq(t, la1.Anyf(func(e int) bool {
		return e < -3
	}), false)

	eq(t, lst.Index(la1, 4), 4)
	eq(t, lst.Index(la1, 5), -1)
	eq(t, la0.Indexf(func(e int) bool {
		return e < 3
	}), -1)
	eq(t, la1.Indexf(func(e int) bool {
		return e < 32
	}), 0)
	eq(t, la1.Indexf(func(e int) bool {
		return e == 4
	}), 4)
	eq(t, la1.Indexf(func(e int) bool {
		return e == 5
	}), -1)

	_, ok := la0.Find(func(e int) bool {
		return e < 3
	})
	eq(t, ok, false)
	var el int
	el, ok = la1.Find(func(e int) bool {
		return e < 32
	})
	eq(t, ok, true)
	eq(t, el, 0)
	el, ok = la1.Find(func(e int) bool {
		return e == 4
	})
	eq(t, ok, true)
	eq(t, el, 4)
	el, ok = la1.Find(func(e int) bool {
		return e == 5
	})
	eq(t, ok, false)

	fnAdd := func(seed, n int) int {
		return seed + n
	}

	sum := 0
	la0.Each(func(e int) {
		sum += e
	})
	eq(t, sum, lst.Reduce(la0, 0, fnAdd))
	la1.Each(func(e int) {
		sum += e
	})
	eq(t, sum, lst.Reduce(la1, 0, fnAdd))
	eq(t, sum, 10)
	sum = 0
	la1.EachIx(func(e, ix int) {
		sum += ix * 2
	})
	eq(t, sum, lst.Reduce(la1, 0, fnAdd)*2)
	eq(t, sum, 20)

	eq(t, lst.NewSplit("", "").Count(), 0)
	eq(t, lst.Join(lst.NewSplit("", ""), ""), "")
	eq(t, lst.NewSplit("a", "").Count(), 1)
	eq(t, lst.Join(lst.NewSplit("a", ""), ""), "a")
	eq(t, lst.NewSplit("añ", "").Count(), 2)
	eq(t, lst.Join(lst.NewSplit("añ", ""), ""), "añ")
	eq(t, lst.NewSplit("", ";").Count(), 1)
	eq(t, lst.Join(lst.NewSplit("", ";"), ";"), "")
	eq(t, lst.NewSplit("ab;cd;", ";").Count(), 3)
	eq(t, lst.Join(lst.NewSplit("ab;cd;", ";"), ";"), "ab;cd;")
	eq(t, lst.NewSplit("ab;cd", ";").Count(), 2)
	eq(t, lst.Join(lst.NewSplit("ab;cd", ";"), ";"), "ab;cd")
	eq(t, lst.NewSplit("ab;", ";").Count(), 2)
	eq(t, lst.Join(lst.NewSplit("ab;", ";"), ";"), "ab;")
	eq(t, lst.NewSplit("ab", ";").Count(), 1)
	eq(t, lst.Join(lst.NewSplit("ab", ";"), ";"), "ab")
	eq(t, lst.NewSplit("", "ñ").Count(), 1)
	eq(t, lst.Join(lst.NewSplit("", "ñ"), "ñ"), "")
	eq(t, lst.NewSplit("abñcdñ", "ñ").Count(), 3)
	eq(t, lst.Join(lst.NewSplit("abñcdñ", "ñ"), "ñ"), "abñcdñ")
	eq(t, lst.NewSplit("abñcd", "ñ").Count(), 2)
	eq(t, lst.Join(lst.NewSplit("abñcd", "ñ"), "ñ"), "abñcd")
	eq(t, lst.NewSplit("abñ", "ñ").Count(), 2)
	eq(t, lst.Join(lst.NewSplit("abñ", "ñ"), "ñ"), "abñ")
	eq(t, lst.NewSplit("ab", "ñ").Count(), 1)
	eq(t, lst.Join(lst.NewSplit("ab", "ñ"), "ñ"), "ab")
	eq(t, lst.NewSplit("", "--").Count(), 1)
	eq(t, lst.Join(lst.NewSplit("", "--"), "--"), "")
	eq(t, lst.NewSplit("ab--cd--", "--").Count(), 3)
	eq(t, lst.Join(lst.NewSplit("ab--cd--", "--"), "--"), "ab--cd--")
	eq(t, lst.NewSplit("ab--cd", "--").Count(), 2)
	eq(t, lst.Join(lst.NewSplit("ab--cd", "--"), "--"), "ab--cd")
	eq(t, lst.NewSplit("ab--", "--").Count(), 2)
	eq(t, lst.Join(lst.NewSplit("ab--", "--"), "--"), "ab--")
	eq(t, lst.NewSplit("ab", "--").Count(), 1)
	eq(t, lst.Join(lst.NewSplit("ab", "--"), "--"), "ab")

	eq(t, lst.NewSplitTrim(" a  ", ";").String(), "[:a:]")
	eq(t, lst.NewSplitTrim("   a  ;   b ;c", ";").String(), "[:a, b, c:]")

	eq(t, lst.NewRange0(-2).String(), "[::]")
	eq(t, lst.NewRange0(5).String(), "[:0, 1, 2, 3, 4:]")

	eq(t, lst.Eq(lst.NewRange0(0), lst.NewRange0(0)), true)
	eq(t, lst.Eq(lst.NewRange0(3), lst.NewRange0(0)), false)
	eq(t, lst.Eq(lst.NewRange0(0), lst.NewRange0(3)), false)
	eq(t, lst.Eq(lst.NewRange0(3), lst.NewRange0(3)), true)
	eq(t, lst.Eq(lst.NewRange(1, 3), lst.NewRange0(0)), false)
	eq(t, lst.Eq(lst.NewRange(0, 2), lst.NewRange0(3)), false)

	fnEq := func(n1, n2 int) bool {
		return n1 == n2
	}

	eq(t, lst.NewRange0(0).Eqf(lst.NewRange0(0), fnEq), true)
	eq(t, lst.NewRange0(3).Eqf(lst.NewRange0(0), fnEq), false)
	eq(t, lst.NewRange0(0).Eqf(lst.NewRange0(3), fnEq), false)
	eq(t, lst.NewRange0(3).Eqf(lst.NewRange0(3), fnEq), true)
	eq(t, lst.NewRange(1, 3).Eqf(lst.NewRange0(0), fnEq), false)
	eq(t, lst.NewRange(0, 2).Eqf(lst.NewRange0(0), fnEq), false)
	eq(t, lst.Eq(lst.NewRangeInf().Take(12), lst.NewRange0(12)), true)

	els, dup := lst.Duplicates(
		lst.NewFromArr([]int{1, 2, 1, 0, 3, 3, 4, -1, -1, -1, 100}))
	eq(t, lst.Eq(els, lst.NewFromArr([]int{1, 2, 0, 3, 4, -1, 100}).Reverse()), true)
	eq(t, lst.Eq(dup, lst.NewFromArr([]int{1, 3, -1, -1}).Reverse()), true)

	els, dup = lst.NewFromArr([]int{1, 2, 1, 0, 3, 3, 4, -1, -1, -1, 100}).
		Duplicatesf(fnEq)
	eq(t, lst.Eq(els, lst.NewFromArr([]int{1, 2, 0, 3, 4, -1, 100}).Reverse()), true)
	eq(t, lst.Eq(dup, lst.NewFromArr([]int{1, 3, -1, -1}).Reverse()), true)

}
