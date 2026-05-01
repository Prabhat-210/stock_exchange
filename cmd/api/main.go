package main

import (
	"fmt"
	"log"
	"userAuth/internal/bootstrap"
)

func main() {
	app, err := bootstrap.Initialize()
	if err != nil {
		fmt.Errorf("Unable to initialize db %w", err)
	}
	defer app.Close()

	log.Println("application initialized successfully")
}
