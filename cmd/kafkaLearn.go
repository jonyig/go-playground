/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"github.com/confluentinc/confluent-kafka-go/kafka"
	"log"
	"os"

	"github.com/spf13/cobra"
)

// kafkaLearnCmd represents the kafkaLearn command
var kafkaLearnCmd = &cobra.Command{
	Use:   "kafkaLearn",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("kafkaLearn called")

		consumer, err := kafka.NewConsumer(&kafka.ConfigMap{
			"bootstrap.servers":  "127.0.0.1:9092",
			"group.id":           "foo2",
			"enable.auto.commit": "false",
			"auto.offset.reset":  "smallest"})

		if err != nil {
			panic(err)
		}
		topics := []string{"quickstart"}
		err = consumer.SubscribeTopics(topics, nil)
		if err != nil {
			panic(err)
		}
		run := true
		for run == true {
			ev := consumer.Poll(100)
			switch e := ev.(type) {
			case *kafka.Message:
				test := string(ev.(*kafka.Message).Value)
				log.Print(test)
				//par, err := consumer.CommitMessage(ev.(*kafka.Message))
				//log.Print(par)
				//log.Print(err)
			case kafka.Error:
				fmt.Fprintf(os.Stderr, "%% Error: %v\n", e)
				run = false
			default:
				fmt.Printf("Ignored %v\n", e)
			}
		}

		consumer.Close()

	},
}

func init() {
	rootCmd.AddCommand(kafkaLearnCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// kafkaLearnCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// kafkaLearnCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
