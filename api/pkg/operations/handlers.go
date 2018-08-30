package operations

import (
	"encoding/json"
	"net/http"
	"os"
)

func retrieveCode() http.Handler {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			var code *Code

			decoder := json.NewDecoder(r.Body)
			err := decoder.Decode(&code)
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				w.Write([]byte(err.Error()))
				return
			}

			fp, err := os.OpenFile("render.html", os.O_RDWR, 0644)
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				w.Write([]byte(err.Error()))
				return
			}
			defer fp.Close()

			// Clear the file before writing to it
			err = fp.Truncate(0)
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				w.Write([]byte(err.Error()))
				return
			}

			// Sets the pointer to the beggining of the file
			fp.Seek(0, 0)

			fp.Write([]byte(`
				<html>
					<head>
						<style>` + code.CSS + `</style>
						<script>` + code.JS + `</script>
					</head>
					<body>` + code.HTML + `</body>
				</html>`))
			return
		},
	)
}

func renderPage() http.Handler {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			http.ServeFile(w, r, "./render.html")
		},
	)
}
