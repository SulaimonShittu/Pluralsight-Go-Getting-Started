package main

import (
	"Pluralsight-Go-Getting-Started/controllers"
	"net/http"
)

func main() {
	controllers.RegisterControllers()
	http.ListenAndServe(":3000", nil)
}
