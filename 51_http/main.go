package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html; charset=utf-8")

		fmt.Fprint(w, `
		<h1>Привет, Интернет!</h1>
		<p>Это обычный текст.</p>
		<p><i>Это курсивный текст.</i></p>
		`)
	})

	// http.HandleFunc("/echo", func(w http.ResponseWriter, r *http.Request) {
	// 	fmt.Fprint(w, r.URL.Query().Get("message"))
	// })

	http.ListenAndServe(":80", nil)
}
