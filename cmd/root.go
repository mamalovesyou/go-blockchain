package cmd

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "gochain",
	Short: "Gochain is a simple cli to use my blockchain from the command line",
	Long:  `Gochain is a simple cli to use my blockchain from the command line`,
}

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}
}

// Execute the commad
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
