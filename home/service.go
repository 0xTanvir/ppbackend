package home

import "github.com/0xTanvir/pp/db"

type Service struct {
	DB *db.DB
}

// GetUI
func (r *Service) GetUI() string {
	return "Here it will render ui for users."
}
