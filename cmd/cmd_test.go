package cmd

import (
	"testing"
)

func testError(err error) {
	if err != nil {
		panic(err)
	}
}

func Test_Cmd(test *testing.T) {
	_ = test
	Execute()
}

func Test_RootCmd(test *testing.T) {
	_ = test
	err := RootCmd.Execute()
	testError(err)
	err = RootCmd.RunE(RootCmd, []string{})
	testError(err)
}

func Test_completionCmd(test *testing.T) {
	_ = test
	err := completionCmd.Execute()
	testError(err)
	err = completionCmd.RunE(completionCmd, []string{})
	testError(err)
}

func Test_docsCmd(test *testing.T) {
	_ = test
	err := docsCmd.Execute()
	testError(err)
	err = docsCmd.RunE(docsCmd, []string{})
	testError(err)
}
