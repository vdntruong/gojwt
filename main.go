package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"

	"gojwt/auth"
	"gojwt/tjwt"

	"github.com/golang-jwt/jwt/v5"
)

func main() {
	gjwt := tjwt.NewGenerator("gojwt", "my-secret", jwt.SigningMethodHS256)
	authSvc := auth.NewAuthSvc(gjwt)
	mux := http.NewServeMux()

	mux.HandleFunc("POST /signin", func(w http.ResponseWriter, r *http.Request) {
		decoder := json.NewDecoder(r.Body)
		var t struct {
			Username string
			Password string
		}
		if err := decoder.Decode(&t); err != nil {
			w.Write([]byte("please submit with username and password"))
			return
		}

		token, _, err := authSvc.GetToken(r.Context(), t.Username, t.Password)
		if err != nil {
			log.Println("Failed to generate token: ", err)
			w.Write([]byte("failed to generate token"))
			return
		}

		w.Write([]byte(token))
		w.WriteHeader(http.StatusOK)
	})

	mux.HandleFunc("GET /task/{id}", func(w http.ResponseWriter, r *http.Request) {
		reqToken := r.Header.Get("Authorization")
		if reqToken == "" {
			w.Write([]byte("please submit with authorization token"))
			return
		}

		splitToken := strings.Split(reqToken, "Bearer ")
		token := splitToken[1]

		_, err := authSvc.VerifyToken(r.Context(), token)
		if err != nil {
			log.Println("Failed to verify token: ", err)
			w.Write([]byte("failed to verify the token"))
			return
		}

		id := r.PathValue("id")
		w.Write([]byte(fmt.Sprintf("task %s has been completed", id)))
		w.WriteHeader(http.StatusOK)
	})

	err := http.ListenAndServe("localhost:8080", mux)
	if err != nil {
		return
	}
}
