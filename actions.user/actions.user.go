package actions

import (
	"encoding/json"
	"fmt"
	"os"
	color "root/colors"
)

var db = []User{}

var C = &User{}

func (a *User) Reset() {
    *a = *C
}

func CreateUser(data User ) {
	urs := User{
		Id: data.Id,
		Name: data.Name,
	}
	db = append(db, urs)
	jsonData, err := json.Marshal(db)
	os.WriteFile("data.json", jsonData , 0644)
	color.Color(color.Colors[3], "Выполнено")
	if err != nil {
		panic(err)
	}
}


func DeleteUser(id int) {
	file, err := os.ReadFile("data.json")
	var u2 []User
	json.Unmarshal(file, &u2)
	if err != nil {
		panic(err)
	}
	db = append(db, u2[id:]...)
	jsonData, err := json.Marshal(db)
	fmt.Print(db)
	os.WriteFile("data.json", jsonData, 0644)
	color.Color(color.Colors[3], "Выполнено")
	if err != nil {
		panic(err)
	}
}

