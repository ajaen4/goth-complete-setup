package server

import (
	"goth-complete-setup/internal/components"
	"net/http"
)

func root(w http.ResponseWriter, r *http.Request) error {
	return Render(w, r, components.Home())
}
