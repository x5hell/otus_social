package main

import (
	"fmt"
	"log"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	seedDataParams, err := GetSeedDataParams()

	fmt.Println(seedDataParams)

	if err != nil {
		log.Fatal(err)
	}
	SeedData(seedDataParams)
}