/*
Copyright © 2026 Josh Rodriguez josh.rodriguez@arm.com

*/
package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
        "github.com/veraison/cmw"
        "os"
)

// collectionCmd represents the collection command
var collectionCmd = &cobra.Command{
	Use:   "collection",
	Short: "Create a Conceptual Message Wrapper for a (Currently JSON only) collection",
	Long: `Create a Conceptual Message Wrapper for a (Currently JSON only) collection
For example:

cmwcli collection --data collection.json`,
	Run: func(cmd *cobra.Command, args []string) {
		// Read flags
		data, err := cmd.Flags().GetString("data")
		if err != nil || data == "" {
			fmt.Println("No data inputted")
			return
		}

		output, err := cmd.Flags().GetString("output")
		if err != nil || data == "" {
			fmt.Println("No output file specified")
		}

		// Read input
		jsonCollection, err := os.ReadFile(data)
		if err != nil {
			fmt.Println("Error accessing json collection at", data)
			return
		}

		// Create collection
		var cmwCollection cmw.CMW
		cmwCollection.UnmarshalJSON(jsonCollection)

		// Write to output
		fmt.Println(cmwCollection)
		cmwBytes, _ := cmwCollection.MarshalJSON()
		err = os.WriteFile(output, cmwBytes, 0666)
		if err != nil {
			fmt.Println("Error writing to file")
		}
	},
}

func init() {
	rootCmd.AddCommand(collectionCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// collectionCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// collectionCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	collectionCmd.Flags().StringP("data", "d", "", "Location of input data")
	collectionCmd.Flags().StringP("output", "o", "", "Location of output file")
}
