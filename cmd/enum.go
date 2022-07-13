package cmd

import (
	"fmt"
	"go-playground/service/enum"

	"github.com/spf13/cobra"
)

var enumCmd = &cobra.Command{
	Use:   "enum",
	Short: "A brief description of your command",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(enum.Add)
		fmt.Println(enum.Subtract)
		fmt.Println(enum.Multiply)
	},
}

func init() {
	rootCmd.AddCommand(enumCmd)
}
