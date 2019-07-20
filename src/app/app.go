package main

import (
	"fmt"
	"log"
	"math"
	"net/http"
)

func greeting(textparam string) string {

	x := 0.001
	for i := 0; i < 10000000; i++ {

		x += math.Sqrt(x)
	}
	return fmt.Sprintf("<b>%s</b>", textparam)
}
func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, greeting("Code.education Rocks!"))
}

func main() {

	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":80", nil))
}
