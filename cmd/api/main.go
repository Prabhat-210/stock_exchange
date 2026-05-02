package main

import (
	"context"
	"log"
	"userAuth/internal/bootstrap"
)

func main() {
	ctx := context.Background()

	app, err := bootstrap.Initialize(ctx)
	
	if err != nil {
		log.Fatalf("Unable to initialize db %v", err)
	}
	defer app.Close()

	log.Println("application initialized successfully")
}
