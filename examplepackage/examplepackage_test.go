package examplepackage

import (
	"context"
	"fmt"
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

// ----------------------------------------------------------------------------
// Test harness
// ----------------------------------------------------------------------------

func TestMain(m *testing.M) {
	err := setup()
	if err != nil {
		fmt.Print(err)
		os.Exit(1)
	}
	code := m.Run()
	err = teardown()
	if err != nil {
		fmt.Print(err)
	}
	os.Exit(code)
}

func setup() error {
	var err error
	return err
}

func teardown() error {
	var err error
	return err
}

// ----------------------------------------------------------------------------
// Test interface functions
// ----------------------------------------------------------------------------

func TestExamplePackageImpl_SaySomething(test *testing.T) {
	ctx := context.TODO()
	testObject := &ExampleImpl{
		Something: "I'm here",
	}
	err := testObject.SaySomething(ctx)
	require.NoError(test, err)
}
