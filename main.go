package main

import (
	"github.com/gelleson/wait/internal/app"
	"log"
)

func main() {
	if err := app.Execute(); err != nil {
		log.Fatalln(err)
	}
}
