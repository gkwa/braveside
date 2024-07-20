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
	Run: func(cmd *cobra.Command, args []string) {
		logger := LoggerFrom(cmd.Context())
		logger.Info("Running hello command")
		core.Hello(logger, showAST)
	},
}

func init() {
	rootCmd.AddCommand(helloCmd)
	helloCmd.Flags().BoolVar(&showAST, "show-ast", false, "Show AST structure")
}
