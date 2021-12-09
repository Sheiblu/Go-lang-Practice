package model

import (
	"customPackage/views"
	"fmt"
)

func ReadAll() ([]views.PostResponse, error) {
	rows, err := con.Query("SELECT id, name, age FROM `user`")

	if err != nil {
		fmt.Println("Select Fail")
		return nil, err
	}
	users := []views.PostResponse{}
	for rows.Next() {
		id := 0
		data := views.PostResponse{}
		rows.Scan(&id, &data.Name, &data.Age)
		users = append(users, data)
	}
	return users, nil
}

func ReadById(id int) ([]views.PostResponse, error) {
	rows, err := con.Query("SELECT id, name, age FROM `user` where id = ?", id)
	// defer insertQ.Close()

	if err != nil {
		fmt.Println("Insert Fail")
		return nil, err
	}
	users := []views.PostResponse{}
	for rows.Next() {
		id := 0
		data := views.PostResponse{}
		rows.Scan(&id, &data.Name, &data.Age)
		users = append(users, data)
	}
	return users, nil
}
