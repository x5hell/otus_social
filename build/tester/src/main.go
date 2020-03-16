package main

import (
	"flag"
	"math/rand"
	"test/helper"
	"test/wrkTest"
	"time"
)

func main() {
	seedDataParams := helper.GetSeedDataParams()

	action := flag.String(
		"action",
		"applyFixture",
		"available actions: applyFixture, removeIndex, AddIndexes, testWithoutIndex, testWithIndex",
		)
	flag.Parse()

	switch *action {
		case "applyFixture":
			rand.Seed(time.Now().UnixNano())
			helper.ApplyFixture(seedDataParams)
			break
		case "removeIndex":
			helper.RemoveIndexes(seedDataParams)
			break
		case "addIndex":
			helper.AddIndexes(seedDataParams)
			break
		case "testWithoutIndex":
			wrkTest.TestLatencyAndThroughput(
				"/search?first-name=%D0%B5%D0%B2&last-name=%D0%BA%D1%83",
				*action + "_",
			)
			break
		case "testWithIndex":
			wrkTest.TestLatencyAndThroughput(
				"/search?first-name=%D0%B5%D0%B2&last-name=%D0%BA%D1%83",
				*action + "_",
			)
			break
	}
}