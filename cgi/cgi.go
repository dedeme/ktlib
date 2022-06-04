// Copyright 01-Jun-2022 ºDeme
// GNU General Public License - V3 <http://www.gnu.org/licenses/>

// Utilities for HTML conections between client - server.
package cgi

import (
	"fmt"
	"github.com/dedeme/ktlib/cryp"
	"github.com/dedeme/ktlib/file"
	"github.com/dedeme/ktlib/js"
	"math/rand"
	"path"
	"time"
)

const (
	// Standad length of passwords
	Klen          = 300
	tNoExpiration = 2592000 // seconds == 30 days
	demeKey       = "nkXliX8lg2kTuQSS/OoLXCk8eS4Fwmc+N7l6TTNgzM1vdKewO0cjok51vcdl" +
		"OKVXyPu83xYhX6mDeDyzapxL3dIZuzwyemVw+uCNCZ01WDw82oninzp88Hef" +
		"bn3pPnSMqEaP2bOdX+8yEe6sGkc3IO3e38+CqSOyDBxHCqfrZT2Sqn6SHWhR" +
		"KqpJp4K96QqtVjmXwhVcST9l+u1XUPL6K9HQfEEGMGcToMGUrzNQxCzlg2g+" +
		"Hg55i7iiKbA0ogENhEIFjMG+wmFDNzgjvDnNYOaPTQ7l4C8aaPsEfl3sugiw"
	noSessionKey = "nosession"
)

type T = map[string]string

var homeV string
var tExpirationV int
var fkey = cryp.Key(demeKey, len(demeKey)) // File encryption key

// Initializes a new interface of commnications.
//    home       : Aboslute path of application directory. For example:
//                    "/peter/wwwcgi/dmcgi/JsMon"
//                    or
//                    "/home/deme/.dmCApp/JsMon".
//    tExpiration: Time in seconds.
func Initialize(home string, tExpiration int) {
	rand.Seed(time.Now().UTC().UnixNano())
	homeV = home
	tExpirationV = tExpiration

	if !file.Exists(home) {
		file.Mkdir(home)
	}

	fusers := path.Join(home, "users.db")
	if !file.Exists(fusers) {
		writeUsers([]*userT{})
		putUser("admin", demeKey, "0")
		writeSessions([]*sessionT{})
	}
}

// User ------------------------------------------------------------------------

type userT struct {
	id    string
	pass  string
	level string
}

func uToJson(us []*userT) string {
	var tmps []string
	for _, u := range us {
		jss := []string{js.Ws(u.id), js.Ws(u.pass), js.Ws(u.level)}
		tmps = append(tmps, js.Wa(jss))
	}
	return js.Wa(tmps)
}
func uFromJson(j string) (us []*userT) {
	jss := js.Ra(j)
	for _, ujs := range jss {
		var ujss []string
		ujss = js.Ra(ujs)
		var id, pass, level string
		id = js.Rs(ujss[0])
		pass = js.Rs(ujss[1])
		level = js.Rs(ujss[2])
		us = append(us, &userT{id, pass, level})
	}
	return
}

func writeUsers(us []*userT) {
	j := uToJson(us)
	file.Write(path.Join(homeV, "users.db"), cryp.Encode(fkey, j))
}

func readUsers() []*userT {
	j := cryp.Decode(fkey, file.Read(path.Join(homeV, "users.db")))
	return uFromJson(j)
}

func putUser(id, pass, level string) {
	pass = cryp.Key(pass, Klen)
	users := readUsers()

	var r *userT
	for _, u := range users {
		if u.id == id {
			r = u
			break
		}
	}
	if r == nil {
		users = append(users, &userT{id, pass, level})
	} else {
		r.pass = pass
		r.level = level
	}
	writeUsers(users)
}

// If check fails, returns "". Otherwise it returns user level.
func checkUser(id, pass string) string {
	pass = cryp.Key(pass, Klen)
	users := readUsers()
	for _, u := range users {
		if u.id == id && u.pass == pass {
			return u.level
		}
	}
	return ""
}

// Session ---------------------------------------------------------------------

type sessionT struct {
	id     string
	comKey string // Communication key
	conKey string // Connection key
	user   string // User id
	level  string
	time   int64 // time.Unix
	lapse  int
}

