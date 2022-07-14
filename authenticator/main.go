package main

import (
	"log"
	"net/http"

	"github.com/ahmetsabri/go-auth/pkg/controllers"
	"github.com/ahmetsabri/go-auth/pkg/middleware"
	"github.com/gorilla/mux"
)

func main() {

	r := mux.NewRouter()

	r.HandleFunc("/auth/signup", controllers.SignUp).Methods(http.MethodPost)
	r.HandleFunc("/auth/login", controllers.Login).Methods(http.MethodPost)

	r.Handle("/user/{id}", middleware.Auth(http.HandlerFunc(me))).Methods(http.MethodPut)

	log.Println("Started on :8080")
	log.Fatal(http.ListenAndServe(":8080", r))

}

func me(res http.ResponseWriter, req *http.Request) {
	log.Println("You can update me")
}
