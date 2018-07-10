package server

import (
	"encoding/json"
	"log"
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
				panic(err)
				// w.WriteHeader(http.StatusInternalServerError)
				// w.Write([]byte(err.Error()))
				// return
			}

			fp, err := os.Create("pages/render.html")
			if err != nil {
				panic(err)
				// w.WriteHeader(http.StatusInternalServerError)
				// w.Write([]byte(err.Error()))
				// return
			}
			defer fp.Close()
			log.Println(code)
			fp.Write([]byte(`
				<html>
					<head>
						<style>` + (*code).CSS + `</style>
						<script>` + (*code).JS + `</script>
					</head>
					<body>` + (*code).HTML + `</body>
				</html>`))
			return
		},
	)
}

func renderPage() http.Handler {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			http.ServeFile(w, r, "pages/render.html")
		},
	)
}
