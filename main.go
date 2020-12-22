package main

import (
	"github.com/phpdi/clockin/httpserver"
)

//go build -o clockinbin main.go
func main() {
	httpserver.HttpServer()

}
