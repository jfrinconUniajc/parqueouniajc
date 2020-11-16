package main

import (
	"log"

	"github.com/jfrinconUniajc/parqueouniajc/bd"
	"github.com/jfrinconUniajc/parqueouniajc/handlers"
)

func main() {
	if bd.CheckConnection() == 0 {
		log.Fatal("Sin conexion")
		return
	}
	handlers.ControllerUSer()
}
