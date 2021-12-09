package controller

import (
	"customPackage/model"
	"customPackage/views"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

func CreateUser() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {
			data := views.PostResponse{}
			json.NewDecoder(r.Body).Decode(&data)
			fmt.Println(data)
			response := views.Response2{}

			if err := model.CreateUser(data.Name, data.Age); err != nil {
				// w.Write([]byte("Some Error"))
				response.Success = false
				response.Message = "Data Not save."
			} else {
				response.Success = true
				response.Message = "Name:  successfully save."
			}
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(response)
		}
	}
}

func ReadAll() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		if r.Method == http.MethodGet {
			data, err := model.ReadAll()

			response := views.Response{}
			if err != nil {
				response.Body = []int{}
				response.Code = 400
			} else {
				response.Body = data
				response.Code = 200
			}
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(response)
		}
	}
}

func ReadById() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		if r.Method == http.MethodGet {
			id, err2 := strconv.ParseInt(r.URL.Query().Get("id"), 10, 64)

			if err2 != nil {
				id = 0
			}

			data, err := model.ReadById(int(id))

			response := views.Response{}
			if err != nil {
				response.Body = []int{}
				response.Code = 400
			} else {
				response.Body = data
				response.Code = 200
			}
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(response)
		}
	}
}

func DeleteById() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		if r.Method == http.MethodDelete {
			id, err2 := strconv.ParseInt(r.URL.Query().Get("id"), 10, 64) //int64convert
			if err2 != nil {
				id = 0
			}
			err := model.DeleteById(int(id))
			if err != nil {
				json.NewEncoder(w).Encode(struct {
					Success bool   `json:success`
					Message string `json:message`
				}{false, "Fails to delete data."})
			}
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(struct {
				Success bool   `json:success`
				Message string `json:message`
			}{true, "Successfully data delete."})
		}
	}
}
