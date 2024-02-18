package xhttp

import (
	"net/http"
	"strings"

	xslice "github.com/frantjc/x/slice"
)

type allowHandler struct {
	handler      http.Handler
	allowMethods []string
}

func (ah *allowHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	Allow(w, r, ah.allowMethods, ah.handler)
}

func Allow(w http.ResponseWriter, r *http.Request, allowMethods []string, h http.Handler) {
	if xslice.Includes([]string{http.MethodHead, http.MethodOptions}, r.Method) {
		allow := strings.Join(allowMethods, ", ")
		w.Header().Set("Allow", allow)

		if http.MethodOptions == r.Method {
			w.Header().Set("Access-Control-Allow-Methods", allow)
			w.WriteHeader(http.StatusOK)
			return
		}
	}

	if !xslice.Includes(allowMethods, r.Method) {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	h.ServeHTTP(w, r)
}

func AllowHandler(allowMethods []string, h http.Handler) http.Handler {
	return &allowHandler{handler: h, allowMethods: allowMethods}
}
