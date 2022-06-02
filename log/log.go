// Copyright 01-Jun-2022 ºDeme
// GNU General Public License - V3 <http://www.gnu.org/licenses/>

// Log utility.
package log

import (
	"github.com/dedeme/ktlib/arr"
	"github.com/dedeme/ktlib/js"
	"github.com/dedeme/ktlib/jstb"
	"github.com/dedeme/ktlib/time"
)

// Entry of 'log'
type T struct {
	IsError bool
	Time    string
	Msg     string
}

func toJs(e *T) string {
	return js.Wa([]string{
		js.Wb(e.IsError),
		js.Ws(e.Time),
		js.Ws(e.Msg),
	})
}

func fromJs(j string) *T {
	a := js.Ra(j)
	return &T{
		js.Rb(a[0]),
		js.Rs(a[1]),
		js.Rs(a[2]),
	}
}

var tb *jstb.T[[]*T]

// Sets file path for Log.
//
// It is mandatory to call this function previously to call any other of this
// file.
//   fpath  : File absolute path.
func Initialize(fpath string) {
	tb = jstb.New(
		fpath,
		[]*T{},
		func(entries []*T) string {
			return js.Wa(arr.Map(entries, toJs))
		},
		func(j string) []*T {
			return arr.Map(js.Ra(j), fromJs)
		},
	)
}

// Read JSON serialization of 'log'.
func ReadJs() string {
	if tb == nil {
		panic("Log not initialized")
	}

	return tb.ReadJs()
}

// Read 'log'.
func Read() []*T {
	if tb == nil {
		panic("Log not initialized")
	}

	return tb.Read()
}

// Adds a warning message to 'log'.
func Warning(msg string) {
	if tb == nil {
		panic("Log not initialized")
	}

	l := Read()
	l = append(l, &T{
		false,
		time.Fmt("%D/%M/%Y(%t)", time.Now()),
		msg,
	})
	tb.Write(l)
}

// Adds an error message to 'log'.
func Error(msg string) {
	if tb == nil {
		panic("Log not initialized")
	}

	l := Read()
	l = append(l, &T{
		true,
		time.Fmt("%D/%M/%Y(%t)", time.Now()),
		msg,
	})
	tb.Write(l)
}

// Resets 'log'.
func Reset() {
	if tb == nil {
		panic("Log not initialized")
	}

	tb.Write([]*T{})
}
