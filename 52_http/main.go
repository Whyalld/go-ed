package main

import (
	"fmt"
	"math"
	"net/http"
	"strconv"
)

func main() {
	http.HandleFunc("/circle", func(w http.ResponseWriter, r *http.Request) {
		radius, err := strconv.ParseFloat(r.URL.Query().Get("radius"), 64)
		if err != nil {
			http.Error(w, "radius должен быть числом", http.StatusBadRequest)
			return
		}

		if radius < 0 {
			http.Error(w, "radius не может быть отрицательным", http.StatusBadRequest)
			return
		}

		area := math.Pi * radius * radius

		fmt.Fprintf(w, "Площадь окржуности: %.2f", area)
	})

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Ошибка сервера:", err)
	}
}
