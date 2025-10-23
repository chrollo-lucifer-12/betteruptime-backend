package main

import "log/slog"

func main() {
	opts := ServerOpts{
		Port: ":8000",
	}
	s := NewServer(opts)
	err := s.Start()
	if err != nil {
		slog.Error("error starting server", "err", err)
	}
}
