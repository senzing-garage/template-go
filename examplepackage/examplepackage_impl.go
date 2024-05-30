package examplepackage

import (
	"context"
	"fmt"
)

// ----------------------------------------------------------------------------
// Types
// ----------------------------------------------------------------------------

// ExampleImpl is an example type-struct.
type ExampleImpl struct {
	Something string
}

// ----------------------------------------------------------------------------
// Constants
// ----------------------------------------------------------------------------

const exampleConstant = "examplePackage"

// ----------------------------------------------------------------------------
// Interface methods
// ----------------------------------------------------------------------------

/*
The SaySomething method simply prints the 'Something' value in the type-struct.

Input
  - ctx: A context to control lifecycle.

Output
  - Nothing is returned, except for an error.  However, something is printed.
    See the example output.
*/
func (examplepackage *ExampleImpl) SaySomething(ctx context.Context) error {
	_ = ctx
	fmt.Printf("%s: %s\n", exampleConstant, examplepackage.Something)
	return nil
}
