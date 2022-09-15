package client

import (
	"log"
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
		RetryTimes:   3, // 设置接口调用失败重试次数
	})
	pdd := p.GoodsImageUpload()
	resp, err := pdd.GoodsImageUpload(imagePath)
	if err != nil {
		log.Printf("上传失败%s", imagePath)
	}
	m.Store(imagePath, resp.ImageUrl)
	<-ch
}
