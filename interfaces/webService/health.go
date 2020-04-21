package webService

import (
	"fmt"
	"net/http"
)

func (ws *webService) Health(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	fmt.Fprint(w, `{"status": "ok"}`)
}
