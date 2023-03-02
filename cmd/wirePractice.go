/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/cobra"
	"go-playground/service/wire-practice/di"
	"net/http"
)

// wirePracticeCmd represents the wirePractice command
var wirePracticeCmd = &cobra.Command{
	Use:   "wirePractice",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		r := gin.Default()
		r.GET("/ping", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"message": "pong",
			})
		})
		di.InitTodoHandler(r)
		r.Run()
	},
}

func InitTodoModule(r *gin.Engine) {
}

func init() {
	rootCmd.AddCommand(wirePracticeCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// wirePracticeCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// wirePracticeCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
