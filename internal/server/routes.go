package server

import (
	"goth-complete-setup/internal"
	"goth-complete-setup/internal/api_errors"
	"log"
	"net/http"

	"github.com/a-h/templ"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func (api *Api) RegisterRoutes() *chi.Mux {
	r := chi.NewRouter()

	r.Use(middleware.Heartbeat("/ping"))
	r.Use(middleware.Logger)

	r.Handle("/static/*", http.FileServer(http.FS(internal.StaticFiles)))

	r.HandleFunc("/", NewHandler(root))

	return r
}

type CustomHandler func(w http.ResponseWriter, request *http.Request) error

func NewHandler(customHandler CustomHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		err := customHandler(w, r)
		if err != nil {
			log.Printf("Error: %s", err.Error())
			if clientErr, ok := err.(*api_errors.ClientErr); ok {
				respondWithJSON(w, clientErr.HttpCode, clientErr)
			} else {
				respondWithJSON(w, http.StatusInternalServerError,
					api_errors.InternalErr{
						HttpCode: http.StatusInternalServerError,
						Message:  "internal server error",
					},
				)
			}
		}
	}
}

func Render(w http.ResponseWriter, r *http.Request, comp templ.Component) error {
	return comp.Render(r.Context(), w)
}
