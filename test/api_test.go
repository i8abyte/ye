package main

import (
	"testing"

	ye "github.com/imran-salim/ye/api"
)

func testDataEmpty(t *testing.T, data []byte) {
	if len(data) < 1 {
		t.Errorf("The data has %d bytes", len(data))
	}
}

func TestGetQuote(t *testing.T) {
	data := ye.GetQuote()

	testDataEmpty(t, data)
}
