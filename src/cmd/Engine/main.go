package main

import (
	"SearchEngine/src/libs/Logger"
)

func main() {
	log := Logger.New()
	log.WithField("test", "123").Error("Errordafdsad")

}
