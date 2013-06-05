package postmaster

import (
	"fmt"
	"net/url"
	"reflect"
	"strings"
)

// urlencode joins parameters from map[string]string with ampersand (&), and
// also escapes their values.
func urlencode(params map[string]string) string {
	arr := make([]string, 0)
	for k, v := range params {
		if fmt.Sprintf("%s", v) != "" {
			arr = append(arr, fmt.Sprintf("%s=%s", k, url.QueryEscape(v)))
		}
	}
	return "&" + strings.Join(arr, "&") + "&"
}

// mapStruct converts struct to map[string]string, using fields' names as keys
// and fields' values as values.
// It also automagically converts any nested structures.
func mapStruct(s interface{}) map[string]string {
	return mapStructNested(s, "")
}

// mapStructNested does all the dirty job that mapStruct was too lazy to do.
func mapStructNested(s interface{}, baseName string) map[string]string {
	result := make(map[string]string)
	// Is s a pointer? We don't want any of those here
	for reflect.TypeOf(s).Kind() == reflect.Ptr {
		s = reflect.ValueOf(s).Elem().Interface()
	}
	fields := reflect.TypeOf(s).NumField()
	for i := 0; i < fields; i++ {
		t := reflect.TypeOf(s).Field(i)
		v := reflect.ValueOf(s).Field(i)
		// Do we even need to parse this field?
		if t.Tag.Get("dontMap") == "true" {
			continue
		}
		// Name is important
		var name string
		if json := t.Tag.Get("json"); json != "" {
			name = json
		} else {
			name = strings.ToLower(t.Name)
		}
		if baseName != "" {
			name = fmt.Sprintf("%s[%s]", baseName, name)
		}
		// I wonder whether this is a nested object
		if v.Kind() == reflect.Struct || v.Kind() == reflect.Ptr { // Nested, activate recursion!
			if v.IsNil() {
				continue
			}
			m := mapStructNested(v.Interface(), name)
			for mk, mv := range m {
				result[mk] = mv
			}
		} else { // Not nested
			value := fmt.Sprintf("%v", v.Interface())
			// Omit all zeros
			k := v.Kind()
			if (k == reflect.Float32 || k == reflect.Int) && value == "0" || value == "" {
				continue
			}
			result[name] = value
		}
	}
	return result
}

// makeUrl creates full URL from baseUrl, version and endpoint.
func (p *Postmaster) makeUrl(version string, endpoint string) string {
	var url string
	if p.baseUrl != "" {
		url = p.baseUrl
	} else {
		url = "https://api.postmaster.io"
	}
	return fmt.Sprintf("%s/%s/%s", url, version, endpoint)
}

// restMockObj is being sent to test case via a buffered channel to make sure
// REST function was called with proper arguments.
type restMockObj struct {
	version  string
	endpoint string
	params   map[string]string
}

// restMock replaces function from rest.go file and just returns given object.
// It communicates with test case via a buffered channel.
func restMock(c chan *restMockObj, mocked interface{}, s int, err error) func(p *Postmaster, version string, endpoint string, params map[string]string, result interface{}) (status int, e error) {
	return func(p *Postmaster, version string, endpoint string, params map[string]string, result interface{}) (status int, e error) {
		result = mocked
		c <- &restMockObj{version: version, endpoint: endpoint, params: params}
		return s, err
	}
}
