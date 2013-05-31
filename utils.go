package postmaster

import (
	"fmt"
	"net/url"
	"strings"
	"reflect"
)


// urlencode joins parameters from map[string]string with ampersand (&), and
// also escapes their values
func urlencode(params map[string]string) string {
	arr := make([]string, len(params))
	for k, v := range params {
		if fmt.Sprintf("%s", v) != "" {
			arr = append(arr, fmt.Sprintf("%s=%s", k, url.QueryEscape(v)))
		}
	}
	return strings.Join(arr, "&")
}


// MapStruct converts struct to map[string]string, using fields' names as keys
// and fields' values as values.
func MapStruct(s interface{}) map[string]string {
	result := make(map[string]string)
	fields := reflect.TypeOf(s).NumField()
	for i := 0; i < fields; i++ {
		field := reflect.TypeOf(s).Field(i)
		var name string
		if json := field.Tag.Get("json"); json != "" {
			name = json
		} else {
			name = strings.ToLower(field.Name)
		}
		/*
		TODO nested objects! better to do this with recursion
		if type == struct {
			for x in value {
				map["parentName[name]"] = value
			}
		}
		yay for pseudocode
		*/
		value := fmt.Sprintf("%v", reflect.ValueOf(s).Field(i).Interface())
		if value != "" {
			result[name] = value
		}
	}
	return result
}
