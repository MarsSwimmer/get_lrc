/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"qinglin.org/get_lrc/biz/get_lrc"
	"qinglin.org/get_lrc/cmd/flag"
)

// lxCmd represents the lx command
var lxCmd = &cobra.Command{
	Use:   "lx",
	Args:  cobra.MatchAll(cobra.ExactArgs(0), cobra.OnlyValidArgs),
	Short: "Get lx's lrc",
	Long: `Get lx's lrc. For example:

get_lrc lx [options].`,
	Run: func(cmd *cobra.Command, args []string) {
		lxCmdRun(flag.PlayerUrl)
	},
}

func init() {
	rootCmd.AddCommand(lxCmd)
	lxCmd.Flags().StringVarP(&flag.PlayerUrl, "playerUrl", "", "http://127.0.0.1:23330/status", "specific the playerUrl of lx music")
	lxCmd.Flags().BoolVarP(&flag.ShowErrMsg, "showErrMsg", "", false, "show error message when internal error")
	lxCmd.Flags().StringVarP(&flag.DefaultValue, "defaultValue", "d", "", "specific the default value when internal error and not show error message")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// lxCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// lxCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func lxCmdRun(playerUrl string) bool{
	service, err := get_lrc.GetLxService(playerUrl)
	if err != nil {
		if flag.ShowErrMsg {
			fmt.Println(err.Error())
		} else {
			fmt.Println(flag.DefaultValue)
		}

		return false
	}

	// showSentenceCount无效，落雪API只返回当前一行歌词
	result, err := service.GetProgressLrc(0)
	if err != nil {
		if flag.ShowErrMsg {
			fmt.Println(err.Error())
		} else {
			fmt.Println(flag.DefaultValue)
		}

		return true
	}

	fmt.Println(result)
	return true
}