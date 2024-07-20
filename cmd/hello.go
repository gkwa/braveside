package cmd

import (
	"github.com/gkwa/braveside/core"
	"github.com/spf13/cobra"
)

var showAST bool

var helloCmd = &cobra.Command{
	Use:   "hello",
	Short: "A brief description of your command",
	Long:  `A longer description that spans multiple lines and likely contains examples and usage of using your command.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		ctx := cmd.Context()
		logger := core.LoggerFrom(ctx)
		ctx = core.ContextWithShowAST(ctx, showAST)
		logger.Info("Running hello command")
		return core.Hello(ctx)
	},
}

func init() {
	rootCmd.AddCommand(helloCmd)
	helloCmd.Flags().BoolVar(&showAST, "show-ast", false, "Show AST structure")
}
