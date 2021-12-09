package main

import  (
	"customPackage/controller"
	"customPackage/model"
	"net/http"
)


func main() {
	mux := controller.Register()

	// Create connection
	db := model.Connection()
	defer db.Close()
	http.ListenAndServe(":3000", mux)
}
