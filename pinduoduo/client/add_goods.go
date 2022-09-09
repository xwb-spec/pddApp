package client

import (
	"fmt"
	"log"
	"pddApp/pinduoduo/sdk"
)

func GoodCat() {
	p := sdk.NewPdd(&sdk.Config{
		ClientId:     "your client id",
		ClientSecret: "your client secret",
		RetryTimes:   3, // 设置接口调用失败重试次数
	})
	pdd := p.GoodAPI()
	params := sdk.NewParams()
	params.Set("custom_parameters", "test")
	params.Set("generate_short_url", true)
	resp, err := pdd.GoodsCatRuleGet(params)
	if err != nil {
		log.Println("获取token失败")
	}
	fmt.Println(resp)
}
