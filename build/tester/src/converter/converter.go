package converter

import (
	"fmt"
	"reflect"
	"strings"
	"structure"
)

func RowListToSqlInsertQuery(rowList []structure.Table) string {
	if len(rowList) == 0 {
		return ""
	}

	firstRow := rowList[0]
	tableName := firstRow.GetTableName()
	propertyList := structure.GetTablePropertyList(firstRow)
	fieldList := structure.GetTableFieldList(firstRow)
	sqlQueryStart := fmt.Sprintf(
		"INSERT INTO %s (%s)",
		tableName,
		strings.Join(fieldList, ","))

	sqlQueryValues := RowListToSqlInsertQueryValues(rowList, propertyList)
	return fmt.Sprintf("%s VALUES %s;", sqlQueryStart, sqlQueryValues)
}

func RowListToSqlInsertQueryValues(rowList []structure.Table, propertyList []string) string {
	var rowListQuery []string
	for _, row := range rowList {
		rowValue := reflect.ValueOf(row)
		var rowValueList []string
		for _, property := range propertyList {
			fieldValue := rowValue.FieldByName(property).Interface()
			field, _ := fieldValue.(structure.NullString)
			fieldQuery := NullStringToQuery(field)
			rowValueList = append(rowValueList, fieldQuery)
		}
		rowQuery := "(" + strings.Join(rowValueList, ",") + ")"
		rowListQuery = append(rowListQuery, rowQuery)
	}
	return strings.Join(rowListQuery, ",")
}



func NullStringToQuery (field structure.NullString) string {
	if field.Valid {
		return "'" + field.String + "'"
	} else {
		return "NULL"
	}
}
