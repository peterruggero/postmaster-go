package postmaster

import (
	"testing"
	"reflect"
)

func TestTime(t *testing.T) {
	// Mock
	c := make(chan *restMockObj, 1)
	post = restMock(c, nil, 100, nil)

	pm := New("apikey")
	time := new(TimeMessage)
	tr, _ := pm.Time(time)
	ret := <-c
	if ret.endpoint != "times" {
		t.Error("wrong endpoint")
	}
	if ret.version != "v1" {
		t.Error("wrong version")
	}
	if reflect.TypeOf(tr) != reflect.TypeOf(new(TimeResponse)) {
		t.Error("wrong response type")
	}
}
