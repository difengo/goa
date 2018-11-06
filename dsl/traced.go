package dsl

import (
	"goa.design/goa/design"
	"goa.design/goa/eval"
)

// Traced defines if the associated method is traced.
//
// Traced is optional in Method or Security Scheme definitions.
//
// Traced takes no argument.
//
// Example:
//
//    Traced()
//
func Traced() {

	switch expr := eval.Current().(type) {
	case *design.MethodExpr:
		expr.Traced = true
	case *design.SchemeExpr:
		expr.Traced = true
	default:
		eval.IncompatibleDSL()
	}
}
