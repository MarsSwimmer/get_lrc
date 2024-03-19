package get_lrc

type GetLrcService interface {

	// 获取进度歌词
	//
	// @param
	// -showSentenceCount: 需要展示几行歌词，exp：1（展示当前正在播放的歌词），2（展示当前播放的歌词及后一行歌词）
	//
	// @return
	// -progressSentence: 进度歌词
	// -err: 内部异常，如网络请求中断等
	GetProgressLrc(showSentenceCount int) (progressSentence string, err error)

	// 获取播放歌曲的歌词
	//
	// @param
	//
	// @return
	// -allSentence: 歌词
	// -err: 内部异常，如网络请求中断等
	GetLrc() (allSentence []string, err error)
}
