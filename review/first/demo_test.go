package source_test

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/dlsniper/u/review/first"
)

var billData = `some text/json/xml`


func TestGetSourceFromUrl(t *testing.T) {
	f := func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprint(w, billData)
	}
	server := httptest.NewServer(http.HandlerFunc(f))
	srv := mockedServer()
	defer srv.Close()

	resp, err := http.Get(srv.URL)
	defer resp.Body.Close()
	if err != nil {
		t.Logf("Should be making a Get call, but got: %v", err)
		t.Fail()
	}


	t.Log("Made get request")

	sOK := http.StatusOK
	if resp.StatusCode != sOK {
		t.Fatalf("Should receive status %d, but received: %d", sOK, resp.StatusCode)
	}
	t.Log("Status code is", http.StatusOK)

	b, err := source.GetData(srv.URL)
	if b != billData {
		t.Fatal("Should be able to match element %s, but got %s", billData, b)
	}
	t.Logf("Body matched %s matched", string(b))
}
