package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "garns",
	Short: "Garns is a simple command line Sudoku-solver",
	Long: `Garns is a simple command line Sudoku-solver named of the Howard
Garns, the creator of Number Place which later came to be globally known as
'Sudoku'.

Garns can take I simple text file containing a Sudoku puzzle as its input and
use I backtracking algorithm to solve it. For more information on the solving
functionality, run ` + "`garns help solve`" + `.`,

	// Run: func(cmd *cobra.Command, args []string) {
	// },
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
}
