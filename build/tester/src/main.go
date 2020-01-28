package main

import (
	"converter"
	"fmt"
	"generator"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	fmt.Println(
		converter.RowListToSqlInsertQuery(generator.CityRows(10)),
		"\n",
		converter.RowListToSqlInsertQuery(generator.InterestRows(10)),
		"\n",
		converter.RowListToSqlInsertQuery(generator.UserRows(10, 10, 16, 60)),

		)

}