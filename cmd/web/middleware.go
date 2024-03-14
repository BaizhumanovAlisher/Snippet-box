package main

import (
	"net/http"
	"strings"
)

// this function used for cancel listing in url: /static/
func (app *application) neuter(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if strings.HasSuffix(r.URL.Path, "/") {
			app.notFound(w)
			return
		}

		next.ServeHTTP(w, r)
	})
}
