// Copyright 25-May-2020 ºDeme
// GNU General Public License - V3 <http://www.gnu.org/licenses/>

package _tests

import (
	"github.com/dedeme/ktlib/arr"
	"github.com/dedeme/ktlib/sys"
	"github.com/dedeme/ktlib/thread"
	"testing"
)

func TestThread(t *testing.T) {
	a := []int{}
	populate := func() {
		for i := 1; i < 10; i++ {
			arr.Push(&a, i)
			sys.Sleep(10)
		}
	}

	thread.Run(populate)
	sys.Sleep(5)
	thread.Run(populate)

	sys.Sleep(500)
	eq(t, arr.ToStr(a), "[1, 1, 2, 2, 3, 3, 4, 4, 5, 5, 6, 6, 7, 7, 8, 8, 9, 9]")

	arr.Clear(&a)
	th1 := thread.Start(populate)
	sys.Sleep(5)
	th2 := thread.Start(populate)

	thread.Join(th2)
	thread.Join(th1)
	eq(t, arr.ToStr(a), "[1, 1, 2, 2, 3, 3, 4, 4, 5, 5, 6, 6, 7, 7, 8, 8, 9, 9]")

	arr.Clear(&a)
	th1 = thread.Start(func() {
		thread.Sync(populate)
	})
	sys.Sleep(5)
	th2 = thread.Start(func() {
		thread.Sync(populate)
	})

	thread.Join(th2)
	thread.Join(th1)
	eq(t, arr.ToStr(a), "[1, 2, 3, 4, 5, 6, 7, 8, 9, 1, 2, 3, 4, 5, 6, 7, 8, 9]")
}
