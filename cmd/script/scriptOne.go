package script

import (
	"fmt"
	"github.com/spf13/cobra"
	"strings"
)

var OneCmd = &cobra.Command{
	Use:   "one",
	Short: "脚本一短提示",
	Long:  "我是脚本一的常提示",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("脚本一的参数: " + strings.Join(args, " "))
	},
}