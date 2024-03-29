package main

import (
	"net/http"
	"testing"
)

func testIsTheHttpRequestSuccessful(t *testing.T, url string) {
	resp, err := http.Get(url)
	if err != nil {
		t.Fatalf("There was no response from the server: %s", url)
	}
	if resp.StatusCode > 299 {
		t.Fatalf("Response failed with the status code: %d", resp.StatusCode)
	}
}

func TestGetQuote(t *testing.T) {
	url := "https://api.kanye.rest"
	testIsTheHttpRequestSuccessful(t, url)
}
