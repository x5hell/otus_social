package main

import (
	"component/fixture"
	"generator"
	"log"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	seedDataParams, err := generator.GetSeedDataParams()
	if err != nil {
		log.Fatal(err)
	}
	err = generator.GenerateFixture(seedDataParams)
	if err != nil {
		log.Fatal(err)
	}
	err = fixture.Apply(seedDataParams.FixtureGeneratedScript)
	if err != nil {
		log.Fatal(err)
	}
}