package postmaster

import (
	"testing"
)

func TestTrackExternal(t *testing.T) {
	// Mock
	var err error
	c := make(chan *restMockObj, 1)
	post = restMock(c, nil, 100, nil)

	pm := New("apikey")
	tr := pm.TrackingExternal()

	_, err = tr.Put()
	if err != nil {
		t.Error("err should be nil")
	}
	ret := <-c
	if ret.endpoint != "track" {
		t.Error("wrong endpoint")
	}
	if ret.version != "v1" {
		t.Error("wrong version")
	}
}

func TestTrackRef(t *testing.T) {
	// Mock
	var err error
	c := make(chan *restMockObj, 1)
	get = restMockGet(c, nil, 100, nil)

	pm := New("apikey")

	_, err = pm.TrackRef("abcde")
	if err != nil {
		t.Error("err should be nil")
	}
	ret := <-c
	if ret.endpoint != "track" {
		t.Error("wrong endpoint")
	}
	if ret.version != "v1" {
		t.Error("wrong version")
	}
}
