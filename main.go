package main

import (
	"github.com/natealcedo/go-server-mux/handlers"
	"log/slog"
	"net/http"
)

const PORT = ":3000"

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /", handlers.MakeHandler(handlers.HandleGet))

	server := &http.Server{
		Addr:    PORT,
		Handler: mux,
	}

	slog.Info("Server started on", "port", PORT)

	if err := server.ListenAndServe(); err != nil {
		slog.Error("Server failed to start: ", err)
		panic(err)
	}
}
