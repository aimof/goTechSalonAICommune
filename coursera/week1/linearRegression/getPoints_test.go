package main

import "testing"

func TestReadConfig(t *testing.T) {
	conf := readConfig()
	actual := conf.inputFileName
	expected := "~/Work/tech_salon_AI_commune/coursera/week1/leastSquare/json/sample.json"
	if actual != expected {
		t.Errorf("got %s\nwant %v", actual, expected)
	}
}
