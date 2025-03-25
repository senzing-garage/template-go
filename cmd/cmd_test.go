package cmd_test

import (
	"os"
	"testing"

	"github.com/senzing-garage/template-go/cmd"
	"github.com/stretchr/testify/require"
)

// ----------------------------------------------------------------------------
// Test public functions
// ----------------------------------------------------------------------------

func Test_CompletionCmd(test *testing.T) {
	_ = test
	err := cmd.CompletionCmd.Execute()
	require.NoError(test, err)
	err = cmd.CompletionCmd.RunE(cmd.CompletionCmd, []string{})
	require.NoError(test, err)
}

func Test_DocsCmd(test *testing.T) {
	_ = test
	err := cmd.DocsCmd.Execute()
	require.NoError(test, err)
	err = cmd.DocsCmd.RunE(cmd.DocsCmd, []string{})
	require.NoError(test, err)
}

//

func Test_Execute(test *testing.T) {
	_ = test
	os.Args = []string{"command-name", "--help"}

	cmd.Execute()
}

func Test_Execute_completion(test *testing.T) {
	_ = test
	os.Args = []string{"command-name", "completion"}

	cmd.Execute()
}

func Test_Execute_docs(test *testing.T) {
	_ = test
	os.Args = []string{"command-name", "docs"}

	cmd.Execute()
}

func Test_Execute_help(test *testing.T) {
	_ = test
	os.Args = []string{"command-name", "--help"}

	cmd.Execute()
}

func Test_PreRun(test *testing.T) {
	_ = test
	args := []string{"command-name", "--help"}

	cmd.PreRun(cmd.RootCmd, args)
}

func Test_RunE(test *testing.T) {
	test.Setenv("SENZING_TOOLS_AVOID_SERVING", "true")

	err := cmd.RunE(cmd.RootCmd, []string{})
	require.NoError(test, err)
}

func Test_RootCmd(test *testing.T) {
	_ = test
	err := cmd.RootCmd.Execute()
	require.NoError(test, err)
	err = cmd.RootCmd.RunE(cmd.RootCmd, []string{})
	require.NoError(test, err)
}

// ----------------------------------------------------------------------------
// Test private functions
// ----------------------------------------------------------------------------

// func Test_completionAction(test *testing.T) {
// 	var buffer bytes.Buffer
// 	err := completionAction(&buffer)
// 	require.NoError(test, err)
// }

// func Test_docsAction_badDir(test *testing.T) {
// 	var buffer bytes.Buffer
// 	badDir := "/tmp/no/directory/exists"
// 	err := docsAction(&buffer, badDir)
// 	require.Error(test, err)
// }
