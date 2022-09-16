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
		ClientId:     "your client id",
		ClientSecret: "your client secret",
		EndPoint:     "https://gw-api.pinduoduo.com/api/router",
		RetryTimes:   3, // 设置接口调用失败重试次数
	})
	pdd := p.GoodAPI()
	resp, _ := pdd.GoodsImageUpload(imagePath)
	m.Store(imagePath, resp)
	<-ch
}
