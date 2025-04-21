package examplepackage

import (
	"context"
	"fmt"
)

// ----------------------------------------------------------------------------
// Types
// ----------------------------------------------------------------------------

// BasicExamplePackage is an example type-struct.
type BasicExamplePackage struct {
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
func (examplepackage *BasicExamplePackage) SaySomething(ctx context.Context) error {
	_ = ctx

	outputf("%s: %s\n", exampleConstant, examplepackage.Something)

	return nil
}

// ----------------------------------------------------------------------------
// Private functions
// ----------------------------------------------------------------------------

func outputf(format string, message ...any) {
	fmt.Printf(format, message...) //nolint
}
