package controllers

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/ahmetsabri/go-auth/models/token"
	"github.com/ahmetsabri/go-auth/models/user"
	"github.com/ahmetsabri/go-auth/pkg/helpers"
)

type ResponseSuccess struct {
	Token string
	User  user.User `json:"user"`
}

type ResponseError struct {
	Message string `json:"message"`
}

type LoginRequest struct {
	Email    string
	Password string
}

func SignUp(res http.ResponseWriter, req *http.Request) {
	var u user.User
	var r ResponseSuccess
	var data ResponseError

	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-type", "application/json")

	err := json.NewDecoder(req.Body).Decode(&u)

	if err != nil {
		panic(err)
	}

	t := user.Create(&u)

	if len(t) == 0 {
		data.Message = "Error while creating new user , check the log"
		res.WriteHeader(422)
		json.NewEncoder(res).Encode(data)
		return
	}

	r = ResponseSuccess{
		User:  u,
		Token: t,
	}

	json.NewEncoder(res).Encode(r)
}

func Login(res http.ResponseWriter, req *http.Request) {

	var u user.User
	var body LoginRequest
	var r ResponseSuccess
	var data ResponseError

	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-type", "application/json")

	err := json.NewDecoder(req.Body).Decode(&body)
	if err != nil {
		log.Fatal(err)
	}

	result := helpers.DB.Last(&u, "email = ?", body.Email)

	if result.RowsAffected == 0 {
		data.Message = "Error in credentials"
		res.WriteHeader(401)
		json.NewEncoder(res).Encode(data)
		return
	}

	err = (user.CheckPassword(u.Password, body.Password))

	if err != nil {
		data.Message = "Error in credentials"
		res.WriteHeader(401)
		json.NewEncoder(res).Encode(data)
		return
	}

	t := token.GenerateToken(body.Email)
	r = ResponseSuccess{
		User:  u,
		Token: t,
	}

	saveToken := &token.Token{
		Token:     t,
		UserId:    u.ID,
		ExpiredAt: time.Now().Add(1 * time.Hour),
	}

	token.Create(saveToken)
	json.NewEncoder(res).Encode(r)
}
