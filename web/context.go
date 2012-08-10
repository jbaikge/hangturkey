package web

import (
	"log"
	"net/http"

	"code.google.com/p/gorilla/sessions"
	"github.com/jbaikge/hangturkey/app"
)

const (
	correctScore   = 5
	incorrectScore = -1
)

var (
	store = sessions.NewCookieStore([]byte("MMMMMMM, TURKEYS"))
)

type Context struct {
	Session *sessions.Session
	State   app.GameState
	request *http.Request
	writer  http.ResponseWriter
}

func NewContext(w http.ResponseWriter, req *http.Request) (*Context, error) {
	// Ignore session from blank error
	session, _ := store.Get(req, "state")
	state := stateFromSession(session)
	log.Printf("STATE: %+v", state)
	return &Context{
		Session: session,
		State:   state,
		request: req,
		writer:  w,
	}, nil
}

func (c *Context) Close() {
	// noop
}

func (c *Context) SaveSession() {
	c.Session.Values["state"] = c.State
	c.Session.Save(c.request, c.writer)
}

func (c Context) TotalScore() int {
	return c.State.TotalScore(correctScore, incorrectScore)
}

func (c Context) WordScore() int {
	return c.State.CurrentWordScore(correctScore, incorrectScore)
}

func stateFromSession(session *sessions.Session) (state app.GameState) {
	if s, ok := session.Values["state"]; ok {
		state = s.(app.GameState)
	} else {
		state = app.GameState{}
	}
	return
}
