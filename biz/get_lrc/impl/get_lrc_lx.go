package impl

import (
	"github.com/thedevsaddam/gojsonq"
	"qinglin.org/get_lrc/biz/util/convert"
	"qinglin.org/get_lrc/biz/util/http_help"
)

type GetLrcLX struct {
	PlayerUrl        string

	
	progressSentence string // lrc歌词
}

func NewGetLrcLX(plarUrl string) *GetLrcLX {
	return &GetLrcLX{
		PlayerUrl: plarUrl,
	}
}

func (s *GetLrcLX) Init() error {
	playerResp, err := http_help.DoHttp(s.PlayerUrl)
	if err != nil {
		return err
	}
	lyricLineText := gojsonq.New().FromString(playerResp).Find("lyricLineText")
	s.progressSentence = convert.ToAny[string](lyricLineText)
	return nil
}

func (s *GetLrcLX) GetProgressLrc(showSentenceCount int) (progressSentence string, err error) {
	// showSentenceCount无效，落雪API只返回当前一行歌词
	return s.progressSentence, nil
}

func (s *GetLrcLX) GetLrc() (allSentence []string, err error) {
	// 落雪音乐暂未提供api
	return nil, nil
}
