package examplepackage_test

import (
	"context"
	"fmt"

	"github.com/senzing-garage/template-go/examplepackage"
)

// ----------------------------------------------------------------------------
// Examples for godoc documentation
// ----------------------------------------------------------------------------

func ExampleBasicExamplePackage_SaySomething() {
	// For more information, visit
	// https://github.com/senzing-garage/template-go/blob/main/examplepackage/examplepackage_test.go
	ctx := context.TODO()
	examplePackage := &examplepackage.BasicExamplePackage{
		Something: "I'm here",
	}

	err := examplePackage.SaySomething(ctx)
	if err != nil {
		fmt.Print(err)
	}
	// Output:
	// examplePackage: I'm here
}
