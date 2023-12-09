//User{Name: "gg"}
package main

import (
	
	actions "root/actions.user"
	color "root/colors"
)


func main() {

	actions.DeleteUser(1)
	// actions.CreateUser(actions.User{Name: "gg"})
	color.Color(color.Colors[4], "main package")
}



