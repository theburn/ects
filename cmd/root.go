package cmd

import (
	"fmt"
	"github.com/betterde/ects/internal/utils/system"
	"github.com/spf13/cobra"
	"os"
)

var rootCmd = &cobra.Command{
	Use:     "ects",
	Short:   "Elastic Crontab System",
	Long:    "Elastic Crontab System",
	Version: system.Version,
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}