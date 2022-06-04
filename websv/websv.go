// Copyright 02-Jun-2022 ºDeme
// GNU General Public License - V3 <http://www.gnu.org/licenses/>

// Web server.
package websv

import (
	"github.com/dedeme/ktlib/file"
	"github.com/dedeme/ktlib/path"
	"github.com/dedeme/ktlib/str"
	"github.com/dedeme/ktlib/sys"
	"github.com/dedeme/ktlib/tcp"
	"github.com/dedeme/ktlib/thread"
)

var mimetypes = map[string]string{
	".aac":   "audio/aac",
	".abw":   "application/x-abiword",
	".avi":   "video/x-msvideo",
	".bin":   "application/octet-stream",
	".bmp":   "image/bmp",
	".bz":    "application/x-bzip",
	".bz2":   "application/x-bzip2",
	".cda":   "application/x-cdf",
	".css":   "text/css",
	".csv":   "text/csv",
	".doc":   "application/msword",
	".epub":  "application/epub+zip",
	".gz":    "application/gzip",
	".gif":   "image/gif",
	".htm":   "text/html",
	".html":  "text/html",
	".ico":   "image/vnd.microsoft.icon",
	".jar":   "application/java-archive",
	".jpeg":  "image/jpeg",
	".js":    "text/javascript",
	".json":  "application/json",
	".mid":   "audio/x-midi",
	".midi":  "audio/x-midi",
	".mjs":   "text/javascript",
	".mp3":   "audio/mpeg",
	".mp4":   "video/mp4",
	".mpeg":  "video/mpeg",
	".odp":   "application/vnd.oasis.opendocument.presentation",
	".ods":   "application/vnd.oasis.opendocument.spreadsheet",
	".odt":   "application/vnd.oasis.opendocument.text",
	".oga":   "audio/ogg",
	".ogv":   "video/ogg",
	".ogx":   "application/ogg",
	".otf":   "font/otf",
	".png":   "image/png",
	".pdf":   "application/pdf",
	".php":   "application/x-httpd-php",
	".ppt":   "application/vnd.ms-powerpoint",
	".rar":   "application/vnd.rar",
	".rtf":   "application/rtf",
	".svg":   "image/svg+xml",
	".tar":   "application/x-tar",
	".tif":   "image/tiff",
	".tiff":  "image/tiff",
	".ttf":   "font/ttf",
	".txt":   "text/plain",
	".wav":   "audio/wav",
	".xhtml": "application/xhtml+xml",
	".xls":   "application/vnd.ms-excel",
	".xml":   "application/xml",
	".xul":   "application/vnd.mozilla.xul+xml",
	".zip":   "application/zip",
	".7z":    "application/x-7z-compressed",
}

// Tcp server.
var sv tcp.ServerT

// Tcp stopped?
var stopCode string

// Web pages root in the file system.
var wroot string

// Starts server.
//   port    : Connection port.
//   tm      : Maximum 'tm' milliseconds to finish any I/O operation.
//             If 'tm' <= 0 there will not have limit for waiting.
//   root    : Server working directory.
//               -Web pages must be in 'root'/www.
//   stopcode: Code for stopping server.
//   handler : Function to process requests.
// Connection request/response is limited to 10.000.000 bits.
func Start(
	port, tm int, root, stopcode string, handler func(string) string,
) {
	if sv != nil {
		return
	}

	stopCode = stopcode
	sv = tcp.Server(port)
	wroot = path.Cat(root, "www")
	file.Cd(root)

	for {
		conn, err := tcp.Accept(sv, tm)
		if err != nil {
			panic(err)
		}

		tx := tcp.Read(conn, 10_000_000)
		if tx == stopCode {
			tcp.CloseConnection(conn)
			tcp.CloseServer(sv)
			break
		}

		thread.Run(func() {
			tcp.WriteBin(conn, []byte(handler(tx)))
			tcp.CloseConnection(conn)
		})
	}
}

// Stops server.
func Stop(port int) {
	conn, err := tcp.Dial("localhost:"+str.Fmt("%d", port), 0)
	if err == nil {
		tcp.Write(conn, stopCode)
	}
}

