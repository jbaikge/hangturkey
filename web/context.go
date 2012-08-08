package web

import (
	"net/http"

	"code.google.com/p/gorilla/sessions"
	"github.com/jbaikge/hangturkey/app"
)

var (
	store = sessions.NewCookieStore([]byte("MMMMMMM, TURKEYS"))
)

type Context struct {
	Session *sessions.Session
	State   app.GameState
}

func (c *Context) Close() {
	// noop
}

func NewContext(req *http.Request) (*Context, error) {
	// Ignore session from blank error
	session, _ := store.Get(req, "state")
	state := stateFromSession(session)
	return &Context{
		Session: session,
		State:   state,
	}, nil
}

func stateFromSession(session *sessions.Session) (state app.GameState) {
	if s, ok := session.Values["state"]; ok {
		state = s.(app.GameState)
	} else {
		state = app.GameState{}
	}
	return
}
