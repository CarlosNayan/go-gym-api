package utils

import (
	"io"
	"net/http"

	"github.com/valyala/fasthttp"
)

func FiberToHttpHandler(handler fasthttp.RequestHandler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		req := &fasthttp.Request{}
		req.SetRequestURI(r.URL.String())
		req.Header.SetMethod(r.Method)

		for k, vv := range r.Header {
			for _, v := range vv {
				req.Header.Add(k, v)
			}
		}

		if r.Body != nil {
			bodyBytes, _ := io.ReadAll(r.Body)
			req.SetBody(bodyBytes)
		}

		ctx := &fasthttp.RequestCtx{}

		handler(ctx)

		w.WriteHeader(ctx.Response.StatusCode())
		ctx.Response.Header.VisitAll(func(key, value []byte) {
			w.Header().Set(string(key), string(value))
		})

		w.Write(ctx.Response.Body())
	})
}