// Returns web file data or 'ok=false' if 'rq' is not a GET request.
//
// Web pages must be in 'root'/www.
//
// NOTE: Only HTTP/1.1 requests are allowed.
//   rq: Complete client request.
func GetRq(rq string) (rp string, ok bool) {
	ix := str.Index(rq, "\n")
	if ix == -1 {
		rp = BadRequestRp(rq)
		ok = true
		return
	}

	rq = str.Trim(rq[0:ix])
	if !str.Starts(rq, "GET ") || !str.Ends(rq, " HTTP/1.1") {
		return
	}
	rq = str.Trim(rq[4:])
	ix = str.Index(rq, " ")
	if ix == -1 {
		rp = BadRequestRp(rq)
		ok = true
		return
	}

	rq = rq[:ix]
	ix = str.Index(rq, "?")
	r2 := rq
	if ix != -1 {
		r2 = rq[:ix]
	}

	ext := path.Extension(r2)
	if ext == "" {
		lc := path.Cat(r2, "index.html")
		if ix != -1 {
			lc += rq[ix:]
		}
		rp = "HTTP/1.1 308 Permanent Redirect\n" +
			"Location: " + lc + "\n\n"
		ok = true
		return
	}

	fpath := path.Canonical(path.Cat(wroot, r2))
	if !file.Exists(fpath) {
		rp = NotFoundRp()
		ok = true
		return
	}

	tp, ok := mimetypes[ext]
	if !ok {
		tp = "text/plain"
	}
	body := file.Read(fpath)

	rp = "HTTP/1.1 200 OK\n" +
		"Server: Kut Webserver\n" +
		"Content-type: " + tp + "\n" +
		"Content-length: " + str.Fmt("%d", len(body)) + "\n\n" +
		body
	ok = true
	return
}

// Executes a 'ccgi.sh' command and returns its result or 'ok=false' if
// rq is not a valid POST request ('POST /cgi-bin/ccgi.sh HTTP/1.1').
//
// Binary files must be in 'root'/bin.
func DmCgiRq(rq string) (rp string, ok bool) {
	rqUnix := str.Replace(rq, "\r", "")
	ix := str.Index(rq, "\n")
	if ix == -1 {
		rp = BadRequestRp(rq)
		ok = true
		return
	}

	rq = str.Trim(rq[:ix])
	if rq != "POST /cgi-bin/ccgi.sh HTTP/1.1" {
		return
	}

	ix = str.Index(rqUnix, "\n\n")
	if ix == -1 {
		rp = BadRequestRp(rq)
		ok = true
		return
	}

	rqUnix = str.Trim(rqUnix[ix+2:])
	ix = str.Index(rqUnix, ":")
	if ix == -1 {
		rp = BadRequestRp(rq)
		ok = true
		return
	}

	cmd := rqUnix[:ix]
	par := rqUnix[ix+1:]
	stdout, stderr := sys.Cmd(path.Cat("bin", cmd), par)

	if stderr != "" {
		rp = stderr
		if stdout != "" {
			rp = stdout
		}
	} else {
		rp = TextRp(stdout)
	}
	ok = true
	return
}

// Return a html valid response.
//   h: HTML text.
func HtmlRp(h string) string {
	return "HTTP/1.1 200 OK\n" +
		"Server: Kut Webserver\n" +
		"Content-type: text/html\n" +
		"Content-length: " + str.Fmt("%d", len(h)) + "\n\n" +
		h
}

// Return a text valid response.
//   h: HTML text.
func TextRp(tx string) string {
	return "HTTP/1.1 200 OK\n" +
		"Server: Kut Webserver\n" +
		"Content-type: text/text\n" +
		"Content-length: " + str.Fmt("%d", len(tx)) + "\n\n" +
		tx
}

// Return a 'Not Found' response.
func NotFoundRp() string {
	return "HTTP/1.1 404 Not Found\n" +
		"Server: Kut Webserver\n" +
		"Content-type: text/plain\n\n" +
		"Page not found"
}

// Return a 'Bad Request' response.
//   rq: Request.
func BadRequestRp(rq string) string {
	return "HTTP/1.1 400 Bad Request\n" +
		"Server: Kut Webserver\n" +
		"Content-type: text/plain\n\n" +
		"Bad Request:\n" + rq
}
