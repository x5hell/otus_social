package main

import (
	"converter"
	"fmt"
	"generator"
	"math/rand"
	"structure"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	fmt.Println(generator.City())

	var tableList []structure.Table
	city1 := structure.City{ID: structure.NullString{String:"1", Valid:true}, Name:structure.NullString{String:"Moscow", Valid:true}}
	city2 := structure.City{ID: structure.NullString{String:"2", Valid:true}, Name:structure.NullString{String:"Vladimir", Valid:true}}

	tableList = append(tableList, city1)
	tableList = append(tableList, city2)

	converter.RowListToSqlInsertQuery(tableList)

}