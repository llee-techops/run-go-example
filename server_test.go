package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"testing"
)

func TestOutput(t *testing.T) {

	resp, err := http.Get("http://localhost:5000")
	if err != nil {
		t.Fatal(err)
	} else {
		defer resp.Body.Close()
		contents, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			fmt.Printf("%s", err)
			os.Exit(1)
		}
		if string(contents) != "Hello!\n" {
			t.Fatal("Not right content")
		}
	}
}
