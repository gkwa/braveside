package cmd

import (
	"bytes"
	"os"

	"github.com/gkwa/braveside/core"
	"github.com/spf13/cobra"
)

var helloCmd = &cobra.Command{
	Use:   "hello",
	Short: "A brief description of your command",
	Long:  `A longer description that spans multiple lines and likely contains examples and usage of using your command.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		input, err := os.Open("testdata/input.md")
		if err != nil {
			return err
		}
		defer input.Close()

		var output bytes.Buffer
		var diffOutput bytes.Buffer

		err = core.Hello(cmd.Context(), input, &output, &diffOutput, core.CompareDiff)
		if err != nil {
			return err
		}

		err = os.WriteFile("output.md", output.Bytes(), 0o644)
		if err != nil {
			return err
		}

		_, err = os.Stdout.Write(diffOutput.Bytes())
		return err
	},
}

func init() {
	rootCmd.AddCommand(helloCmd)
}
