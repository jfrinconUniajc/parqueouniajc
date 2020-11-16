package main

import (
	"log"

	"github.com/parqueouniajc/bd"
	"github.com/parqueouniajc/handlers"
)

func main() {
	if bd.CheckConnection() == 0 {
		log.Fatal("Sin conexxion")
		return
	}
	handlers.ControllerUSer()
}
