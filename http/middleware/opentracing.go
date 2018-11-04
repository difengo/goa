package middleware

import (
	"net/http"

	opentracing "github.com/opentracing/opentracing-go"
	ext "github.com/opentracing/opentracing-go/ext"
)

// OpenTracing returns a middleware that traces HTTP requests using the globally defined
// opentracing tracer
func OpenTracing() func(h http.Handler) http.Handler {
	return func(h http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

			wireCtx, _ := opentracing.GlobalTracer().Extract(
				opentracing.HTTPHeaders,
				opentracing.HTTPHeadersCarrier(r.Header))

			serverSpan := opentracing.StartSpan(r.URL.Path, ext.RPCServerOption(wireCtx))
			defer serverSpan.Finish()

			r = r.WithContext(opentracing.ContextWithSpan(r.Context(), serverSpan))

			rw := CaptureResponse(w)
			h.ServeHTTP(rw, r)
		})
	}
}
