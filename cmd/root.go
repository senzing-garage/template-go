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
	"github.com/senzing/senzing-tools/helper"
	"github.com/senzing/senzing-tools/option"
	"github.com/senzing/template-go/examplepackage"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

const (
	defaultConfiguration           string = ""
	defaultEngineConfigurationJson string = ""
	defaultEngineLogLevel          int    = 0
	defaultLogLevel                string = "INFO"
)

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
	viper.AutomaticEnv()
	replacer := strings.NewReplacer("-", "_")
	viper.SetEnvKeyReplacer(replacer)
	viper.SetEnvPrefix(constant.SetEnvPrefix)

	// Ints

	intOptions := map[string]int{
		option.EngineLogLevel: defaultEngineLogLevel,
	}
	for optionKey, optionValue := range intOptions {
		viper.SetDefault(optionKey, optionValue)
		viper.BindPFlag(optionKey, cobraCommand.Flags().Lookup(optionKey))
	}

	// Strings

	stringOptions := map[string]string{
		option.EngineConfigurationJson: defaultEngineConfigurationJson,
		option.LogLevel:                defaultLogLevel,
	}
	for optionKey, optionValue := range stringOptions {
		viper.SetDefault(optionKey, optionValue)
		viper.BindPFlag(optionKey, cobraCommand.Flags().Lookup(optionKey))
	}
}

// RootCmd represents the command.
var RootCmd = &cobra.Command{
	Use:   "template-go",
	Short: "template-go short description",
	Long: `
template-go long description.
	`,
	PreRun: func(cobraCommand *cobra.Command, args []string) {
		loadConfigurationFile(cobraCommand)
		loadOptions(cobraCommand)
		cobraCommand.SetVersionTemplate(constant.VersionTemplate)
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		var err error = nil
		ctx := context.TODO()
		examplePackage := &examplepackage.ExamplePackageImpl{
			Something: "Main says 'Hi!'",
		}
		err = examplePackage.SaySomething(ctx)
		return err
	},
	Version: helper.MakeVersion(githubVersion, githubIteration),
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the RootCmd.
func Execute() {
	err := RootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

// Since init() is always invoked, define command line parameters.
func init() {
	RootCmd.Flags().Int(option.EngineLogLevel, defaultEngineLogLevel, fmt.Sprintf("Log level for Senzing Engine [%s]", envar.EngineLogLevel))
	RootCmd.Flags().String(option.Configuration, defaultConfiguration, fmt.Sprintf("Path to configuration file [%s]", envar.Configuration))
	RootCmd.Flags().String(option.EngineConfigurationJson, defaultEngineConfigurationJson, fmt.Sprintf("JSON string sent to Senzing's init() function [%s]", envar.EngineConfigurationJson))
}
