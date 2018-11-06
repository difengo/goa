package design

import (
	"fmt"
)

type (
	// TracingExpr defines tracing configuration.
	TracingExpr struct {
		// Endpoint used to send the traces to.
		Endpoint string
	}
)

// EvalName returns the generic expression name used in error messages.
func (m *TracingExpr) EvalName() string {
	var prefix, suffix string
	if m.Endpoint != "" {
		suffix = fmt.Sprintf("tracing %#v", m.Endpoint)
	} else {
		suffix = "undefined tracing endpoint"
	}
	prefix = "server "
	return prefix + suffix
}
