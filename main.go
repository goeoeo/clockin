package main

import (
	"github.com/phpdi/clockin/core"
	"github.com/phpdi/clockin/httpserver"
	"os"
)

//go build -o clockinbin main.go
func main() {
	httpserver.HttpServer()
	//cmd()
}

func cmd() {

	if len(os.Args) > 0 {
		core.Run(os.Args[0])
		return
	}

	core.Run("")
}
