package client

import (
	"pddApp/pinduoduo/sdk"
	"sync"
)

// 获取token
func UploadImage(imagePath string, wg *sync.WaitGroup, m *sync.Map, ch chan struct{}) {
	wg.Add(1)
	defer wg.Done()
	p := sdk.NewPdd(&sdk.Config{
		ClientId:     CliId,
		ClientSecret: CliSecret,
		EndPoint:     EndPoint,
		RetryTimes:   3, // 设置接口调用失败重试次数
	})
	pdd := p.GoodAPI()
	params := sdk.NewParams()
	params.Set("access_token", AccessToken)
	resp, _ := pdd.GoodsImageUpload(imagePath, params)
	m.Store(imagePath, resp)
	<-ch
}
