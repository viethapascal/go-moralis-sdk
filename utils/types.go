package utils

import "reflect"

func EncodeToQueryString(input interface{}) string {
	//rt := reflect.TypeOf(input)
	rv := reflect.ValueOf(input)
	return rv.String()
}
