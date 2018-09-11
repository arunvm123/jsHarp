package operations

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

func TestCodeRoute(t *testing.T) {

	// TODO: Check other paths
	os.Create("render.html")

	srv := httptest.NewServer(InitRoutes())
	defer srv.Close()

	b, err := json.Marshal(Code{
		HTML: "Hello",
		CSS:  "{ color: 'red'}",
		JS:   "alert('Hello')",
	})
	if err != nil {
		t.Fatalf("Error marshalling code struct: %v", err.Error())
	}

	byteReader := bytes.NewReader(b)

	res, err := http.Post(srv.URL+"/code", "application/json", byteReader)
	if err != nil {
		t.Fatal(err.Error())
	}

	if res.StatusCode != http.StatusOK {
		t.Fatalf("Expected StatusOK code, got %v", res.StatusCode)
	}

	os.Remove("render.html")
}

func TestRenderRoute(t *testing.T) {
	// TODO: Check other paths
	os.Create("render.html")

	srv := httptest.NewServer(InitRoutes())
	defer srv.Close()

	res, err := http.Get(srv.URL + "/render")
	if err != nil {
		t.Fatal(err.Error())
	}

	if res.StatusCode != http.StatusOK {
		t.Fatalf("Expected StatusOK code, got %v", res.StatusCode)
	}

	os.Remove("render.html")

}