func sToJson(ss []*sessionT) string {
	var tmps []string
	for _, s := range ss {
		jss := []string{
			js.Ws(s.id), js.Ws(s.comKey), js.Ws(s.conKey),
			js.Ws(s.user), js.Ws(s.level),
			js.Wl(s.time), js.Wi(s.lapse),
		}
		tmps = append(tmps, js.Wa(jss))
	}
	return js.Wa(tmps)
}
func sFromJson(j string) (ss []*sessionT, err error) {
	jss := js.Ra(j)
	for _, sjs := range jss {
		var sjss []string
		sjss = js.Ra(sjs)
		var id, comKey, conKey, user, level string
		var time int64
		var lapse int
		id = js.Rs(sjss[0])
		comKey = js.Rs(sjss[1])
		conKey = js.Rs(sjss[2])
		user = js.Rs(sjss[3])
		level = js.Rs(sjss[4])
		time = js.Rl(sjss[5])
		lapse = js.Ri(sjss[6])
		ss = append(ss, &sessionT{id, comKey, conKey, user, level, time, lapse})
	}
	return
}

// Returns 'false' if 's' is out of date.
func (s *sessionT) update() bool {
	now := time.Now().Unix()
	if now > s.time+int64(s.lapse) {
		return false
	}
	s.time = now + int64(s.lapse)
	return true
}

func writeSessions(ss []*sessionT) {
	j := sToJson(ss)
	file.Write(path.Join(homeV, "sessions.db"), cryp.Encode(fkey, j))
}

func readSessions() []*sessionT {
	j := cryp.Decode(fkey, file.Read(path.Join(homeV, "sessions.db")))
	r, err := sFromJson(j)
	if err != nil {
		panic(err)
	}
	return r
}

// Adds session and purge sessions.
func addSession(sessionId, comKey, conKey, user, level string, lapse int) {
	now := time.Now().Unix()
	ss := readSessions()
	var newSs []*sessionT
	for _, s := range ss {
		if now <= s.time+int64(s.lapse) {
			newSs = append(newSs, s)
		}
	}
	newSs = append(newSs,
		&sessionT{sessionId, comKey, conKey, user, level, now, lapse})
	writeSessions(newSs)
}

// Replace a session with a new date and a new connection key
func replaceSession(mdss *sessionT) {
	ss := readSessions()
	var newSs []*sessionT
	for _, s := range ss {
		if s.id != mdss.id {
			newSs = append(newSs, s)
		}
	}
	newSs = append(newSs, mdss)
	writeSessions(newSs)
}

// Public interface ------------------------------------------------------------

// Root application directory.
func Home() string {
	return homeV
}

// Sends to client 'communicationKey', 'userId' and 'userLevel'. If conection
// fails every one is "".
//    sessionId: Session identifier.
//    return   : {key: String, conKey:String, user: String, level: String}.
func Connect(sessionId string) string {
	var r *sessionT
	for _, s := range readSessions() {
		if s.id == sessionId && s.update() {
			r = s
			break
		}
	}

	comKey := ""
	conKey := ""
	user := ""
	level := ""
	if r != nil {
		comKey = r.comKey
		user = r.user
		level = r.level
		conKey = cryp.GenK(Klen)
		r.conKey = conKey
		replaceSession(r)
	}
	return Rp(sessionId, T{
		"key":    js.Ws(comKey),
		"conKey": js.Ws(conKey),
		"user":   js.Ws(user),
		"level":  js.Ws(level),
	})
}

// Sends to client 'sessionId', 'communicationKey' and 'userLevel'. If
// conection fails every one is "".
//    key           : Communication key
//    user          : User id.
//    pass          : User password.
//    withExpiration: If is set to false, session will expire after 30 days.
//    return        : {sessionId: String, key: String, conKey: String,
//                     level: String}.
func Authentication(key, user, pass string, withExpiration bool) string {
	sessionId := ""
	comKey := ""
	conKey := ""
	level := checkUser(user, pass)
	if level != "" {
		sessionId = cryp.GenK(Klen)
		comKey = cryp.GenK(Klen)
		conKey = cryp.GenK(Klen)

		lapse := tNoExpiration
		if withExpiration {
			lapse = tExpirationV
		}
		addSession(sessionId, comKey, conKey, user, level, lapse)
	}

	return Rp(key, T{
		"sessionId": js.Ws(sessionId),
		"key":       js.Ws(comKey),
		"conKey":    js.Ws(conKey),
		"level":     js.Ws(level),
	})
}

