package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/mux"
)

func getRoutes() http.Handler {
	route := mux.NewRouter()

	//Write the route here..
	route.HandleFunc("/health", func(c *gin.Context) {
		c.IntendedJSON(http.StatusOK, "OK")
	}).Methods("GET", "OPTIONS")

	// Placeholder for generated code. Do not remove or modify this comment.

	route.Use(checkMethodOptions(route))

	return route
}
