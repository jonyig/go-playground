package cmd

import (
	"github.com/spf13/cobra"
	"go-playground/service/receiver_interface"
	"log"
)

var receiverInterfaceCmd = &cobra.Command{
	Use:   "receiverInterface",
	Short: "A brief description of your command",
	Run: func(cmd *cobra.Command, args []string) {

		sVals := map[int]receiver_interface.S{1: {"A"}}

		// 你只能通过值调用 Read
		sVals[1].Read()

		// 这不能编译通过：
		//sVals[1].Write("test")

		sPtrs := map[int]*receiver_interface.S{1: {"A"}}

		// 通过指针既可以调用 Read，也可以调用 Write 方法
		sPtrs[1].Read()
		sPtrs[1].Write("test")

		s := receiver_interface.S{Data: "A"}
		// 雖可以成功呼叫，但無法成功寫入
		s.Write("123")
		log.Print(s.Read())

		s1Val := receiver_interface.S1{}
		s1Ptr := &receiver_interface.S1{}
		s2Val := receiver_interface.S2{}
		s2Ptr := &receiver_interface.S2{}

		var i receiver_interface.F
		i = s1Val
		i = s1Ptr
		i = s2Ptr
		// s2Val 為 value struct 無法對應到 指針方法的實作
		//i = s2Val

		log.Print(s2Val)
		log.Print(i)
	},
}

func init() {
	rootCmd.AddCommand(receiverInterfaceCmd)
}
