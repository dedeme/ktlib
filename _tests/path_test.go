// Copyright 29-May-2022 ºDeme
// GNU General Public License - V3 <http://www.gnu.org/licenses/>

package _tests

import (
	"github.com/dedeme/ktlib/path"
	"testing"
)

func TestPath(t *testing.T) {
	eq(t, path.Base(""), ".")
	eq(t, path.Base("ab.c"), "ab.c")
	eq(t, path.Base("p/ab.c"), "ab.c")
	eq(t, path.Base("/ab.c"), "ab.c")
	eq(t, path.Base("/p/ab.c"), "ab.c")
	eq(t, path.Base("ab"), "ab")
	eq(t, path.Base("p/ab"), "ab")
	eq(t, path.Base("/ab"), "ab")
	eq(t, path.Base("/p/ab"), "ab")

	eq(t, path.Extension(""), "")
	eq(t, path.Extension("ab.c"), ".c")
	eq(t, path.Extension("p/ab.c"), ".c")
	eq(t, path.Extension("/ab.c"), ".c")
	eq(t, path.Extension("/p/ab.c"), ".c")
	eq(t, path.Extension("ab"), "")
	eq(t, path.Extension("p/ab"), "")
	eq(t, path.Extension("/ab"), "")
	eq(t, path.Extension("/p/ab"), "")

	eq(t, path.Parent(""), ".")
	eq(t, path.Parent("ab.c"), ".")
	eq(t, path.Parent("p/ab.c"), "p")
	eq(t, path.Parent("/ab.c"), "/")
	eq(t, path.Parent("/p/ab.c"), "/p")
	eq(t, path.Parent("ab"), ".")
	eq(t, path.Parent("p/ab"), "p")
	eq(t, path.Parent("/ab"), "/")
	eq(t, path.Parent("/p/ab"), "/p")

	eq(t, path.Cat("s"), "s")
	eq(t, path.Cat("s", "b"), "s/b")
	eq(t, path.Cat("/s", "b"), "/s/b")
	eq(t, path.Cat("/s", "/b"), "/s/b")
	eq(t, path.Cat("/s/x", "../b"), "/s/b")

	eq(t, path.Canonical(""), ".")
	eq(t, path.Canonical("///"), "/")
	eq(t, path.Canonical("a/b"), "a/b")
	eq(t, path.Canonical("a/b/"), "a/b")
	eq(t, path.Canonical("a////b"), "a/b")
	eq(t, path.Canonical("a/x/../b"), "a/b")
	eq(t, path.Canonical("/a////b"), "/a/b")
	eq(t, path.Canonical("/a/x/../b"), "/a/b")
}
