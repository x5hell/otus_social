package converter

import (
	"fmt"
	"reflect"
	"structure"
)

func RowListToSqlInsertQuery(rowList []structure.Table) string {
	if len(rowList) == 0 {
		return ""
	}

	firstRow := rowList[0]
	//tableName := firstRow.GetTableName()
	propertyList := structure.GetTablePropertyList(firstRow)
	//fieldList := structure.GetTableFieldList(firstRow)
	/*sqlQueryStart := fmt.Sprintf(
		"INSERT INTO %s (%s) ",
		tableName,
		strings.Join(fieldList, ","))
*/
	for _, row := range rowList {
		rowValue := reflect.ValueOf(row)
		var rowValueList []string
		for _, property := range propertyList {
			fieldValue := rowValue.FieldByName(property).Interface()
			field, _ := fieldValue.(structure.NullString)
			fieldQuery := NullStringToQuery(field)
			rowValueList = append(rowValueList, fieldQuery)
		}
		fmt.Println(rowValueList)
	}
	return ""
}

func NullStringToQuery (field structure.NullString) string {
	if field.Valid {
		return "'" + field.String + "'"
	} else {
		return "NULL"
	}
}
