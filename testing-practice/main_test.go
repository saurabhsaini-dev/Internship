package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

type AddResult struct {
	x        int
	y        int
	expected int
}

var addResult = []AddResult{
	{1, 1, 2},
}

// Testing of Add function
func TestAdd(t *testing.T) {
	for _, test := range addResult {
		result := Add(test.x, test.y)
		if result != test.expected {
			t.Fatal("Expected result not given")
		}
	}
}

// Testing of reading a file
func TestReadFile(t *testing.T) {
	data, err := ioutil.ReadFile("testdata/test.data")
	if err != nil {
		t.Fatal("Could not open file")
	}
	if string(data) != "hello world" {
		t.Fatal("String contents do not match expected")
	}
}

// Testing of http request
func TestHttpRequest(t *testing.T) {
	handler := func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "{ \"status\": \"good\" }")
	}

	req := httptest.NewRequest("GET", "https://tutorialedge.net", nil)
	w := httptest.NewRecorder()
	handler(w, req)

	resp := w.Result()
	body, _ := ioutil.ReadAll(resp.Body)

	fmt.Println(string(body))

	if 200 != resp.StatusCode {
		t.Fatal("Status code not OK")
	}

}
