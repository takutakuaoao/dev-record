/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"github.com/spf13/cobra"
)

// categoriesCmd represents the categories command
var categoriesCmd = &cobra.Command{
	Use:   "categories",
	Short: "Show All Category Info.",
	Long:  "Show All Category Info.",
	Run: func(cmd *cobra.Command, args []string) {
		// categoryRepository := categoryapi[{
		// 	Items []category.Category `json:"items"`
		// }].NewCategoryApi("http://app")
		// categoryJson := categoryRepository.FetchAll()

		// table := tablewriter.NewWriter(os.Stdout)
		// table.SetHeader([]string{"ID", "Name", "Slug"})

		// for _, v := range categoryJson.Data.Items {
		// 	table.Append(v.ToArray())
		// }
		// table.Render()
	},
}

func init() {
	rootCmd.AddCommand(categoriesCmd)
}
