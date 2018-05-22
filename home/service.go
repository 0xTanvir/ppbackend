package home

import (
	"github.com/0xTanvir/pp/auth"
	"github.com/0xTanvir/pp/db"
)

// Service all logic functionality of User
type Service struct {
	DB   *db.DB
	Auth *auth.Service
}

// GetUI handle
func (r *Service) GetUI() string {
	return "Yes, Programmer's Playground back-end is running"
}