// Returns the session communication key.
//		ssId  : Session identifier.
//		conKey: Connection key. If its value is "", this parameter is not used.
func GetComKey(ssId, conKey string) (comKey string, ok bool) {
	ss := readSessions()
	for _, s := range ss {
		if s.id == ssId && (conKey == "" || conKey == s.conKey) && s.update() {
			comKey = s.comKey
			ok = true
			return
		}
	}
	return
}

// Changes user password.
//    ck    : Communication key
//    user  : User name to change password.
//    old   : Old password.
//    new   : New password.
//    return: After call 'Rp()', boolean field {ok:true|false}, sets to true
//            if operation succeeded. A fail can come up if 'user'
//            authentication fails.
func ChangePass(ck, user, old, new string) (rp string) {
	rp = Rp(ck, T{"ok": js.Wb(false)})

	us := readUsers()
	var u *userT
	for _, u0 := range us {
		if u0.id == user {
			u = u0
			break
		}
	}
	if u == nil {
		return
	}

	old2 := cryp.Key(old, Klen)
	if old2 != u.pass {
		return
	}

	u.pass = cryp.Key(new, Klen)
	writeUsers(us)
	rp = Rp(ck, T{"ok": js.Wb(true)})
	return
}

// Deletes 'sessionId' and returns an empty response.
func DelSession(ck string, sessionId string) string {
	ss := readSessions()
	var newss []*sessionT
	for _, s := range ss {
		if s.id != sessionId {
			newss = append(newss, s)
		}
	}
	writeSessions(newss)
	return RpEmpty(ck)
}

// Messages --------------------------------------------------------------------

// Returns a response to send to client.
//	 ck: Communication key.
//	 rp: Response.
func Rp(ck string, rp T) string {
	j := js.Wo(rp)
	return cryp.Encode(ck, j)
}

// Returns an empty response.
//	 ck: Communication key.
func RpEmpty(ck string) string {
	return Rp(ck, T{})
}

// Returns a message with an only field "error" with value 'msg'.
//	 ck: Communication key.
func RpError(ck, msg string) string {
	return Rp(ck, T{"error": js.Ws(msg)})
}

// Returns a message with an only field "expired" with value 'true',
// codified with the key 'noSessionKey' ("nosession")
func RpExpired() string {
	return Rp(noSessionKey, T{"expired": js.Wb(true)})
}

// Requests --------------------------------------------------------------------

// Reads a bool value
func RqBool(rq T, key string) (v bool) {
	j, ok := rq[key]
	if !ok {
		panic(fmt.Sprintf("Key '%v' not found in request", key))
	}
	v = js.Rb(j)
	return
}

// Reads a int value
func RqInt(rq T, key string) (v int) {
	j, ok := rq[key]
	if !ok {
		panic(fmt.Sprintf("Key '%v' not found in request", key))
	}
	v = js.Ri(j)
	return
}

// Reads a int64 value
func RqLong(rq T, key string) (v int64) {
	j, ok := rq[key]
	if !ok {
		panic(fmt.Sprintf("Key '%v' not found in request", key))
	}
	v = js.Rl(j)
	return
}

// Reads a float32 value
func RqFloat(rq T, key string) (v float32) {
	j, ok := rq[key]
	if !ok {
		panic(fmt.Sprintf("Key '%v' not found in request", key))
	}
	v = js.Rf(j)
	return
}

// Reads a float64 value
func RqDouble(rq T, key string) (v float64) {
	j, ok := rq[key]
	if !ok {
		panic(fmt.Sprintf("Key '%v' not found in request", key))
	}
	v = js.Rd(j)
	return
}

// Reads a string value
func RqString(rq T, key string) (v string) {
	j, ok := rq[key]
	if !ok {
		panic(fmt.Sprintf("Key '%v' not found in request", key))
	}
	v = js.Rs(j)
	return
}
