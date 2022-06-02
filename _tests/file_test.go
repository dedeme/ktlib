// Copyright 26-May-2022 ºDeme
// GNU General Public License - V3 <http://www.gnu.org/licenses/>

package _tests

import (
	"github.com/dedeme/ktlib/arr"
	"github.com/dedeme/ktlib/file"
	"github.com/dedeme/ktlib/path"
	"github.com/dedeme/ktlib/str"
	"testing"
)

func TestFile(t *testing.T) {
	pdf := "db/physics.pdf"
	work := "db/work"
	work2 := "db/work2"
	eq(t, file.IsDirectory("db"), true)
	file.Del(work)
	eq(t, !file.IsDirectory(work), true)
	eq(t, file.Exists(pdf), true)
	eq(t, !file.Exists("db/physic.pdf"), true)
	eq(t, file.Size(pdf), 20298536)
	eq(t, file.Tm(pdf) >= 1545549713565, true)

	file.Mkdir(work)
	eq(t, file.IsDirectory(work), true)
	file.Copy(pdf, work)
	eq(t, file.Exists(path.Cat(work, "physics.pdf")), true)
	file.Rename(path.Cat(work, "physics.pdf"), path.Cat(work, "ph.pdf"))

	file.Write(path.Cat(work, "tx1.txt"), "A text\nwith\nfour\nlines.")
	fw1 := file.Wopen(path.Cat(work, "tx2.txt"))
	file.WriteText(fw1, "A text\nwith")
	file.Close(fw1)
	fa1 := file.Aopen(path.Cat(work, "tx2.txt"))
	file.WriteText(fa1, "\nfour\nlines.")
	file.Close(fa1)
	fw2 := file.Wopen(path.Cat(work, "tx3.txt"))
	file.WriteText(fw2, "A text\nwith")
	file.Close(fw2)
	fa2 := file.Aopen(path.Cat(work, "tx3.txt"))
	file.WriteBin(fa2, []byte("\nfour\nlines."))
	file.Close(fa2)

	tx1 := file.Read(path.Cat(work, "tx1.txt"))

	fr1 := file.Ropen(path.Cat(work, "tx2.txt"))
	var tx2 []string
	for {
		if l, ok := file.ReadLine(fr1); ok {
			arr.Push(&tx2, l)
			continue
		}
		break
	}
	file.Close(fr1)
	eq(t, tx1, arr.Join(tx2, "\n"))

	fr2 := file.Ropen(path.Cat(work, "tx3.txt"))
	var bs []byte
	for {
		if bs2 := file.ReadBin(fr2, 4); len(bs2) > 0 {
			arr.Cat(&bs, bs2)
			continue
		}
		break
	}
	file.Close(fr2)
	eq(t, tx1, string(bs))

	eq(t, arr.Join(file.Dir(work), "-"), "ph.pdf-tx1.txt-tx2.txt-tx3.txt")

	file.Del(work2)
	file.Copy(work, work2)

	//t.Fatal(file.Home())
	tmp1 := file.Tmp("", "abc")
	eq(t, str.Starts(tmp1, "abc") && len(tmp1) == 11, true)
	tmp2 := file.Tmp("/tmp", "abc")
	eq(t, str.Starts(tmp2, "/tmp/abc") && len(tmp2) == 16, true)
	eq(t, str.Ends(file.Wd(), "_tests"), true)
	file.Cd("db/work2")
	eq(t, str.Ends(file.Wd(), "work2"), true)
	file.Cd("../../")
	eq(t, str.Ends(file.Wd(), "_tests"), true)
	eq(t, file.IsDirectory("db"), true)
}
