package convert

import (
	"component/handler"
	"crypto/md5"
	"fmt"
	"io"
)

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

func StringToMd5(str string) (hash string) {
	h := md5.New()
	_, err := io.WriteString(h, str)
	handler.ErrorLog(err)
	return fmt.Sprintf("%x", h.Sum(nil))
}
