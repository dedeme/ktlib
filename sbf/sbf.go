// Copyright 21-Jul-2022 ºDeme
// GNU General Public License - V3 <http://www.gnu.org/licenses/>

// Strings constructor (string buffer)
package sbf

import (
	"strings"
)

type T struct {
	sb strings.Builder
}

// Returns a new string builder.
func New() *T {
	var sb strings.Builder
	return &T{sb}
}

// Returns a new string builder intialized with 's'
func NewStr(s string) *T {
	var r strings.Builder
	r.WriteString(s)
	return &T{r}
}

// Adds a string to 'sb'
func (sb *T) Add(s string) {
	sb.sb.WriteString(s)
}

// Adds a byte so 'sb'
func (sb *T) AddByte(b byte) {
	sb.sb.WriteByte(b)
}

// Returns the bytes number of 'sb'
func (sb *T) Len() int {
	return sb.sb.Len()
}

// Resets 'sb'
func (sb *T) Reset() {
	sb.sb.Reset()
}

// Return the String value of 'sb'
func (sb *T) String() string {
	return sb.sb.String()
}
