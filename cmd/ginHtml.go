/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/cobra"
	"go-playground/service/gin-html/handler"
	"net/http"
)

// ginHtmlCmd represents the ginHtml command
var ginHtmlCmd = &cobra.Command{
	Use:   "ginHtml",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		r := gin.Default()
		r.LoadHTMLGlob("service/gin-html/template/*")
		r.HTMLRender = handler.CreateMyRender()
		r.Static("/assetPath", "service/gin-html/asset")
		r.GET("/ping", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"message": "pong",
			})
		})

		r.GET("/test", handler.Index)
		r.Run()
	},
}

func init() {
	rootCmd.AddCommand(ginHtmlCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// ginHtmlCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// ginHtmlCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
