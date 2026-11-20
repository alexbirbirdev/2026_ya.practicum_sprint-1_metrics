package server

import (
	"context"
	"database/sql"
	"time"
)

type ServerUsecase struct {
	db *sql.DB
}

func NewServerUsecase(db *sql.DB) *ServerUsecase {
	return &ServerUsecase{db: db}
}

func (u *ServerUsecase) Ping() error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	return u.db.PingContext(ctx)
}
