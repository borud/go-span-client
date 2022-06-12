package main

import (
	"context"
	"fmt"
	"log"
	"time"

	client "github.com/borud/go-span-client"
)

func main() {
	config := client.NewDefaultConfig()
	client, err := client.COAPConnect(config)
	if err != nil {
		log.Fatalf("COAPConnect failed: %v", err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Minute)
	defer cancel()

	// buf := bytes.NewReader([]byte("hello"))
	// msg, err := client.Post(ctx, "/", message.TextPlain, buf)
	res, err := client.Get(ctx, "/")
	if err != nil {
		log.Fatalf("post failed: %v", err)
	}

	log.Printf("result: %+v", res)

	body, err := res.ReadBody()
	if err != nil {
		log.Fatalf("ReadBody failed: %v", err)
	}

	fmt.Printf("len=%d body=[%s]\n", len(body), string(body))
}
