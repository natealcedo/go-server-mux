package main

import (
	"github.com/natealcedo/go-server-mux/handlers"
	"github.com/natealcedo/go-server-mux/middleware"
	"log/slog"
	"net/http"
)

const PORT = ":3000"

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /",
		middleware.LoggingMiddleware(
			handlers.MakeHandler(handlers.HandleGet),
		),
	)

	mux.HandleFunc("POST /",
		middleware.LoggingMiddleware(
			handlers.MakeHandler(handlers.HandlePost),
		),
	)

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
