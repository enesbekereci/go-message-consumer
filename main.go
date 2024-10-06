package main

import (
	"fmt"
	"log"
	"sync"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env %s", err)
	}
	fmt.Println("MQ-CONSUMER")
	cc := ClientManager{}
	cc.SetClients()

	var wg sync.WaitGroup
	cc.ConsumeClients(&wg)
	wg.Wait()
}
