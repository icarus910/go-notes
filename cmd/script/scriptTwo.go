package script

import (
	"fmt"
	"github.com/spf13/cobra"
	"strings"
)

var TwoCmd = &cobra.Command{
	Use:   "two",
	Short: "脚本二短提示",
	Long:  "我是脚本二的常提示",
	Run: func(cmd *cobra.Command, args []string) {
		twoFuncCmd(args)
	},
}

func twoFuncCmd(args []string)  {
	fmt.Println("脚本二的参数: " + strings.Join(args, " "))
}