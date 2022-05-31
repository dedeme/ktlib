// Copyright 31-May-2020 ºDeme
// GNU General Public License - V3 <http://www.gnu.org/licenses/>

// TCP connections.
package tcp

import (
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
func Read(conn ConnT, lim int) string {
	return string(ReadBin(conn, lim))
}

// Reads a request from connection 'conn' with a maximum bytes length of 'lim'
// ('lim' must be greater than '0').
func ReadBin(conn ConnT, lim int) []byte {
	if lim < 1 {
		panic("Connection limit less than 1")
	}
	bs := make([]byte, lim+1)
	n, err := conn.Read(bs)
	if err != nil {
		if err == io.EOF {
			return []byte{}
		}
		panic(err)
	}
	if n > lim {
		panic(fmt.Sprintf("Bytes read out of limit (%v)", lim))
	}
	bs2 := make([]byte, n)
	for i := 0; i < n; i++ {
		bs2[i] = bs[i]
	}
	return bs2
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
func Write(conn ConnT, s string) {
	fmt.Fprintf(conn, s)
}

// Writes a []byte through connection 'conn'.
func WriteBin(conn ConnT, bs []byte) {
	_, err := conn.Write(bs)
	if err != nil {
		panic(err)
	}
}
