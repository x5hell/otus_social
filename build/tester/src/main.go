package main

import (
	"component/databaseWorkMode"
	"flag"
	"math/rand"
	"os"
	"test/helper"
	"test/wrkTest"
	"time"
)

func main() {
	seedDataParams := helper.GetSeedDataParams()

	action := flag.String(
		"action",
		"applyFixture",
		"available actions: applyFixture, removeIndex, AddIndexes, testWithoutIndex, testWithIndex, setAppWorkModeUseReplica, setAppWorkModeMasterOnly",
		)
	flag.Parse()

	masterHostname := os.ExpandEnv("$MYSQL_MASTER_HOSTNAME")
	slaveHostname := os.ExpandEnv("$MYSQL_SLAVE_HOSTNAME")

	switch *action {
		case "applyFixture":
			rand.Seed(time.Now().UnixNano())
			helper.ApplyFixture(seedDataParams, masterHostname)
			break
		case "removeIndex":
			helper.RemoveIndexes(seedDataParams, slaveHostname)
			break
		case "addIndex":
			helper.AddIndexes(seedDataParams, slaveHostname)
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
		case "setAppWorkModeUseReplica":
			_ = databaseWorkMode.Set("UseReplica")
			break
		case "setAppWorkModeMasterOnly":
			_ = databaseWorkMode.Set("MasterOnly")
			break
	}
}