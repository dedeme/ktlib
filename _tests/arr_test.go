// Copyright 25-May-2022 ºDeme
// GNU General Public License - V3 <http://www.gnu.org/licenses/>

package _tests

import (
	"github.com/dedeme/ktlib/arr"
	"github.com/dedeme/ktlib/str"
	"github.com/dedeme/ktlib/sys"
	"testing"
)

func TestArr(t *testing.T) {
	a := []int{}
	arr.Cat(&a, []int{})
	eq(t, arr.Empty(a), true)
	arr.Cat(&a, []int{1, 2})
	eq(t, arr.ToStr(a), "[1, 2]")
	arr.Cat(&a, []int{})
	eq(t, arr.ToStr(a), "[1, 2]")
	arr.Cat(&a, []int{3})
	eq(t, arr.ToStr(a), "[1, 2, 3]")

	a0 := []int{}
	a1 := []int{0, 1, 2, 3, 4}

	eq(t, arr.Empty(a0), true)

	eq(t, arr.ToStr(arr.Filter(a0, func(i int) bool {
		return i%2 == 0
	})), "[]")
	eq(t, arr.ToStr(arr.Filter(a1, func(i int) bool {
		return i%2 == 0
	})), "[0, 2, 4]")
	eq(t, arr.ToStr(arr.Filter(a1, func(i int) bool {
		return i%2 != 0
	})), "[1, 3]")
	eq(t, arr.ToStr(arr.Filter(a1, func(i int) bool {
		return i > 1000
	})), "[]")

	ac := arr.Copy(a0)
	arr.FilterIn(&ac, func(i int) bool {
		return i%2 == 0
	})
	eq(t, arr.ToStr(ac), "[]")
	ac = arr.Copy(a1)
	arr.FilterIn(&ac, func(i int) bool {
		return i%2 == 0
	})
	eq(t, arr.ToStr(ac), "[0, 2, 4]")
	ac = arr.Copy(a1)
	arr.FilterIn(&ac, func(i int) bool {
		return i%2 != 0
	})
	eq(t, arr.ToStr(ac), "[1, 3]")
	ac = arr.Copy(a1)
	arr.FilterIn(&ac, func(i int) bool {
		return i > 1000
	})
	eq(t, arr.ToStr(ac), "[]")

	el, ok := arr.Find(a0, func(i int) bool {
		return i == 3
	})
	eq(t, ok, false)
	el, ok = arr.Find(a1, func(i int) bool {
		return i == 3
	})
	eq(t, ok, true)
	eq(t, el, 3)
	el, ok = arr.Find(a1, func(i int) bool {
		return i == -100
	})
	eq(t, ok, false)

	eq(t, arr.Join([]string{}, "-"), "")
	eq(t, arr.Join([]string{"a"}, "-"), "a")
	eq(t, arr.Join([]string{"a", "b", "c"}, "-"), "a-b-c")

	ac = arr.Copy(a1)
	ac[1] = 101
	eq(t, arr.ToStr(ac), "[0, 101, 2, 3, 4]")
	arr.Push(&ac, 33)
	eq(t, arr.ToStr(ac), "[0, 101, 2, 3, 4, 33]")
	arr.Unshift(&ac, -12)
	eq(t, arr.ToStr(ac), "[-12, 0, 101, 2, 3, 4, 33]")
	el = arr.Pop(&ac)
	eq(t, el, 33)
	eq(t, arr.ToStr(ac), "[-12, 0, 101, 2, 3, 4]")
	el = arr.Shift(&ac)
	eq(t, el, -12)
	eq(t, arr.ToStr(ac), "[0, 101, 2, 3, 4]")
	eq(t, arr.Peek(ac), 4)
	eq(t, arr.Peek(ac), 4)

	fnEq := func(n1, n2 int) bool {
		return n1 == n2
	}

	eq(t, arr.Eq(ac, a1), false)
	eq(t, arr.Eq(ac[2:], a1[2:]), true)
	eq(t, arr.Eq(ac[3:], a1[2:]), false)
	eq(t, arr.Eqf(ac, a1, fnEq), false)
	eq(t, arr.Eqf(ac[2:], a1[2:], fnEq), true)
	eq(t, arr.Eqf(ac[3:], a1[2:], fnEq), false)
	arr.Clear(&ac)
	eq(t, arr.Eq(ac, a0), true)
	eq(t, arr.Eqf(ac, a0, fnEq), true)

	els, dup := arr.Duplicates([]int{1, 2, 1, 0, 3, 3, 4, -1, -1, -1, 100})
	eq(t, arr.Eq(els, []int{1, 2, 0, 3, 4, -1, 100}), true)
	eq(t, arr.Eq(dup, []int{1, 3, -1, -1}), true)

	els, dup = arr.Duplicatesf([]int{1, 2, 1, 0, 3, 3, 4, -1, -1, -1, 100}, fnEq)
	eq(t, arr.Eq(els, []int{1, 2, 0, 3, 4, -1, 100}), true)
	eq(t, arr.Eq(dup, []int{1, 3, -1, -1}), true)

	fnAdd := func(seed, n int) int {
		return seed + n
	}

	sum := 0
	arr.Each(a0, func(e int) {
		sum += e
	})
	eq(t, sum, arr.Reduce(a0, 0, fnAdd))
	arr.Each(a1, func(e int) {
		sum += e
	})
	eq(t, sum, arr.Reduce(a1, 0, fnAdd))
	eq(t, sum, 10)
	sum = 0
	arr.EachIx(a1, func(e, ix int) {
		sum += ix * 2
	})
	eq(t, sum, arr.Reduce(a1, 0, fnAdd)*2)
	eq(t, sum, 20)

	eq(t, arr.All(a0, 3), true)
	eq(t, arr.All(a1, 3), false)
	eq(t, arr.All(arr.Take(a1, 1), 0), true)
	eq(t, arr.Allf(a0, func(e int) bool {
		return e < 3
	}), true)
	eq(t, arr.Allf(a1, func(e int) bool {
		return e < 32
	}), true)
	eq(t, arr.Allf(a1, func(e int) bool {
		return e < 3
	}), false)
	eq(t, arr.Any(a0, 3), false)
	eq(t, arr.Any(a1, 3), true)
	eq(t, arr.Any(a1, 33), false)
	eq(t, arr.Anyf(a0, func(e int) bool {
		return e < 3
	}), false)
	eq(t, arr.Anyf(a1, func(e int) bool {
		return e < 32
	}), true)
	eq(t, arr.Anyf(a1, func(e int) bool {
		return e < 3
	}), true)
	eq(t, arr.Anyf(a1, func(e int) bool {
		return e < -3
	}), false)

	eq(t, arr.Index(a1, 4), 4)
	eq(t, arr.Index(a1, 5), -1)
	eq(t, arr.Indexf(a0, func(e int) bool {
		return e < 3
	}), -1)
	eq(t, arr.Indexf(a1, func(e int) bool {
		return e < 32
	}), 0)
	eq(t, arr.Indexf(a1, func(e int) bool {
		return e == 4
	}), 4)
	eq(t, arr.Indexf(a1, func(e int) bool {
		return e == 5
	}), -1)

	eq(t, arr.ToStr(arr.Take([]string{}, 2)), "[]")
	eq(t, arr.ToStr(arr.Take([]string{"a"}, 2)), "[a]")
	eq(t, arr.ToStr(arr.Take(a1, 2)), "[0, 1]")
	eq(t, arr.ToStr(arr.Take(a1, -2)), "[]")
	eq(t, arr.ToStr(arr.Drop([]string{}, 2)), "[]")
	eq(t, arr.ToStr(arr.Drop([]string{"a"}, 2)), "[]")
	eq(t, arr.ToStr(arr.Drop(a1, 2)), "[2, 3, 4]")
	eq(t, arr.ToStr(arr.Drop(a1, -2)), "[0, 1, 2, 3, 4]")

	eq(t, arr.ToStr(arr.TakeWhile([]string{}, func(e string) bool {
		return e == "a"
	})), "[]")
	eq(t, arr.ToStr(arr.TakeWhile([]string{"a"}, func(e string) bool {
		return e == "a"
	})), "[a]")
	eq(t, arr.ToStr(arr.TakeWhile(a1, func(e int) bool {
		return e < 2
	})), "[0, 1]")
	eq(t, arr.ToStr(arr.TakeWhile(a1, func(e int) bool {
		return e > 2
	})), "[]")
	eq(t, arr.ToStr(arr.DropWhile([]string{}, func(e string) bool {
		return e == "a"
	})), "[]")
	eq(t, arr.ToStr(arr.DropWhile([]string{"a"}, func(e string) bool {
		return e == "a"
	})), "[]")
	eq(t, arr.ToStr(arr.DropWhile(a1, func(e int) bool {
		return e < 2
	})), "[2, 3, 4]")
	eq(t, arr.ToStr(arr.DropWhile(a1, func(e int) bool {
		return e > 2
	})), "[0, 1, 2, 3, 4]")

	ac = arr.Copy(a1)
	arr.Remove(&ac, 2)
	eq(t, arr.Eq(ac, []int{0, 1, 3, 4}), true)

	ac = arr.Copy(a1)
	arr.RemoveRange(&ac, 2, 4)
	eq(t, arr.Eq(ac, []int{0, 1, 4}), true)

	ac = arr.Reverse(a1)
	eq(t, arr.Eq(ac, a1), false)
	eq(t, arr.Eq(ac, []int{4, 3, 2, 1, 0}), true)
	arr.ReverseIn(ac)
	eq(t, arr.Eq(ac, a1), true)
	ac = arr.Reverse(a0)
	eq(t, arr.Eq(ac, a0), true)
	arr.ReverseIn(ac)
	eq(t, arr.Eq(ac, a0), true)

	ac = arr.Reverse(a1)
	arr.Sort(ac, func(e1, e2 int) bool {
		return e1 < e2
	})
	eq(t, arr.Eq(ac, a1), true)
	ac2 := arr.Reverse(a1)
	arr.Sort(ac, func(e1, e2 int) bool {
		return e1 > e2
	})
	eq(t, arr.Eq(ac, ac2), true)

	ac = arr.Copy(a0)
	arr.Sort(ac, func(e1, e2 int) bool {
		return e1 < e2
	})
	eq(t, arr.Eq(ac, a0), true)

	eq(t, arr.Eq(arr.Map(a0, func(e int) string {
		return str.Fmt("%d", e)
	}), []string{}), true)
	eq(t, arr.Eq(arr.Map(a1, func(e int) string {
		return str.Fmt("%d", e)
	}), []string{"0", "1", "2", "3", "4"}), true)

	sys.Rand()
	//  ac = arr.Copy(a1)
	//  arr.Shuffle(ac)
	//  t.Fatal(arr.ToStr(ac))
}
