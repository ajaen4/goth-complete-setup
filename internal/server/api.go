package server

import (
	"fmt"
	"net/http"
	"time"
)

type Api struct{}

func NewServer(port int) *http.Server {

	api := &Api{}

	server := &http.Server{
		Addr:         fmt.Sprintf(":%d", port),
		Handler:      api.RegisterRoutes(),
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	return server
}
