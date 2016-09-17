package main

import "testing"

func TestReadConfig(t *testing.T) {
	conf := readConfig()
	actual := conf.InputFileName
	expected := "/json/sample.json"
	if actual != expected {
		t.Errorf("got %s\nwant %s", actual, expected)
	}
}
