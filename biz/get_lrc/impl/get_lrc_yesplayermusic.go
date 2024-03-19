package impl

import (
	"fmt"
	"strings"

	"github.com/thedevsaddam/gojsonq"
	"qinglin.org/get_lrc/biz/util/convert"
	"qinglin.org/get_lrc/biz/util/http_help"
	"qinglin.org/get_lrc/biz/util/lrc_scentence"
)

type GetLrcYesplaymusic struct {
	PlayerUrl       string // 正在播放的音乐
	CurrentTrackUrl string // 获取正在播放音乐的lrc歌词
	SplitStr        string // 多行歌词拼接符号
	PrefixStr       string // 展示歌词的前缀
	Delay           int    // 延迟毫秒

	lrcNotExist bool // lrc是否存在
	timeDotArr  []int
	sentenceArr []string
	progress    int
}

func NewGetLrcYesplaymusic(playerUrl string, currentTrackUrl string, splitStr string, prefixStr string, delay int) *GetLrcYesplaymusic {
	return &GetLrcYesplaymusic{
		PlayerUrl:       playerUrl,
		CurrentTrackUrl: currentTrackUrl,
		SplitStr:        splitStr,
		PrefixStr:       prefixStr,
		Delay:           delay,
	}
}

func (s *GetLrcYesplaymusic) Init() error {

	playerResp, err := http_help.DoHttp(s.PlayerUrl)
	if err != nil {
		return err
	}

	s.progress = convert.ToAny[int](convert.ToAny[float32](gojsonq.New().FromString(playerResp).Find("progress")) * 100)

	id := gojsonq.New().FromString(playerResp).Find("currentTrack.id")
	currentTrackId := convert.ToAny[string](id)
	if currentTrackId == "" {
		return fmt.Errorf("currentTrack.id:[%v] convert to string return empty value", id)
	}

	resp, err := http_help.DoHttp(fmt.Sprintf("%s%s", s.CurrentTrackUrl, currentTrackId))
	if err != nil {
		return err
	}

	lrc := gojsonq.New().FromString(resp).Find("lrc.lyric")
	if lrc == nil {
		s.lrcNotExist = true
		return nil
	}

	lrcText := convert.ToAny[string](lrc)
	if lrcText == "" {
		return fmt.Errorf("lrc.lyric:[%v] convert to string return empty value", lrc)
	}

	lrcArr := strings.Split(lrcText, "\n")
	for _, str := range lrcArr {
		str = strings.TrimSpace(str)

		if isTime, start, end := lrc_scentence.Format(str); isTime {
			dateStr := str[start+1 : end]

			min := convert.ToAny[int](strings.TrimLeft(dateStr[:2], "0"))
			sec := convert.ToAny[int](strings.TrimLeft(dateStr[3:5], "0"))
			mSec := convert.ToAny[int](strings.TrimLeft(dateStr[6:], "0"))
			// 可能是3位
			if mSec > 100 {
				mSec = mSec / 10
			}
			tDot := min*6000 + sec*100 + mSec

			sentence := strings.TrimSpace(str[end+1:])
			if sentence != "" {
				s.timeDotArr = append(s.timeDotArr, tDot)
				s.sentenceArr = append(s.sentenceArr, sentence)
			}
		}
	}

	return nil
}

func (s *GetLrcYesplaymusic) GetLrc() (allSentence []string, err error) {
	return s.sentenceArr, nil
}

func (s *GetLrcYesplaymusic) GetProgressLrc(showSentenceCount int) (progressSentence string, err error) {

	if s.lrcNotExist {
		return "", nil
	}

	stLength := len(s.timeDotArr)
	if stLength == 0 {
		return "", nil
	}

	if stLength == 1 {
		return fmt.Sprintf("%s%s", s.PrefixStr, s.sentenceArr[stLength-1]), nil
	}

	startTime := s.progress
	startTimeAfterD := s.progress + (s.Delay / 10)

	// 获取未开始部分
	index := -1
	indexAfterD := -1
	for i := 0; i < stLength; i++ {
		if gap := s.timeDotArr[i] - startTime; gap >= 0 && index == -1 {
			index = i
			if gap > (s.Delay/10) && i > 0 {
				index--
			}
		}
		if gap := s.timeDotArr[i] - startTimeAfterD; gap > 0 {
			indexAfterD = i
			if gap > (s.Delay/10) && i > 0 {
				indexAfterD--
			}
			break
		}
	}

	// 所有lrc行都播放结束
	if index == -1{
		return fmt.Sprintf("%s%s", s.PrefixStr, s.sentenceArr[stLength-1]), nil
	}

	if indexAfterD == -1{
		indexAfterD = stLength - 1
	}

	indexEnd := indexAfterD
	if indexAfterD-index < 1 {
		indexEnd += showSentenceCount
	}

	res := ""
	for i := index; i < stLength && i < indexEnd; i++ {
		if res == "" {
			res = s.sentenceArr[i]
		} else {
			res = fmt.Sprintf("%s %s %s", res, s.SplitStr, s.sentenceArr[i])
		}
	}

	return fmt.Sprintf("%s%s", s.PrefixStr, res), nil
}
