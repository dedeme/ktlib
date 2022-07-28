// Copyright 01-Jun-2022 ºDeme
// GNU General Public License - V3 <http://www.gnu.org/licenses/>

// Functions using external programs.
package ext

import (
	"github.com/dedeme/ktlib/file"
	"github.com/dedeme/ktlib/js"
	"github.com/dedeme/ktlib/path"
	"github.com/dedeme/ktlib/sys"
	"github.com/dedeme/ktlib/str"
)

// If 'isMozilla' is 'false', it calls "wget -q --no-cache -O - 'url'" and
// returns the text read.
//
// If 'isMozilla' is 'true', it calls:
//   wget --user-agent=Mozilla
//     --load-cookies=/home/deme/.mozilla/firefox/bfrqeymk.default/cookies.sqlite
//     -q --no-cache -O - url
//
// If the reading fails, it returns an empty string.
func Wget(url string, isMozilla bool) string {
	opts := []string{
		"--user-agent=Mozilla",
		"--load-cookies=/home/deme/" +
			".mozilla/firefox/bfrqeymk.default/cookies.sqlite",
		"-q", "--no-cache", "-O", "-", url}
	if !isMozilla {
		opts = opts[2:]
	}
	stdout, stderr := sys.Cmd("wget", opts...)
	if stderr != "" || stdout == "" {
		return ""
	}
	return stdout
}

// Returns the md5 sum of a file. It calls "md5sum 'f'".
//   f: File to check.
func Md5(f string) string {
	stdout, stderr := sys.Cmd("md5sum", f)
	if stderr != "" {
		panic(stderr)
	}
	return stdout[:32]
}

// Returns the md5 sum of a string. It calls
//     bash -c "echo -n " + js.Ws('s') + " | md5sum"
// Parameter:
//   s: String to check.
func Md5Str(s string) string {
	stdout, stderr := sys.Cmd("bash", "-c", "echo -n "+js.Ws(s)+" | md5sum")
	if stderr != "" {
		panic(stderr)
	}
	return stdout[:32]
}

// Returns the sha256 sum of a file. It calls "sha256sum 'f'".
//   f: File to check.
func Sha256(f string) string {
	stdout, stderr := sys.Cmd("sha256sum", f)
	if stderr != "" {
		panic(stderr)
	}
	return stdout[:64]
}

// Returns the sha256 sum of a string. It calls
//     bash -c "echo -n " + js.Ws('s') + " | sha256sum"
// Parameter:
//   s: String to check.
func Sha256Str(s string) string {
	stdout, stderr := sys.Cmd("bash", "-c", "echo -n "+js.Ws(s)+" | sha256sum")
	if stderr != "" {
		panic(stderr)
	}
	return stdout[:64]
}

// Zip compress source in target. It calls:
//   zip -q 'target' 'source'
// If 'target' already exists, source will be added to it. If you require a
// fresh target file, you have to delete it previously.
//   source: can be a file or directory,
//   target: Zip file. If it is a relative path, it hangs on source parent.
func Zip(source, target string) {
	wd := file.Wd()

	if source == "/" || source == "" {
		panic("Bad source path '" + source + "'")
	}

	parent := path.Parent(source)
	name := path.Base(source)
	file.Cd(parent)

	opts := []string{"-r", "-q", target, name}
	if !file.IsDirectory(name) {
		opts = opts[1:]
	}

	stdout, stderr := sys.Cmd("zip", opts...)
	exists := file.Exists(target)
	file.Cd(wd)

	if stdout != "" || stderr != "" || !exists {
		msg := stdout
		if msg == "" {
			msg = stderr
		}
		if msg != "" {
			msg = "\n" + msg
		}

		panic("Fail running Zip(" + source + ", " + target + ")." + msg)
	}
}

// Unnzip uncompress source in target, It calls:
//   unzip -q 'source' -d 'target'
// Parameters:
//   source: Zip file.
//   target: A directory. It it does not exist, it is created.
func Unzip(source, target string) {
	if file.Exists(target) && !file.IsDirectory(target) {
		panic("'" + target + "' is not a directory")
	}

	stdout, stderr := sys.Cmd("unzip", "-q", source, "-d", target)
	if stdout != "" || stderr != "" {
		msg := stdout
		if msg == "" {
			msg = stderr
		}
		if msg != "" {
			msg = "\n" + msg
		}

		panic("Fail running Unzip(" + source + ", " + target + ")." + msg)
	}
}

// Returns 'Files' converted to 'pdf'. It calls
//     If 'isBook' is true:
//     htmldoc --book --charset utf-8 --footer ... --size A4 -t pdf ['Files']
//     If 'isBook' is false:
//     htmldoc --webpage --charset utf-8 --footer ... --size A4 -t pdf ['Files']
// htmldoc documentation: https://www.msweet.org/htmldoc/htmldoc.html
//   isBook: 'true' if a book shuld be generated.
//   Files : Files in HTML to convert.
func Htmldoc(isBook bool, files ...string) []byte {
  book := "--webpage"
  if isBook {
    book = "--book"
  }
  out, err := sys.Cmd(
    "htmldoc",
    append([]string{
        book,"--charset", "utf-8", "--footer", "...", "--size", "A4", "-t", "pdf",
      }, files...,
    )...,
  )
  if err != "" && !str.Starts(str.Trim(err), "PAGES:") {
    panic(err)
  }

  return []byte(out)
};

// Returns 'html' converted to 'pdf'. It calls 'Htmldoc'.
//   isBook: 'true' if a book shuld be generated.
//   html  : Text in HTML to convert.
func HtmldocStr(isBook bool, html string) []byte {
  sys.Rand()
  tmp := file.Tmp("./", "kut_htmldoc")
  file.Write(tmp, html)
  r := Htmldoc(isBook, tmp)
  file.Del(tmp)
  return r
};

