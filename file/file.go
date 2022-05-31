// Copyright 26-May-2017 ºDeme
// GNU General Public License - V3 <http://www.gnu.org/licenses/>

// Functions to manage file system.
package file

import (
	"bufio"
	"crypto/rand"
	"github.com/dedeme/ktlib/b64"
	"github.com/dedeme/ktlib/path"
	"github.com/dedeme/ktlib/str"
	"io"
	"io/ioutil"
	"os"
	"os/user"
)

type T struct {
	f    *os.File
	scan *bufio.Scanner
}

// Open 'fpath' for appending.
func Aopen(fpath string) *T {
	f, err := os.OpenFile(fpath, os.O_APPEND|os.O_WRONLY, 0755)
	if err != nil {
		panic(err)
	}
	return &T{f, nil}
}

// Changes working directory.
func Cd(dir string) {
	os.Chdir(dir)
}

// Closes 'f' and frees resources.
func Close(f *T) {
	f.f.Close()
}

// Copy a file.
//    source: Can be a regular file or a directory.
//    target: - If source is a directory, target must be the target parent
//              directory.
//            - If source is a regular file, target can be the target parent
//              directory or a regular file.
// NOTE: Target will be overwritte if already exists.
func Copy(source, target string) {
	if IsDirectory(source) {
		p := path.Cat(target, path.Base(source))
		if Exists(p) {
			if !IsDirectory(p) {
				panic(str.Fmt(
					"Copying '%v' to '%v', when the later exists and is not a directory",
					source, p,
				))
			}
			Del(p)
		}
		Mkdir(p)

		for _, e := range Dir(source) {
			Copy(path.Cat(source, e), p)
		}
		return
	}

	if IsDirectory(target) {
		target = path.Cat(target, path.Base(source))
	}

	sourcef, err := os.Open(source)
	if err != nil {
		panic(err)
	}

	targetf, err := os.Create(target)
	if err != nil {
		panic(err)
	}

	defer sourcef.Close()
	defer targetf.Close()

	buf := make([]byte, 8192)
	for {
		n, err := sourcef.Read(buf)
		if err != nil && err != io.EOF {
			panic(err)
		}
		if n == 0 {
			break
		}

		if _, err := targetf.Write(buf[:n]); err != nil {
			panic(err)
		}
	}

	return
}

// Deletes 'fpath'. If 'fpath' is a directory, 'del' deletes it and all its
// subdirectories, although they are not empty.
//
// If 'fpath' does not exists, 'del' do nothing.
func Del(fpath string) {
	os.RemoveAll(fpath)
}

// Returns base names of files and directories of 'dir'.
func Dir(fpath string) []string {
	fis, err := ioutil.ReadDir(fpath)
	if err != nil {
		panic(err)
	}
	var r []string
	for _, fi := range fis {
		r = append(r, fi.Name())
	}
	return r
}

// Returns 'true' if 'fpath' is a file or directory.
func Exists(fpath string) bool {
	if _, err := os.Stat(fpath); err == nil {
		return true
	}
	return false
}

// Returns home directory.
func Home() string {
	u, err := user.Current()
	if err != nil {
		panic(err)
	}
	return u.HomeDir
}

// Returns 'true' if 'fpath' is a directory.
func IsDirectory(fpath string) bool {
	if info, err := os.Stat(fpath); err == nil && info.IsDir() {
		return true
	}
	return false
}

// Creates the directory 'fpath' and its parents directory if is necessary.
func Mkdir(fpath string) {
	os.MkdirAll(fpath, os.FileMode(0755))
}

// Read the complete file 'fpath'.
func Read(fpath string) string {
	bs, err := ioutil.ReadFile(fpath)
	if err != nil {
		panic(err)
	}
	return string(bs)
}

// Read 'buf' bytes of 'f'.
//    f  : An open file.
//    buf: Bytes buffer.
// When 'f' is at end, it returns a bytes with no element.
func ReadBin(f *T, buf int) []byte {
	bs := make([]byte, buf)
	n, err := f.f.Read(bs)
	if n == 0 && err == io.EOF {
		err = nil
	}
	if err != nil {
		panic(err)
	}
	if n == buf {
		return bs
	}
	bs2 := make([]byte, n)
	for i := 0; i < n; i++ {
		bs2[i] = bs[i]
	}
	return bs2
}

// Returns the next line of 'f' or 'ok=false' if there is no more elements.
//    f: An open file.
// Lines are read without carriage return.
func ReadLine(f *T) (tx string, ok bool) {
	if f.scan.Scan() {
		tx = f.scan.Text()
		ok = true
	}
	return
}

// Changes (moves) 'oldPath' by 'newPath'
func Rename(oldPath, newPath string) {
	err := os.Rename(oldPath, newPath)
	if err != nil {
		panic(err)
	}
}

// Open 'fpath' for reading.
func Ropen(fpath string) *T {
	f, err := os.Open(fpath)
	if err != nil {
		panic(err)
	}
	return &T{f, bufio.NewScanner(f)}
}

// Returns the size of file 'fpath'.
func Size(fpath string) int64 {
	info, err := os.Stat(fpath)
	if err != nil {
		panic(err)
	}
	return info.Size()
}

// Returns the time of the last modification of file 'fpath'.
func Tm(fpath string) int64 {
	info, err := os.Stat(fpath)
	if err != nil {
		panic(err)
	}
	return info.ModTime().UnixMilli()
}

// Returns a random name for a file in "dir" + "/" + 'fpath' + "xxxxxxxx"
// ('sys.Rand()' should be called previously).
//
// The return is checked that does not match any existing file.
func Tmp(dir, fpath string) (tmpPath string) {
	for {
		lg := 8
		a := make([]byte, lg)
		_, err := rand.Read(a)
		if err != nil {
			panic(err)
		}

		tmpPath = path.Cat(
			dir,
			fpath+str.Replace(b64.EncodeBytes(a)[:lg], "/", "+"),
		)
		if _, err := os.Stat(tmpPath); err != nil {
			break
		}
	}
	return
}

// Returns working directory.
func Wd() string {
	d, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	return d
}

// Open 'fpath' for writing.
func Wopen(fpath string) *T {
	f, err := os.Create(fpath)
	if err != nil {
		panic(err)
	}
	return &T{f, nil}
}

// Creates, writes and closes 'fpath'.
func Write(fpath, tx string) {
	err := ioutil.WriteFile(fpath, []byte(tx), 0755)
	if err != nil {
		panic(err)
	}
}

// Writes 'bs' in 'f'.
func WriteBin(f *T, bs []byte) {
	_, err := f.f.Write(bs)
	if err != nil {
		panic(err)
	}
}

// Writes 'tx' in 'f'.
func WriteText(f *T, tx string) {
	_, err := f.f.WriteString(tx)
	if err != nil {
		panic(err)
	}
}
