package handlers

import "net/http"

type Handler func(w http.ResponseWriter, r *http.Request) error

func MakeHandler(handler Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := handler(w, r); err != nil {
			http.Error(w, err.Error(), 500)
		}
	}
}
