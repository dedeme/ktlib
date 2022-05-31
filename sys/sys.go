// Copyright 25-May-2020 ºDeme
// GNU General Public License - V3 <http://www.gnu.org/licenses/>

// System utilities.
package sys

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"math/rand"
	"os"
	"os/exec"
	"runtime/debug"
	"strings"
	"time"
)

var rndV int64

// Adds the stack trace to 'msg'.
func Fail(msg string) string {
	return msg + "\n" + string(debug.Stack())
}

// Returns program arguments.
func Args() []string {
	return os.Args
}

// Executes 'Cmd' and returns its stdout and stderror results.
// If an error happens, it panics.
//    c   : Command.
//    args: Arguments.
//    RETURN ------
//    stdout : stdout response.
//    stderr : stderr response.
func Cmd(c string, args ...string) (stdout, stderr string) {
	cmd := exec.Command(c, args...)
	sOut, err := cmd.StdoutPipe()
	if err != nil {
		panic(err)
	}
	sErr, err := cmd.StderrPipe()
	if err != nil {
		panic(err)
	}
	err = cmd.Start()
	if err != nil {
		panic(err)
	}

	bytesOut, err := ioutil.ReadAll(sOut)
	if err != nil {
		panic(err)
	}
	stdout = string(bytesOut)

	bytesErr, err := ioutil.ReadAll(sErr)
	if err != nil {
		panic(err)
	}
	stderr = string(bytesErr)

	err = cmd.Wait()
	if err != nil && len(bytesErr) == 0 {
		stderr = err.Error()
	}

	return
}

// Returns EVIRONMENT values
func Environ() map[string]string {
	r := map[string]string{}
	for _, e := range os.Environ() {
		ix := strings.IndexByte(e, '=')
		if ix != -1 {
			r[e[:ix]] = e[ix+1:]
		}
	}
	return r
}

// Shows 'o'
func Print[T any](o T) {
	fmt.Print(o)
}

// Shows 'o' + "\n"
func Println[T any](o T) {
	fmt.Println(o)
}

// Starts the random generator of numbers.
// This functions must be called previously to use random functions.
func Rand() {
	rndV += time.Now().UnixMilli()
	rand.Seed(rndV)
	return
}

// Reads a line on Console.
// The '\n' byte is not read.
func ReadLine() string {
	rd := bufio.NewReader(os.Stdin)
	s, err := rd.ReadString('\n')
	if err != nil && err != io.EOF {
		panic(err)
	}
	return s[:len(s)-1]
}

// Stop the current thread 'millis' milliseconds.
func Sleep(millis int) {
	time.Sleep(time.Duration(millis) * time.Millisecond)
}
