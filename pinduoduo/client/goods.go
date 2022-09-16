package client

import (
	"log"
	"pddApp/common"
	"pddApp/pinduoduo/sdk"
)

func GoodsCatRuleGet() {
	p := sdk.NewPdd(&sdk.Config{
		ClientId:     CliId,
		ClientSecret: CliSecret,
		RetryTimes:   3, // 设置接口调用失败重试次数
	})
	pdd := p.GoodAPI()
	params := sdk.NewParams()
	params.Set("generate_short_url", true)
	resp, err := pdd.GoodsCatRuleGet()
	if err != nil {
		log.Println("获取商品发布规则失败")
	}
	pro, _ := common.GetGoodsProperties("", "")
	var rules []*sdk.GoodsProperties
	for _, d := range resp {
		refPid := d.RefPid
		name := d.Name
		val, ok := pro[name]
		if ok {
			for _, v := range d.Values {
				for _, n := range val {
					if v.Value == n {
						rules = append(rules, &sdk.GoodsProperties{
							RefPid: refPid,
							Vid:    v.Vid,
							Value:  v.Value,
							SpecId: v.SpecId,
						})
					}
				}

			}
		}

	}
}
