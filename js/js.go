// Copyright 29-May-2022 ºDeme
// GNU General Public License - V3 <http://www.gnu.org/licenses/>

// JSON utilities.
package js

import (
	"encoding/json"
	"fmt"
	"github.com/dedeme/ktlib/math"
	"strconv"
	"strings"
)

// Writes a JSON null value.
func Wn() string {
	return "null"
}

// Returns 'true' if json is "null" or 'false' in another case.
func IsNull(j string) bool {
	return string(j) == "null"
}

// Writes a JSON boolean value.
func Wb(v bool) string {
	j, err := json.Marshal(v)
	if err != nil {
		panic(err)
	}
	return string(j)
}

// Reads a JSON boolean value.
func Rb(j string) (v bool) {
	err := json.Unmarshal([]byte(j), &v)
	if err != nil {
		panic(err)
	}
	return
}

// Writes a JSON int value.
func Wi(v int) string {
	j, err := json.Marshal(v)
	if err != nil {
		panic(err)
	}
	return string(j)
}

// Reads a JSON int value.
func Ri(j string) (v int) {
	err := json.Unmarshal([]byte(j), &v)
	if err != nil {
		panic(err)
	}
	return
}

// Writes a JSON int64 value.
func Wl(v int64) string {
	j, err := json.Marshal(v)
	if err != nil {
		panic(err)
	}
	return string(j)
}

// Reads a JSON int64 value.
func Rl(j string) (v int64) {
	err := json.Unmarshal([]byte(j), &v)
	if err != nil {
		panic(err)
	}
	return
}

// Writes a JSON float32 value.
func Wf(v float32) string {
	j, err := json.Marshal(v)
	if err != nil {
		panic(err)
	}
	return string(j)
}

// Reads a JSON float32 value.
func Rf(j string) (v float32) {
	err := json.Unmarshal([]byte(j), &v)
	if err != nil {
		panic(err)
	}
	return
}

// Writes a JSON float64 value.
func Wd(v float64) string {
	j, err := json.Marshal(v)
	if err != nil {
		panic(err)
	}
	return string(j)
}

// Writes a JSON float64 value with 'dec' decimals (between 0 and 9, both inclusive).
func WdDec(v float64, dec int) string {
	fm := "%." + strconv.Itoa(dec) + "f"
	return fmt.Sprintf(fm, math.Round(v, dec))
}

// Reads a JSON float64 value.
func Rd(j string) (v float64) {
	err := json.Unmarshal([]byte(j), &v)
	if err != nil {
		panic(err)
	}
	return
}

// Writes a JSON string value.
func Ws(v string) string {
	j, err := json.Marshal(v)
	if err != nil {
		panic(err)
	}
	return string(j)
}

// Reads a JSON string value.
func Rs(j string) (v string) {
	err := json.Unmarshal([]byte(j), &v)
	if err != nil {
		panic(err)
	}
	return
}

// Writes a JSON string value from a slice of js-strings.
func Wa(v []string) string {
	var b strings.Builder
	b.WriteByte('[')
	for i, j := range v {
		if i > 0 {
			b.WriteByte(',')
		}
		b.WriteString(string(j))
	}
	b.WriteByte(']')
	return b.String()
}

// Reads a JSON slice of js-string values.
func Ra(j string) (v []string) {
	if !strings.HasPrefix(j, "[") {
		panic(fmt.Sprintf("Array does not start with '[' in\n'%v'", j))
	}
	if !strings.HasSuffix(j, "]") {
		panic(fmt.Sprintf("Array does not end with ']' in\n'%v'", j))
	}
	s2 := strings.TrimSpace(j[1 : len(j)-1])
	l := len(s2)
	if l == 0 {
		return
	}
	i := 0
	var e string
	for {
		if i2, ok := nextByte(s2, ',', i); ok {
			e = strings.TrimSpace(s2[i:i2])
			if e == "" {
				panic(fmt.Sprintf("Missing elements in\n'%v'", j))
			}
			v = append(v, e)
			i = i2 + 1
			continue
		}
		e = strings.TrimSpace(s2[i:l])
		if e == "" {
			panic(fmt.Sprintf("Missing elements in\n'%v'", j))
		}
		v = append(v, e)
		break
	}
	return
}

// Writes a JSON string value from a map of [string]js-strings.
func Wo(v map[string]string) string {
	var b strings.Builder
	b.WriteByte('{')
	more := false
	for k, j := range v {
		if more {
			b.WriteByte(',')
		} else {
			more = true
		}
		b.WriteString(string(Ws(k)))
		b.WriteByte(':')
		b.WriteString(j)
	}
	b.WriteByte('}')
	return b.String()
}

// Reads a JSON object (map of string[js-string] values).
func Ro(j string) (v map[string]string) {
	v = make(map[string]string)
	if !strings.HasPrefix(j, "{") {
		panic(fmt.Sprintf("Object does not start with '{' in\n'%v'", j))
	}
	if !strings.HasSuffix(j, "}") {
		panic(fmt.Sprintf("Object does not end with '}' in\n'%v'", j))
	}
	s2 := strings.TrimSpace(j[1 : len(j)-1])
	l := len(s2)
	if l == 0 {
		return
	}
	i := 0
	var kjs string
	var k string
	var val string
	for {
		i2, ok := nextByte(s2, ':', i)
		if !ok {
			panic(fmt.Sprintf("Expected ':' in\n'%v'", s2))
		}
		kjs = strings.TrimSpace(s2[i:i2])
		if kjs == "" {
			panic(fmt.Sprintf("Key missing in\n'%v'", j))
		}
		k = Rs(kjs)

		i = i2 + 1

		if i2, ok := nextByte(s2, ',', i); ok {
			val = strings.TrimSpace(s2[i:i2])
			if val == "" {
				panic(fmt.Sprintf("Value missing in\n'%v'", j))
			}
			v[k] = val
			i = i2 + 1
			continue
		}
		val = strings.TrimSpace(s2[i:l])
		if val == "" {
			panic(fmt.Sprintf("Value missing in\n'%v'", j))
		}
		v[k] = val
		break
	}
	return
}
