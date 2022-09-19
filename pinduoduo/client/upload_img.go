package client

import (
	"pddApp/pinduoduo/sdk"
	"sync"
)

// 获取token
func UploadImage(imagePath string, wg *sync.WaitGroup, m *sync.Map, ch chan struct{}) {
	wg.Add(1)
	defer wg.Done()
	pdd := NewClient.GoodAPI()
	params := sdk.NewParams()
	params.Set("access_token", AccessToken)
	resp, _ := pdd.GoodsImageUpload(imagePath, params)
	m.Store(imagePath, resp)
	<-ch
}
