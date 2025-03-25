/*
 */
package cmd

import (
	"fmt"
	"io"
	"os"

	"github.com/spf13/cobra"
)

// CompletionCmd represents the completion command.
var CompletionCmd = &cobra.Command{
	Use:   "completion",
	Short: "Generate bash completion for the command",
	Long: `To load completions, run:
source < (template-go completion)

To load completions automaticallon on login, add this line to your .bashrc file:
source < (template-go completion)
`,
	RunE: func(cmd *cobra.Command, args []string) error {
		_ = cmd
		_ = args

		return completionAction(os.Stdout)
	},
}

func init() {
	RootCmd.AddCommand(CompletionCmd)
}

func completionAction(out io.Writer) error {
	err := RootCmd.GenBashCompletion(out)
	if err != nil {
		return fmt.Errorf("completionAction: %w", err)
	}

	return nil
}
