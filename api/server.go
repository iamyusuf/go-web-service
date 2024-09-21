package api

import "gorm.io/gorm"

type Server struct {
	Db *gorm.DB
}

func NewServer(db *gorm.DB) *Server {
	return &Server{
		Db: db,
	}
}
