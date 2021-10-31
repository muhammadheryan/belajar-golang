package middleware

import (
	"belajar-golang/belajar-restful-api/helper"
	"belajar-golang/belajar-restful-api/model/web"
	"net/http"
)

type AuthMiddleware struct {
	Handler http.Handler
}

func NewAuthMiddleware(handler http.Handler) *AuthMiddleware {
	return &AuthMiddleware{Handler: handler}
}

func (middleware *AuthMiddleware) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if "x666x" == r.Header.Get("X-API-Key") {
		//API KEY TRUE
		middleware.Handler.ServeHTTP(w, r)
	} else {
		//API KEY FALSE
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusUnauthorized)

		webResponse := web.WebResponse{
			Code:   http.StatusNotFound,
			Status: "Unauthorized Request",
		}

		helper.WriteToResponseBody(w, webResponse)
	}
}
