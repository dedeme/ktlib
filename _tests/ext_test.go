// Copyright 01-Jun-2022 ºDeme
// GNU General Public License - V3 <http://www.gnu.org/licenses/>

package _tests

import (
	"github.com/dedeme/ktlib/ext"
	"github.com/dedeme/ktlib/file"
	"testing"
)

func TestExt(t *testing.T) {
	eq(t, ext.Md5Str("ca labaza"), "a544eed57bd849378bd796dc099dca48")
	eq(t, ext.Sha256Str("ca labaza"),
		"91d3e225b58dcc7d76322cccc3a0a678a972a57ec15078d0b9e6ecab6cf8497a")

	//t.Fatal(ext.Wget("http://www.google.es", true))

	zdir := "db/z"
	zdir2 := "db/dtg"
	zfile := "db/z/ztest.txt"
	zfile2 := "db/dtg/z/ztest.txt"
	file.Mkdir(zdir)
	file.Mkdir(zdir2)
	file.Write(zfile, "ab\nc\n")
	ext.Zip(zdir, "z.zip")
	ext.Unzip("db/z.zip", zdir2)
	eq(t, file.Read(zfile), file.Read(zfile2))
	file.Del(zdir)
	file.Del(zdir2)
	file.Del("db/z.zip")

	file.Write(
		"db/HtmlDocTest.pdf",
		string(ext.HtmldocStr(false, "<p><b>H</b>ello<br>€ cañón</p>")),
	)
	eq(t, file.Exists("db/HtmlDocTest.pdf"), true)
	file.Del("db/HtmlDocTest.pdf")
}
