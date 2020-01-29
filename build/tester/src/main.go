package main

import (
	"log"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	seedDataParams, err := GetSeedDataParams()
	if err != nil {
		log.Fatal(err)
	}
	SeedData(seedDataParams)
}