// Copyright 25-May-2020 ºDeme
// GNU General Public License - V3 <http://www.gnu.org/licenses/>

package _tests

import (
	"fmt"
	"github.com/dedeme/ktlib/sys"
	"testing"
)

func eq[T comparable](t *testing.T, actual, expected T) {
	if actual != expected {
		t.Fatal(sys.Fail(fmt.Sprintf(
			"\nActual  : %v\nExpected: %v\n", actual, expected)))
	}
}
