package server

import (
	"github.com/0xTanvir/pp/contest"
	"github.com/0xTanvir/pp/home"
	"github.com/0xTanvir/pp/users"
)

// Controllers are all the controllers
// which is used in server
type Controllers struct {
	User    *users.Controller
	Home    *home.Controller
	Contest *contest.Controller
}
