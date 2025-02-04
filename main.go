package main

import (
	"Pluralsight-Go-Getting-Started/models"
	"fmt"
)

func main() {
	u := models.User{
		ID:        2,
		FirstName: "Sulaimon",
		LastName:  "Shittu",
	}
	fmt.Println(u)
}
