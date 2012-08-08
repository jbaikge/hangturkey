package main

import (
	"code.google.com/p/gorilla/sessions"
)

type Context struct {
	Session sessions.Session
}
