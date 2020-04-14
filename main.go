package main

import (
	"net/http"

	"github.com/tjdghks994/tucker_golang_web/myapp"
)

func main() {
	http.ListenAndServe(":3000", myapp.NewHandler())
}
