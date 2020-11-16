package main

import (
	"log"

	"./bd"
	"./handlers"
)

func main() {
	if bd.CheckConnection() == 0 {
		log.Fatal("Sin conexxion")
		return
	}
	handlers.ControllerUSer()
}
