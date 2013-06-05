package postmaster

import (
	"testing"
	"reflect"
)

func TestRate(t *testing.T) {
	// Mock
	c := make(chan *restMockObj, 1)
	post = restMock(c, nil, 100, nil)

	pm := New("apikey")
	r := new(RateMessage)
	// Non-empty carrier
	r.Carrier = "ups"
	tr, _ := pm.Rate(r)
	ret := <-c
	if ret.endpoint != "rates" {
		t.Error("wrong endpoint")
	}
	if ret.version != "v1" {
		t.Error("wrong version")
	}
	if reflect.TypeOf(tr) != reflect.TypeOf(new(RateResponse)) {
		t.Error("wrong response for non-empty carrier")
	}
	// Empty carrier
	r.Carrier = ""
	tr, _ = pm.Rate(r)
	if reflect.TypeOf(tr) != reflect.TypeOf(new(RateResponseBest)) {
		t.Error("wrong response for empty carrier")
	}
}
