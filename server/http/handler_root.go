package http

import (
	"fmt"
	"net/http"
)

func root(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Invalid link\n")
}
