package sdk

import (
	"encoding/json"
)

type CatRuleGetResponse struct {
	GoodsPropertiesRule GoodsPropertiesRule `json:"goods_properties_rule"`
}

// 发布规则查询接口
func (g *GoodsAPI) GoodsCatRuleGet(catId int) (res CatRuleGetResponse, err error) {
	params := NewParamsWithType("pdd.goods.cat.rule.get")
	params.Set("cat_id", catId)
	r, err := Call(g.Context, params)
	if err != nil {
		return
	}
	bytes, err := GetResponseBytes(r, "goods_cats_get_response", "goods_cats_list")
	json.Unmarshal(bytes, &res)
	return
}
