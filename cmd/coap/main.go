package main

import (
	"context"
	"fmt"
	"log"
	"time"

	client "github.com/borud/go-span-client"
)

func main() {
	client, err := client.COAPConnect(client.NewDefaultConfig())
	if err != nil {
		log.Fatalf("COAPConnect failed: %v", err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Minute)
	defer cancel()

	res, err := client.Get(ctx, "")
	if err != nil {
		log.Fatalf("get failed: %v", err)
	}

	body, err := res.ReadBody()
	if err != nil {
		log.Fatalf("ReadBody failed: %v", err)
	}

	fmt.Println(body)
}
