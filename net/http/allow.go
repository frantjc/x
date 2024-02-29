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
	if xslice.Includes([]string{http.MethodHead, http.MethodOptions}, r.Method) {
		allow := strings.Join(ah.allowMethods, ", ")
		w.Header().Set("Allow", allow)

		if http.MethodOptions == r.Method {
			w.Header().Set("Access-Control-Allow-Methods", allow)
			w.WriteHeader(http.StatusOK)
			return
		}
	}

	if !xslice.Includes(ah.allowMethods, r.Method) {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	ah.handler.ServeHTTP(w, r)
}

func AllowHandler(h http.Handler, allowMethods []string) http.Handler {
	return &allowHandler{handler: h, allowMethods: allowMethods}
}
