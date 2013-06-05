package postmaster

import (
	"testing"
)

func TestValidate(t *testing.T) {
	// Mock
	c := make(chan *restMockObj, 1)
	post = restMock(c, nil, 100, nil)

	pm := New("apikey")
	addr := new(Address)
	addr.City = "Austin"
	addr.State = "TX"
	pm.Validate(addr)
	ret := <-c
	if ret.endpoint != "validate" {
		t.Error("wrong endpoint")
	}
	if ret.version != "v1" {
		t.Error("wrong version")
	}
	if len(ret.params) != 2 {
		t.Error("wrong params length")
	}
	if ret.params["city"] != "Austin" {
		t.Error("wrong param (city)")
	}
	if ret.params["state"] != "TX" {
		t.Error("wrong param (state)")
	}
}
