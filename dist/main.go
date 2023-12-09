package main

import (
	"fmt"
	"reflect"
)

type Person struct {
	Name   string
	Age    int
	Gender string
	Single bool
}

func main() {
	var data = []Person{
		{
			Name:   "John",
			Gender: "Female",
			Age:    17,
			Single: false,
		},
		{
			Name:   "ааа",
			Gender: "апппп",
			Age:    300,
			Single: false,
		},
	}

	for i := range data {
		values := reflect.ValueOf(data[i])
		types := values.Type()
			for i := 0; i < values.NumField(); i++ {
				// fmt.Println(types.Field(i).Index[0], types.Field(i).Name, values.Field(i))
				
			
				if types.Field(i).Index[0] == 0 &&  values.Field(i).String() == "John" {
					fmt.Print("deleted \n")
				}
			}
	}

	// values := reflect.ValueOf(data)
	// types := values.Type()
	// for i := 0; i < values.NumField(); i++ {
	// 	fmt.Println(types.Field(i).Index[0], types.Field(i).Name, values.Field(i))
	// }

	fmt.Println(data)
}


// func DeleteUser(data User) {
// 	file, err := os.ReadFile("data.json")
// 	var u2 []User
// 	json.Unmarshal(file, &u2)
// 	if err != nil {
// 		panic(err)
// 	}
// 	for i := range u2 {
// 		if u2[i] == data {
// 			u2[i].Reset()
// 			fmt.Print(u2[i], "\n", u2, "\n", data, "\n", db, "\n", )
// 			db = append(db, u2...)
// 			jsonData, err := json.Marshal(db)
// 			color.Color(color.Colors[3], "Выполнено")
// 			os.WriteFile("data.json", jsonData, 0644)
// 			if err != nil {
// 				panic(err)
// 			}
// 		}
		
// 	}
	
// }
