package main

import (
	"github.com/PxGo/GoFM/modules"
)

func main() {

	modules.InitReader()
	modules.InitServer()

	select {}
}
