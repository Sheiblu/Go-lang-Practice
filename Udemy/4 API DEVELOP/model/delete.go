package model

import (
	"fmt"
)

func DeleteById(id int) error {
	deleteQ, err := con.Query("DELETE FROM `user` WHERE id = ?", id)
	defer deleteQ.Close()

	if err != nil {
		fmt.Println("Insert Fail")
		return err
	}

	return nil
}
