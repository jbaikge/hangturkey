package main

import (
	"code.google.com/p/gorilla/sessions"
)

var (
	store = sessions.NewCookieStore([]byte("MMMMMMM, TURKEYS"))
)

type Context struct {
	Session sessions.Session
}
