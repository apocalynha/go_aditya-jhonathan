package main

import (
	"prioritas2/config"
	"prioritas2/route"
)

func main() {
	// create a new echo instance
	config.InitDB()
	e := route.New()
	// start the server, and log if it fails
	e.Logger.Fatal(e.Start(":8000"))
}
