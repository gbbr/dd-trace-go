// +build go1.7

package muxtrace

import (
	"net/http"

	"github.com/DataDog/dd-trace-go/tracer"
)

// SetRequestSpan sets the span on the request's context. Under the hood,
// it will use request.Context() if it's available, otherwise falling back
// to using gorilla/context.
func SetRequestSpan(r *http.Request, span *tracer.Span) *http.Request {
	if r == nil || span == nil {
		return r
	}

	ctx := tracer.ContextWithSpan(r.Context(), span)
	return r.WithContext(ctx)
}

// GetRequestSpan will return the span associated with the given request. It
// will return nil/false if it doesn't exist.
func GetRequestSpan(r *http.Request) (*tracer.Span, bool) {
	span, ok := tracer.SpanFromContext(r.Context())
	return span, ok
}