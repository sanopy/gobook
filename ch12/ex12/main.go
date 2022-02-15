// Search is a demo of the params.Unpack function.
package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/sanopy/gobook/ch12/ex12/params"
)

// search implements the /search URL endpoint.
func search(resp http.ResponseWriter, req *http.Request) {
	var data struct {
		Email   string `http:"e" validate:"email"`
		Credit  string `http:"c" validate:"credit"`
		Zipcode int    `http:"z" validate:"zipcode"`
	}
	if err := params.Unpack(req, &data); err != nil {
		http.Error(resp, err.Error(), http.StatusBadRequest) // 400
		return
	}

	// ...rest of handler...
	fmt.Fprintf(resp, "Search: %+v\n", data)
}

func main() {
	http.HandleFunc("/search", search)
	log.Fatal(http.ListenAndServe(":12345", nil))
}
