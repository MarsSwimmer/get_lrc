/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"os"

	"github.com/spf13/cobra"
	"qinglin.org/get_lrc/cmd/flag"
)

var (
	lxPlayerUrl      string
	yesplayPlayerUrl string
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "get_lrc",
	Short: "Get music player's lrc",
	Long: `Get music player's lrc. For example:

get_lrc [cmd] [options], you can just use "get_lrc" to get current lrc sentence, it will auto choose source from all music player.`,
	Run: func(cmd *cobra.Command, args []string) {
		// 优先调用yesplay，如果接口反问不通，说明yesplay并没有在播放
		exist := yesPlayCmdRun(yesplayPlayerUrl)

		// yespaly没有在播放时，调用lx的api
		if !exist {
			lxCmdRun(lxPlayerUrl)
		}
	},
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.get_lrc.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	rootCmd.Flags().StringVarP(&lxPlayerUrl, "lxPlayerUrl", "", "http://localhost:23330/status", "specific the playerUrl of lx music")
	rootCmd.Flags().StringVarP(&yesplayPlayerUrl, "yesplayPlayerUrl", "", "http://127.0.0.1:27232/player", "specific the playerUrl of yesplaymusic")
	rootCmd.Flags().BoolVarP(&flag.ShowErrMsg, "showErrMsg", "", false, "show error message when internal error")
	rootCmd.Flags().StringVarP(&flag.DefaultValue, "defaultValue", "d", "", "specific the default value when internal error and not show error message")
}
