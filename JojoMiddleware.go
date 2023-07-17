package main

import (
	"context"
	"go-generate/entity"
	"net/http"

	"github.com/gorilla/mux"
)

type contextKey int

const (
	jwtKey         contextKey = 0
	userProfileKey contextKey = 1
)

func HandleJojoMiddleware(h http.Handler, middlewares ...mux.MiddlewareFunc) http.Handler {
	for _, m := range middlewares {
		h = m(h)
	}
	return h
}

func addItemToContext(r *http.Request, key contextKey, value any) *http.Request {
	ctxWithItem := context.WithValue(r.Context(), key, value)
	return r.WithContext(ctxWithItem)
}

func GetJWT(r *http.Request) entity.JojoJWT {
	return r.Context().Value(jwtKey).(entity.JojoJWT)
}
