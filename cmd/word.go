package cmd

import (
	"github.com/spf13/cobra"
)

var wordCmd = &cobra.Command{
	Use:   "word",
	Short: "单词格式转换",
	Long:  "支持多种格式",
	Run: func(cmd *cobra.Command, args []string) {

	},
}

func init() {}
