package cmd

import (
	"log"

	"github.com/spf13/cobra"
)

var data string

var addBlockCmd = &cobra.Command{
	Use:   "addblock",
	Short: "Add a new block to the blockchain",
	Run: func(cmd *cobra.Command, args []string) {
		// Check if the blockchain server is running
		if host == nil {
			log.Fatal("The blockchain server is not currently running. \n")
		}
	},
}

func init() {
	addBlockCmd.Flags().StringVarP(&data, "data", "d", "", "Data added to the new block")
	addBlockCmd.MarkFlagRequired("data")
	rootCmd.AddCommand(addBlockCmd)
}
