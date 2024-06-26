/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"qinglin.org/get_lrc/cmd/flag"
	"qinglin.org/get_lrc/biz/get_lrc"
)

var (
	currentTrackUrl string // 获取正在播放音乐的lrc歌词
	splitStr        string // 多行歌词拼接符号
	prefixStr       string // 展示歌词的前缀
	delay           int    // 延迟毫秒

	sentenceCount int    // 实时歌词展示的行数
	showAll       bool   // 展示所有歌词
)

// yesplayCmd represents the yesplay command
var yesplayCmd = &cobra.Command{
	Use:   "yesplay",
	Args:  cobra.MatchAll(cobra.ExactArgs(0), cobra.OnlyValidArgs),
	Short: "Get yesplaymusic's lrc",
	Long: `Get yesplaymusic's lrc. For example:

get_lrc yesplay [options].`,
	Run: func(cmd *cobra.Command, args []string) {
		yesPlayCmdRun(flag.PlayerUrl)
	},
}

func init() {
	rootCmd.AddCommand(yesplayCmd)
	yesplayCmd.Flags().StringVarP(&flag.PlayerUrl, "playerUrl", "", "http://127.0.0.1:27232/player", "specific the playerUrl of yesplaymusic")
	yesplayCmd.Flags().StringVarP(&currentTrackUrl, "currentTrackUrl", "", "http://127.0.0.1:10754/lyric?id=", "specific the currentTrackUrl of yesplaymusic")
	yesplayCmd.Flags().StringVarP(&splitStr, "split", "", "|", "specific the split to merge multi lrc's sentence")
	yesplayCmd.Flags().StringVarP(&prefixStr, "prefix", "", "歌词：", "specific the prefix of lrc's sentence")
	yesplayCmd.Flags().IntVarP(&delay, "delay", "", 1000, "specific the progress delay time for current lrc's sentence show, ms unit")
	yesplayCmd.Flags().BoolVarP(&showAll, "showAll", "", false, "show all sentence of lrc")
	yesplayCmd.Flags().BoolVarP(&flag.ShowErrMsg, "showErrMsg", "", false, "show error message when internal error")
	yesplayCmd.Flags().StringVarP(&flag.DefaultValue, "defaultValue", "d", "", "specific the default value when internal error and not show error message")

	yesplayCmd.Flags().IntVarP(&sentenceCount, "sentenceCount", "", 1, "specific show current sentence count")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// yesplayCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// yesplayCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func yesPlayCmdRun(playerUrl string) bool{
	service, err := get_lrc.GetYesplaymusicService(playerUrl, currentTrackUrl, splitStr, prefixStr, delay)
		if err != nil {
			if flag.ShowErrMsg {
				fmt.Println(err.Error())
			} else {
				fmt.Println(flag.DefaultValue)
			}

			return false
		}

		if showAll {
			arr, err := service.GetLrc()
			if err != nil {
				if flag.ShowErrMsg {
					fmt.Println(err.Error())
				} else {
					fmt.Println(flag.DefaultValue)
				}

				return true
			}

			for _, str := range arr {
				fmt.Println(str)
			}
			return true
		}

		result, err := service.GetProgressLrc(sentenceCount)
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