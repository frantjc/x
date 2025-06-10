package xhttp

import (
	"net/http"
	"slices"
	"strings"
)

type allowHandler struct {
	handler      http.Handler
	allowMethods []string
}

func (ah *allowHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if slices.Contains([]string{http.MethodHead, http.MethodOptions}, r.Method) {
		allow := strings.Join(ah.allowMethods, ", ")
		w.Header().Set("Allow", allow)

		if http.MethodOptions == r.Method {
			w.Header().Set("Access-Control-Allow-Methods", allow)
			w.WriteHeader(http.StatusOK)
			return
		}
	}

	if !slices.Contains(ah.allowMethods, r.Method) {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	ah.handler.ServeHTTP(w, r)
}

func AllowHandler(h http.Handler, allowMethods []string) http.Handler {
	return &allowHandler{handler: h, allowMethods: allowMethods}
}
