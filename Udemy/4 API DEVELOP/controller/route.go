package controller

import (
	"net/http"
)

func Register() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", ping())
	mux.HandleFunc("/pong", pong())
	mux.HandleFunc("/create", CreateUser())
	mux.HandleFunc("/read", ReadAll())
	mux.HandleFunc("/readById", ReadById())
	mux.HandleFunc("/deleteById", DeleteById())

	return mux
}
