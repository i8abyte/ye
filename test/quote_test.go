package main

import (
	"encoding/json"
	"strings"
	"testing"

	goaway "github.com/TwiN/go-away"
	ye "github.com/imran-salim/ye/api"
)

type Response struct {
	Quote string
}

func testRespObjEmpty(t *testing.T, respObj Response) {
	if len(respObj.Quote) < 1 {
		t.Errorf("The quote in the response is empty: %s", respObj.Quote)
	}
}

func testCensor(t *testing.T, quote string) {
	if goaway.IsProfane(quote) {
		censoredQuote := goaway.Censor(quote)

		if !strings.Contains(censoredQuote, "*") {
			t.Errorf("The quote could not be censored despite containing profanity")
		}
	}
}

func TestQuoteKanyeWest(t *testing.T) {
	data := ye.GetQuote()
	var respObj Response
	json.Unmarshal(data, &respObj)
	quote := respObj.Quote

	testRespObjEmpty(t, respObj)
	testCensor(t, quote)
}