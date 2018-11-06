package dsl

import (
	"goa.design/goa/design"
	"goa.design/goa/eval"
)

// Tracing defines the tracing configuration for the service.
//
// Tracing is optional in Service definition.
//
// Tracing takes 1 argument: the host/endpoint to send the traces to.
//
// Example:
//
//    Tracing("localhost:5775")
//
func Tracing(endpoint string) {
	s, ok := eval.Current().(*design.ServiceExpr)
	if !ok {
		eval.IncompatibleDSL()
		return
	}
	ep := &design.TracingExpr{Endpoint: endpoint, Service: s}
	s.Tracing = ep
}
