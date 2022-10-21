// Copyright 31-May-2022 ºDeme
// GNU General Public License - V3 <http://www.gnu.org/licenses/>

// TCP connections.
//
// Example:
//    package main
//
//    import (
//      "github.com/dedeme/ktlib/str"
//      "github.com/dedeme/ktlib/sys"
//      "github.com/dedeme/ktlib/tcp"
//      "github.com/dedeme/ktlib/thread"
//    )
//
//    // Process to run in server
//    func process(conn tcp.ConnT) bool {
//      tx, err := tcp.Read(conn, 10000)
//      if err != nil {
//        panic(err)
//      }
//      if tx == "end" {
//        err := tcp.Write(conn, "Closing server")
//        if err != nil {
//          panic(err)
//        }
//        tcp.CloseConnection(conn)
//        return true
//      }
//      err = tcp.Write(conn, "Send from server: "+tx)
//      if err != nil {
//        panic(err)
//      }
//      tcp.CloseConnection(conn)
//      return false
//    }
//
//    func main() {
//
//      // Launch server.
//      th1 := thread.Start(func() {
//        sv := tcp.Server(23344)
//
//        for {
//          conn, err := tcp.Accept(sv, 0)
//          if err != nil {
//            panic(err)
//          }
//          if process(conn) {
//            break
//          }
//        }
//
//        tcp.CloseServer(sv)
//      })
//
//      // Three connection from client.
//      for i := 0; i < 3; i++ {
//        conn, err := tcp.Dial("localhost:23344", 0)
//        if err != nil {
//          panic(err)
//        }
//        sys.Println(str.Fmt("Sending 'abc%d'", i))
//        err = tcp.Write(conn, str.Fmt("abc%d", i))
//        if err != nil {
//          panic(err)
//        }
//        rq, err0 := tcp.Read(conn, 10000)
//        if err0 != nil {
//          panic(err0)
//        }
//        sys.Println(rq)
//        tcp.CloseConnection(conn)
//      }
//
//      // Ending server
//      conn, err := tcp.Dial("localhost:23344", 0)
//      if err != nil {
//        panic(err)
//      }
//      err = tcp.Write(conn, "end")
//      if err != nil {
//        panic(err)
//      }
//      rq, err0 := tcp.Read(conn, 10000)
//      if err0 != nil {
//        panic(err0)
//      }
//      sys.Println(rq)
//      tcp.CloseConnection(conn)
//
//      // Wait until server is ended.
//      thread.Join(th1)
//    }
package tcp

import (
	"errors"
	"fmt"
	"io"
	"net"
	"strconv"
	"time"
)

// Server type
type ServerT = net.Listener

// Connection type
type ConnT = net.Conn

// Waits for a new connection on server 'sv'.
//
// The connection will have a maximum 'tm' milliseconds to finish any
// I/O operation. If 'tm' <= 0 there will not have limit for waiting.
//
// 'tcp.Accept' returns:
//    - A new tcp connection to use with tcp.Read, tcp.ReadBin,
//      'tcp.Write', 'tcp.WriteBin' or 'tcp.CloseServer'.
//    - 'err != nil' if the connection failed.
// The tcp connection returned shuld be closed with tcp.CloseConnection.
func Accept(sv ServerT, tm int) (conn ConnT, err error) {
	conn, err = sv.Accept()
	if err == nil {
		if tm > 0 {
			err = conn.SetDeadline(time.Now().Add(time.Duration(tm) * time.Millisecond))
		}
	}
	return
}

// Closes 'conn'.
func CloseConnection(conn ConnT) {
	err := conn.Close()
	if err != nil {
		panic(err)
	}
}

// Closes 'sv'.
func CloseServer(sv ServerT) {
	err := sv.Close()
	if err != nil {
		panic(err)
	}
}

// Client connection.
//    server: Indicated as 'server:port'. Examples:
//            'localhost:21786', '127.0.0.1:22454".
//    tm    : Maximun time in milleseconds to wait for a response.
//            If 'tm' <= 0 there will not have limit for waiting.
// tcp.Dial returns:
//    - A new tcp connection to use with tcp.Read, tcp.ReadBin, tcp.Write,
//      tcp.WriteBin or tcp.CloseConnection.
//    - 'err != nil' if the connection failed.
// The tcp connection returned shuld be closed with tcp.CloseConnection.
func Dial(tcpServer string, tm int) (conn ConnT, err error) {
	conn, err = net.Dial("tcp", tcpServer)
	if err == nil {
		if tm > 0 {
			err = conn.SetDeadline(time.Now().Add(time.Duration(tm) * time.Millisecond))
		}
	}
	return
}

// Reads a request from connection 'conn' with a maximum bytes length of 'lim'
// ('lim' must be greater than '0').
func Read(conn ConnT, lim int) (data string, err error) {
	var bs []byte
	bs, err = ReadBin(conn, lim)
	if err == nil {
		data = string(bs)
	}
	return
}

// Reads a request from connection 'conn' with a maximum bytes length of 'lim'
// ('lim' must be greater than '0').
func ReadBin(conn ConnT, lim int) (data []byte, err error) {
	if lim < 1 {
		panic("Connection limit less than 1")
	}
	bs := make([]byte, lim+1)
	n, err0 := conn.Read(bs)
	if err0 != nil {
		if err0 == io.EOF {
			data = []byte{}
			return
		}
		err = err0
		return
	}
	if n > lim {
		err = errors.New(fmt.Sprintf("Bytes read out of limit (%v)", lim))
		return
	}
	bs2 := make([]byte, n)
	for i := 0; i < n; i++ {
		bs2[i] = bs[i]
	}
	data = bs2
	return
}

// Returns a tcp server to use with 'tcp.accept'.
//    port: Comunications port.
// The tcp server returned shuld be closed with tcp.CloseServer.
func Server(port int) ServerT {
	sv, err := net.Listen("tcp", ":"+strconv.Itoa(port))
	if err != nil {
		panic(err)
	}
	return sv
}

// Writes a string through connection 'conn'.
func Write(conn ConnT, s string) (err error) {
	_, err = fmt.Fprintf(conn, s)
	return
}

// Writes a []byte through connection 'conn'.
func WriteBin(conn ConnT, bs []byte) (err error) {
	_, err = conn.Write(bs)
	return
}
