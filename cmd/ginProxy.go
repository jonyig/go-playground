/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go-playground/service/gin-proxy/handler"

	"github.com/spf13/cobra"
)

// ginProxyCmd represents the ginProxy command
var ginProxyCmd = &cobra.Command{
	Use:   "ginProxy",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("ginProxy called")
		r := gin.Default()
		r.Any("/*proxyPath", handler.ProxyHandler)

		r.Run(":8080")
	},
}

func init() {
	rootCmd.AddCommand(ginProxyCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// ginProxyCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// ginProxyCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}