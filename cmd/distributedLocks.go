package cmd

import (
	"go-playground/config"
	service "go-playground/service/distributed-locks"
	"log"
	"sync"
	"time"

	"github.com/spf13/cobra"
)

var (
	cfgFile string
)
var distributedLocksCmd = &cobra.Command{
	Use:   "distributedLocks",
	Short: "A brief description of your command",
	Run: func(cmd *cobra.Command, args []string) {
		config, err := config.NewConfig(cfgFile)
		if err != nil {
			log.Fatal(err)
		}

		rdb := service.NewRedis(config.RedisCluster)
		var wg sync.WaitGroup
		err = rdb.SetInventory(2)
		if err != nil {
			log.Fatal(err)
		}

		for i := 1; i <= 4; i++ {
			wg.Add(1)
			go func() {
				defer wg.Done()
				rdb.Lock()
				time.Sleep(10 * time.Second)
				rdb.SubInventory()
				rdb.UnlockUseLua()
			}()
		}
		wg.Wait()
	},
}

func init() {
	rootCmd.AddCommand(distributedLocksCmd)
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is ./config/config.json)")
}
