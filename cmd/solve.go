package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var solveCmd = &cobra.Command{
	Use:   "solve",
	Short: "Solve a Sudoku given as the input",
	Long: `Solve a Sudoku puzzle given as input. Input can be passed in using
the ` + "`--flag`" + ` option. The output is printed in ` + "`stdout`" + `.

Please note that the input file must be a text file containing nine lines of
nine character. Each known number at the start of the puzzle should be as is,
but the unknown empty places can be represented by either a dot (` + "`.`" +
		`) or a zero (` + "`0`" + `).`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("solve called")

		fileFlag, _ := cmd.Flags().GetString("file")
		fmt.Println("The file was", fileFlag)
	},
}

func init() {
	rootCmd.AddCommand(solveCmd)

	solveCmd.Flags().StringP("file", "f", "", "File containing the Sudoku input")
	solveCmd.MarkFlagFilename("file")
	solveCmd.MarkFlagRequired("file")
}
