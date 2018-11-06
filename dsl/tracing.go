package dsl

import (
	"goa.design/goa/design"
	"goa.design/goa/eval"
)

// Tracing defines the tracing configuration for the server.
//
// Tracing is optional in Server definition.
//
// Tracing takes 1 argument: the host/endpoint to send the traces to.
//
// Example:
//
//    Tracing("localhost:5775")
//
func Tracing(endpoint string) {

	switch expr := eval.Current().(type) {
	case *design.APIExpr:
		expr.Tracing = &design.TracingExpr{Endpoint: endpoint}
	case *design.ServerExpr:
		expr.Tracing = &design.TracingExpr{Endpoint: endpoint}
	default:
		eval.IncompatibleDSL()
	}
}
