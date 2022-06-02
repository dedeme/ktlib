// Copyright 29-May-2022 ºDeme
// GNU General Public License - V3 <http://www.gnu.org/licenses/>

package _tests

import (
	"github.com/dedeme/ktlib/time"
	"testing"
)

func TestTime(t *testing.T) {
	d0 := time.NewDate(12, 2, 2022)
	eq(t, time.ToStr(d0), "20220212")
	eq(t, time.ToIso(d0), "12/02/2022")
	eq(t, time.ToEn(d0), "02-12-2022")

	d1 := time.New(7, 2, 2022, 11, 45, 0)
	eq(t, time.ToStr(d1)+"-"+time.Fmt("%t", d1), "20220207-11:45:00")
	eq(t, time.FromClock(time.AddDays(d0, -5), "11:45:00"), d1)

	eq(t, time.FromStr("20220212"), d0)
	eq(t, time.FromIso("12/02/2022", "/"), d0)
	eq(t, time.FromEn("02-12-2022", "-"), d0)

	eq(t, time.DfDays(d0, d1), 5)
	eq(t, time.EqDay(time.AddDays(d1, 5), d0), true)

	eq(t, time.Weekday(d0), 6)
	eq(t, time.YearDay(d0), 43)

	eq(t, time.Day(d1), 7)
	eq(t, time.Month(d1), 2)
	eq(t, time.Year(d1), 2022)
	eq(t, time.Hour(d1), 11)
	eq(t, time.Minute(d1), 45)
	eq(t, time.Second(d1), 0)

	//t.Fatal(time.Fmt("%D%M%Y-%T", time.Now()));
}
