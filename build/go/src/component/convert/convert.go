package convert

func StringListToInterfaceList(stringList []string) []interface{} {
	interfaceList := make([]interface{}, len(stringList))
	for index, value := range stringList {
		interfaceList[index] = value
	}
	return interfaceList
}

func Int64ListToInterfaceList(int64List []int64) []interface{} {
	interfaceList := make([]interface{}, len(int64List))
	for index, value := range int64List {
		interfaceList[index] = value
	}
	return interfaceList
}
