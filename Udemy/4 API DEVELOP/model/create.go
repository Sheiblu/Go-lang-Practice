package model

import "fmt"

func CreateUser(name string, age int) error {
	insertQ, err := con.Query("INSERT INTO `user` (name, age) VALUES (?, ?)", name, age)
	defer insertQ.Close()

	if err != nil {
		fmt.Println("Insert Fail")
		return err
	} else {
		fmt.Println("Insert to Database")
	}
	return nil
}
