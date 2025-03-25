package examplepackage

import (
	"context"
	"fmt"
)

// ----------------------------------------------------------------------------
// Examples for godoc documentation
// ----------------------------------------------------------------------------

func ExampleExampleImpl_SaySomething() {
	// For more information, visit https://github.com/senzing-garage/template-go/blob/main/examplepackage/examplepackage_test.go
	ctx := context.TODO()
	examplePackage := &ExampleImpl{
		Something: "I'm here",
	}
	err := examplePackage.SaySomething(ctx)
	if err != nil {
		fmt.Print(err)
	}
	// Output:
	// examplePackage: I'm here
}
