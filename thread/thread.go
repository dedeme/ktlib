// Copyright 31-May-2020 ºDeme
// GNU General Public License - V3 <http://www.gnu.org/licenses/>

// Threads management.
package thread

import (
	"sync"
)

var mtx sync.Mutex

// Stops program until the thread with the channel 'ch' is finished.
func Join(ch chan bool) {
	<-ch
}

// Runs 'fn' in a detached thread.
func Run(fn func()) {
	go fn()
}

// Runs 'fn' in a thread with a chanel that is returned. This chanel can be used
// with 'thread.Join' for waiting the tread to end.
func Start(fn func()) chan bool {
	r := make(chan bool)
	go func() {
		fn()
		r <- true
	}()
	return r
}

// Runs the function 'fn' avoiding that another thread accesses it before
// its ending.
//
// It is necessary prevent that 'fn' directly or indirectly call also
// 'thread.Sync'. Otherwise the program will be blocked.
func Sync(fn func()) {
	mtx.Lock()
	fn()
	mtx.Unlock()
}
