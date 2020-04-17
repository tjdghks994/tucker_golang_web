package deco

import "net/http"

type DecoratorFunc func(http.ResponseWriter, *http.Request, http.Handler)

type DecoHandler struct {
	fn DecoratorFunc
	h  http.Handler
}

func (deco *DecoHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	deco.fn(w, r, deco.h)
}

func NewDecoHandler(h http.Handler, fn DecoratorFunc) http.Handler {
	return &DecoHandler{
		fn: fn,
		h:  h,
	}
}
