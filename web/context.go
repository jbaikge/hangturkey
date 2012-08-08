package web

import (
	"code.google.com/p/gorilla/sessions"
)

var (
	store = sessions.NewCookieStore([]byte("MMMMMMM, TURKEYS"))
)

type Context struct {
	Session sessions.Session
	State   GameState
}

func Close() {
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

func stateFromSession(session *sessions.Session) (state GameState) {
	if s, ok := session.Values["state"]; ok {
		state = s.(GameState)
	} else {
		state = GameState{}
	}
	return
}
