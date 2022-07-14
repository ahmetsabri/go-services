package middleware

import (
	"log"
	"net/http"
	"strconv"

	"github.com/ahmetsabri/go-auth/models/token"
	"github.com/gorilla/mux"
)

func Auth(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Contenet-Type", "application/json")
		w.Header().Set("Accept", "application/json")

		authToken := r.Header.Get("token")

		t := token.Verify(authToken)

		if t.UserId == 0 {
			w.WriteHeader(401)
			return
		}

		vars := mux.Vars(r)
		id, err := strconv.Atoi(vars["id"])
		if err != nil {
			w.WriteHeader(403)
			return
		}

		if int(t.UserId) != id {
			log.Println("not authorized", t.UserId, id)
			w.WriteHeader(403)
			return
		}

		h.ServeHTTP(w, r)
	})
}
