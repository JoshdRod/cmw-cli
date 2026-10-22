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
	Short: "Create a Conceptual Message Wrapper for a (CBOR only) collection",
	Long: `Create a Conceptual Message Wrapper for a (CBOR only) collection
For example:

cmwcli collection --data collection.cbor`,
	Run: func(cmd *cobra.Command, args []string) {
		// Read flags
		data, err := cmd.Flags().GetStringSlice("data")
		if err != nil || data == nil {
			fmt.Println("No data inputted")
			return
		}

		output, err := cmd.Flags().GetString("output")
		if err != nil || output == "" {
			fmt.Println("No output file specified")
			return
		}

		cmwc_t, err := cmd.Flags().GetString("cmwc_t")
		if err != nil {
			fmt.Println("Invalid __cmwc_t specified")
			fmt.Println(err)
			return
		}

		// Read input
		var cmwOutput cmw.CMW

		// If multiple input files, create a collection and add each file's data to it
		if len(data) > 1 {
			cmwOutputRef, err := cmw.NewCollection(cmwc_t)
			if err != nil {
				fmt.Println("Error creating collection")
				fmt.Println(err)
				return
			}
			cmwOutput = *cmwOutputRef
		}
		for i, d := range data {
			newCMWItemBytes, err := os.ReadFile(d)
			if err != nil {
				fmt.Println("Error accessing cbor data in file", d)
				fmt.Println(err)
				return
			}

			// Create Object to be added to CMW
			var newCMWItem cmw.CMW
			err = newCMWItem.UnmarshalCBOR(newCMWItemBytes)
			if err != nil {
				fmt.Println("Error unmarshalling input CBOR file", d, "to CMW")
				fmt.Println(err)
				return
			}

			// If only 1 file inputted, make item the whole CMW
			if len(data) == 1 {
				cmwOutput = newCMWItem
				break
			}

			// Else, Add Object to CMW
			err = cmwOutput.AddCollectionItem(fmt.Sprintf("key%f", i), &newCMWItem)
			if err != nil {
				fmt.Println("Error adding item from file", d, "to CMW collection")
				fmt.Println(err)
				return
			}
		}

		// Write to output file
		fmt.Println(cmwOutput)

		cmwBytes, err := cmwOutput.MarshalCBOR()
		if err != nil {
			fmt.Println("Error marshalling CMW collection to a CBOR file. Maybe check the input file is valid CBOR?")
			return
		}

		err = os.WriteFile(output, cmwBytes, 0666)
		if err != nil {
			fmt.Println("Error writing to file")
			return
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
	collectionCmd.Flags().StringSliceP("data", "d", nil, "Location of input data")
	collectionCmd.Flags().StringP("output", "o", "", "Location of output file")
	collectionCmd.Flags().StringP("cmwc_t", "c", "", "The __cmwc_t of the parent collection")
}
