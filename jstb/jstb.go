// Copyright 01-Jun-2022 ºDeme
// GNU General Public License - V3 <http://www.gnu.org/licenses/>

// JSON table.
package jstb

import (
	"github.com/dedeme/ktlib/file"
)

type T[Tb any] struct {
	fpath  string
	init   Tb
	toJs   func(Tb) string
	fromJs func(string) Tb
}

// Creates a jstb with a JSON serializable object.
//    fpath  : Path of table file.
//    initVal: Initial value for the table.
//    toJs   : Object JSON serialization
//    fromJs : Objset JSON deserialization
func New[Tb any](
	fpath string, initVal Tb, toJs func(Tb) string, fromJs func(string) Tb,
) *T[Tb] {
	if !file.Exists(fpath) {
		file.Write(fpath, toJs(initVal))
	}
	return &T[Tb]{fpath, initVal, toJs, fromJs}
}

// Returns the JSON representation of the saved object.
func (tb T[Tb]) ReadJs() string {
	return file.Read(tb.fpath)
}

// Returns the the saved object.
func (tb T[Tb]) Read() Tb {
	return tb.fromJs(file.Read(tb.fpath))
}

// Writes the JSON representation of the table object in the table file.
func (tb T[Tb]) WriteJs(j string) {
	file.Write(tb.fpath, j)
}

// Writes 'o' in the table file.
func (tb T[Tb]) Write(o Tb) {
	file.Write(tb.fpath, tb.toJs(o))
}
