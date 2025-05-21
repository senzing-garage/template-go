/*
 */
package cmd

import (
	"context"
	"os"

	"github.com/senzing-garage/go-cmdhelping/cmdhelper"
	"github.com/senzing-garage/go-cmdhelping/option"
	"github.com/senzing-garage/go-cmdhelping/option/optiontype"
	"github.com/senzing-garage/go-helpers/wraperror"
	"github.com/senzing-garage/template-go/examplepackage"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

const (
	Short string = "template-go short description"
	Use   string = "template-go"
	Long  string = `
template-go long description.
    `
)

// ----------------------------------------------------------------------------
// Context variables
// ----------------------------------------------------------------------------

var SomethingToSay = option.ContextVariable{
	Arg:     "something-to-say",
	Default: option.OsLookupEnvString("SENZING_TOOLS_SOMETHING_TO_SAY", "Main says 'Hi!'"),
	Envar:   "SENZING_TOOLS_SOMETHING_TO_SAY",
	Help:    "Just a phrase to say [%s]",
	Type:    optiontype.String,
}

var ContextVariablesForMultiPlatform = []option.ContextVariable{
	option.Configuration,
	option.EngineSettings,
	option.LogLevel,
	SomethingToSay,
}

var ContextVariables = append(ContextVariablesForMultiPlatform, ContextVariablesForOsArch...)

// ----------------------------------------------------------------------------
// Command
// ----------------------------------------------------------------------------

// RootCmd represents the command.
var RootCmd = &cobra.Command{
	Use:     Use,
	Short:   Short,
	Long:    Long,
	PreRun:  PreRun,
	RunE:    RunE,
	Version: Version(),
}

// ----------------------------------------------------------------------------
// Public functions
// ----------------------------------------------------------------------------

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the RootCmd.
func Execute() {
	err := RootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

// Used in construction of cobra.Command.
func PreRun(cobraCommand *cobra.Command, args []string) {
	cmdhelper.PreRun(cobraCommand, args, Use, ContextVariables)
}

// Used in construction of cobra.Command.
func RunE(_ *cobra.Command, _ []string) error {
	ctx := context.Background()
	examplePackage := &examplepackage.BasicExamplePackage{
		Something: viper.GetString(SomethingToSay.Arg),
	}

	err := examplePackage.SaySomething(ctx)
	if err != nil {
		return wraperror.Errorf(err, "examplePackage.SaySomething failed")
	}

	return wraperror.Errorf(err, wraperror.NoMessage)
}

// Used in construction of cobra.Command.
func Version() string {
	return cmdhelper.Version(githubVersion, githubIteration)
}

// ----------------------------------------------------------------------------
// Private functions
// ----------------------------------------------------------------------------

// Since init() is always invoked, define command line parameters.
func init() {
	cmdhelper.Init(RootCmd, ContextVariables)
}
