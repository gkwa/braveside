package cmd

import (
	"context"

	"github.com/gkwa/braveside/core"
	"github.com/spf13/cobra"
)

var showAST bool

var helloCmd = &cobra.Command{
	Use:   "hello",
	Short: "A brief description of your command",
	Long:  `A longer description that spans multiple lines and likely contains examples and usage of using your command.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		logger := LoggerFrom(cmd.Context())
		logger.Info("Running hello command")
		ctx := context.WithValue(cmd.Context(), core.ShowASTKey, showAST)
		return core.Hello(ctx, logger)
	},
}

func init() {
	rootCmd.AddCommand(helloCmd)
	helloCmd.Flags().BoolVar(&showAST, "show-ast", false, "Show AST structure")
}
