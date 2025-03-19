package main

import (
	"sinarmas/app"
)

func main() {
	var a app.App
	a.CreateConnection()
	a.Routes()
	a.Run()
}
