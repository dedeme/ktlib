// Copyright 26-May-2020 ºDeme
// GNU General Public License - V3 <http://www.gnu.org/licenses/>

package _tests

import (
	"github.com/dedeme/ktlib/sys"
	"testing"
)

func TestSys(t *testing.T) {
	eq(t, len(sys.Environ()["USER"]) > 0, true)
}
