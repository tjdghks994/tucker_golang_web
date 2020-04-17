package main

import (
	"log"
	"net/http"
	"time"

	"github.com/tjdghks994/decorator/deco"
	"github.com/tjdghks994/tucker_golang_web/myapp"
)

//logger Decorator func
func logger(w http.ResponseWriter, r *http.Request, h http.Handler) {
	start := time.Now()

	log.Println("[Logger1] Started")
	h.ServeHTTP(w, r)
	log.Println("[Logger1] Completed time : ", time.Since(start).Milliseconds())
}

func logger2(w http.ResponseWriter, r *http.Request, h http.Handler) {
	start := time.Now()

	log.Println("[Logger2] Started")
	h.ServeHTTP(w, r)
	log.Println("[Logger2] Completed time : ", time.Since(start).Milliseconds())
}

func main() {

	mux := myapp.NewHandler()
	decoHandler := deco.NewDecoHandler(mux, logger)
	decoHandler = deco.NewDecoHandler(decoHandler, logger2)

	http.ListenAndServe(":3000", mux)
	http.ListenAndServe(":3001", decoHandler)
}
