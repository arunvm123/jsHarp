package operations

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

func TestRetrieveCodeHandler(t *testing.T) {
	tt := []struct {
		name string
		code Code
		err  string
	}{
		{name: "Render file", code: Code{
			HTML: "Hello",
			CSS:  "{ color: 'red'}",
			JS:   "alert('Hello')",
		}},
		{name: "Missing request body", err: "Error in request body data"},
		{name: "Missing file", err: "Could not open file", code: Code{HTML: "Hello"}},
	}

	var byteReader *bytes.Reader

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {

			if tc.err == "" {
				os.Create("render.html")
			}

			if tc.code != (Code{}) {
				b, err := json.Marshal(tc.code)
				if err != nil {
					t.Fatalf("Error marshalling code struct: %v", err.Error())
				}

				byteReader = bytes.NewReader(b)
			}

			req, err := http.NewRequest("POST", "/code", byteReader)
			if err != nil {
				t.Fatalf("Failed to construct request: %v", err.Error())
			}

			rec := httptest.NewRecorder()
			retrieveCode().ServeHTTP(rec, req)

			res := rec.Result()
			defer res.Body.Close()

			b, err := ioutil.ReadAll(res.Body)
			if err != nil {
				t.Fatalf("could not read response: %v", err)
			}

			if tc.err != "" {
				if (res.StatusCode != http.StatusBadRequest) && (res.StatusCode != http.StatusInternalServerError) {
					t.Errorf("expected error code; got %v", res.StatusCode)
				}
				if msg := string(bytes.TrimSpace(b)); msg != tc.err {
					t.Errorf("expected message %q; got %q", tc.err, msg)
				}

				return
			}

			if res.StatusCode != http.StatusOK {
				t.Fatalf("Expected StatusOK code, got %v", res.StatusCode)
			}

			if _, err := os.Stat("render.html"); err == nil {
				os.Remove("render.html")
			}
		})

	}

}

func TestRenderPage(t *testing.T) {
	tt := []struct {
		name string
		err  string
	}{
		{name: "Successful Response"},
		{name: "Error Response"},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {

			if tc.name != "Error Response" {
				os.Create("render.html")
			}

			req, err := http.NewRequest("GET", "/render", nil)
			if err != nil {
				t.Fatalf("Failed to construct request: %v", err.Error())
			}

			rec := httptest.NewRecorder()
			renderPage().ServeHTTP(rec, req)

			res := rec.Result()
			defer res.Body.Close()

			if tc.name == "Error Response" {
				if res.StatusCode != http.StatusNotFound {
					t.Fatalf("Expected StatusNotFound code, got %v", res.StatusCode)
				}
				return
			}

			if res.StatusCode != http.StatusOK {
				t.Fatalf("Expected StatusOK code, got %v", res.StatusCode)
			}

			if _, err := os.Stat("render.html"); err == nil {
				os.Remove("render.html")
			}
		})
	}
}
