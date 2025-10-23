package main

import (
	"log/slog"

	"github.com/chrollo-lucifer-12/betteruptime/db"
)

func main() {
	db, err := db.NewGorm()
	if err != nil {
		slog.Error("error connecting to db", "err", err)
	}
	opts := ServerOpts{
		Port: ":8000",
		DB:   db.Db,
	}
	s := NewServer(opts)
	err = s.Start()
	if err != nil {
		slog.Error("error starting server", "err", err)
	}
}
