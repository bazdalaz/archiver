package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Short: "File archiving tool",
}

// Execute executes the root command.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		handelError(err)
	}
}


// handelError prints the error and exits the program.
func handelError(err error) {
	_, _ = fmt.Fprintln(os.Stderr, err)
	os.Exit(1)
	
}