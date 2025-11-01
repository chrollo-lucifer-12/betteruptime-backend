package main

import (
	"log/slog"

	"github.com/chrollo-lucifer-12/betteruptime/db"
	"github.com/chrollo-lucifer-12/betteruptime/server"
)

func main() {
	db, err := db.NewGorm()
	if err != nil {
		slog.Error("error connecting to db", "err", err)
	}
	opts := server.ServerOpts{
		Port: ":8000",
		DB:   db.Db,
	}
	s := server.NewServer(opts)
	err = s.Start()
	if err != nil {
		slog.Error("error starting server", "err", err)
	}
}
