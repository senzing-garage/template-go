package examplepackage

import (
	"context"
)

// ----------------------------------------------------------------------------
// Examples for godoc documentation
// ----------------------------------------------------------------------------

func ExampleExamplePackageImpl_SaySomething() {
	// For more information, visit https://github.com/Senzing/template-go/blob/main/examplepackage/examplepackage_test.go
	ctx := context.TODO()
	examplePackage := &ExamplePackageImpl{
		Something: "I'm here",
	}
	examplePackage.SaySomething(ctx)
	//Output:
	//examplePackage: I'm here
}
