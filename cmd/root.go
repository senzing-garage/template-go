/*
 */
package cmd

import (
	"context"
	"fmt"
	"os"
	"strings"

	"github.com/senzing/senzing-tools/constant"
	"github.com/senzing/senzing-tools/envar"
	"github.com/senzing/senzing-tools/help"
	"github.com/senzing/senzing-tools/helper"
	"github.com/senzing/senzing-tools/option"
	"github.com/senzing/template-go/examplepackage"
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

var ContextBools = []struct {
	Dfault bool
	Envar  string
	Help   string
	Option string
}{
	{
		Dfault: false,
		Envar:  envar.EnableAll,
		Help:   help.EnableAll,
		Option: option.EnableAll,
	},
}

var ContextInts = []struct {
	Dfault int
	Envar  string
	Help   string
	Option string
}{
	{
		Dfault: 0,
		Envar:  envar.EngineLogLevel,
		Help:   help.EngineLogLevel,
		Option: option.EngineLogLevel,
	},
}

var ContextStrings = []struct {
	Dfault string
	Envar  string
	Help   string
	Option string
}{
	{
		Dfault: "",
		Envar:  envar.Configuration,
		Help:   help.Configuration,
		Option: option.Configuration,
	},
	{
		Dfault: "",
		Envar:  envar.EngineConfigurationJson,
		Help:   help.EngineConfigurationJson,
		Option: option.EngineConfigurationJson,
	},
	{
		Dfault: "INFO",
		Envar:  envar.LogLevel,
		Help:   help.LogLevel,
		Option: option.LogLevel,
	},
}

var ContextStringSlices = []struct {
	Dfault []string
	Envar  string
	Help   string
	Option string
}{
	{
		Dfault: []string{},
		Envar:  envar.XtermAllowedHostnames,
		Help:   help.XtermAllowedHostnames,
		Option: option.XtermAllowedHostnames,
	},
}

// ----------------------------------------------------------------------------
// Private functions
// ----------------------------------------------------------------------------

// Since init() is always invoked, define command line parameters.
func init() {
	for _, contextBool := range ContextBools {
		RootCmd.Flags().Bool(contextBool.Option, contextBool.Dfault, fmt.Sprintf(contextBool.Help, contextBool.Envar))
	}
	for _, contextInt := range ContextInts {
		RootCmd.Flags().Int(contextInt.Option, contextInt.Dfault, fmt.Sprintf(contextInt.Help, contextInt.Envar))
	}
	for _, contextString := range ContextStrings {
		RootCmd.Flags().String(contextString.Option, contextString.Dfault, fmt.Sprintf(contextString.Help, contextString.Envar))
	}
	for _, contextStringSlice := range ContextStringSlices {
		RootCmd.Flags().StringSlice(contextStringSlice.Option, contextStringSlice.Dfault, fmt.Sprintf(contextStringSlice.Help, contextStringSlice.Envar))
	}
}

// If a configuration file is present, load it.
func loadConfigurationFile(cobraCommand *cobra.Command) {
	configuration := ""
	configFlag := cobraCommand.Flags().Lookup(option.Configuration)
	if configFlag != nil {
		configuration = configFlag.Value.String()
	}
	if configuration != "" { // Use configuration file specified as a command line option.
		viper.SetConfigFile(configuration)
	} else { // Search for a configuration file.

		// Determine home directory.

		home, err := os.UserHomeDir()
		cobra.CheckErr(err)

		// Specify configuration file name.

		viper.SetConfigName("template-go")
		viper.SetConfigType("yaml")

		// Define search path order.

		viper.AddConfigPath(home + "/.senzing-tools")
		viper.AddConfigPath(home)
		viper.AddConfigPath("/etc/senzing-tools")
	}

	// If a config file is found, read it in.

	if err := viper.ReadInConfig(); err == nil {
		fmt.Fprintln(os.Stderr, "Applying configuration file:", viper.ConfigFileUsed())
	}
}

// Configure Viper with user-specified options.
func loadOptions(cobraCommand *cobra.Command) {
	var err error = nil
	viper.AutomaticEnv()
	replacer := strings.NewReplacer("-", "_")
	viper.SetEnvKeyReplacer(replacer)
	viper.SetEnvPrefix(constant.SetEnvPrefix)

	// Bools

	for _, contextVar := range ContextBools {
		viper.SetDefault(contextVar.Option, contextVar.Dfault)
		err = viper.BindPFlag(contextVar.Option, cobraCommand.Flags().Lookup(contextVar.Option))
		if err != nil {
			panic(err)
		}
	}

	// Ints

	for _, contextVar := range ContextInts {
		viper.SetDefault(contextVar.Option, contextVar.Dfault)
		err = viper.BindPFlag(contextVar.Option, cobraCommand.Flags().Lookup(contextVar.Option))
		if err != nil {
			panic(err)
		}
	}

	// Strings

	for _, contextVar := range ContextStrings {
		viper.SetDefault(contextVar.Option, contextVar.Dfault)
		err = viper.BindPFlag(contextVar.Option, cobraCommand.Flags().Lookup(contextVar.Option))
		if err != nil {
			panic(err)
		}
	}
	// StringSlice

	for _, contextVar := range ContextStringSlices {
		viper.SetDefault(contextVar.Option, contextVar.Dfault)
		err = viper.BindPFlag(contextVar.Option, cobraCommand.Flags().Lookup(contextVar.Option))
		if err != nil {
			panic(err)
		}
	}
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

// Used in construction of cobra.Command
func PreRun(cobraCommand *cobra.Command, args []string) {
	loadConfigurationFile(cobraCommand)
	loadOptions(cobraCommand)
	cobraCommand.SetVersionTemplate(constant.VersionTemplate)
}

// Used in construction of cobra.Command
func RunE(_ *cobra.Command, _ []string) error {
	var err error = nil
	ctx := context.TODO()
	examplePackage := &examplepackage.ExamplePackageImpl{
		Something: "Main says 'Hi!'",
	}
	err = examplePackage.SaySomething(ctx)
	return err
}

// Used in construction of cobra.Command
func Version() string {
	return helper.MakeVersion(githubVersion, githubIteration)
}

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
