package utils

import (
	"reflect"
	"runtime"
	"strings"
)

// solution from https://stackoverflow.com/questions/7052693/how-to-get-the-name-of-a-function-in-go
func GetFunctionName(i interface{}) string {
	tokens := strings.Split(runtime.FuncForPC(reflect.ValueOf(i).Pointer()).Name(), "/")
	return tokens[len(tokens)-1]
}
