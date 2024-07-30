package main

import (
	"context"
	"fmt"
	"log"

	"github.com/elv-todd/solana-go-sdk/client"
	"github.com/elv-todd/solana-go-sdk/rpc"
)

func main() {
	c := client.NewClient(rpc.DevnetRPCEndpoint)
	ok, err := c.GetHealth(context.TODO())
	if err != nil {
		log.Fatalf("failed to get health, err: %v", err)
	}

	fmt.Println(ok)
}
