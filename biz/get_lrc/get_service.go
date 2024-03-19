package get_lrc

import "qinglin.org/get_lrc/biz/get_lrc/impl"

func GetYesplaymusicService(playerUrl string, currentTracUrl string, splitStr string, prefixStr string, delay int) (GetLrcService, error) {
	service := impl.NewGetLrcYesplaymusic(playerUrl, currentTracUrl, splitStr, prefixStr, delay)

	return service, service.Init()
}
