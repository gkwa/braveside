package cmd

import (
	"github.com/gkwa/braveside/core"
	"github.com/spf13/cobra"
)

var helloCmd = &cobra.Command{
	Use:   "hello",
	Short: "A brief description of your command",
	Long:  `A longer description that spans multiple lines and likely contains examples and usage of using your command.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		return core.Hello(cmd.Context())
	},
}

func init() {
	rootCmd.AddCommand(helloCmd)
	helloCmd.Flags().BoolVar(&core.ShowAST, "show-ast", false, "Show AST structure")
}
