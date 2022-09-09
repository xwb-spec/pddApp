package sdk

import (
	"encoding/json"
)

type RuleValues struct {
	SpecId int64  `json:"spec_id"`
	Value  string `json:"value"` // OPPO
	Vid    int64  `json:"vid"`
}
type RuleProperties struct {
	IsSale bool         `json:"is_sale"`
	IsSku  bool         `json:"is_sku"`
	Name   string       `json:"name"`    // 适用品牌, 适用型号
	RefPid int64        `json:"ref_pid"` //
	Values []RuleValues `json:"values"`
}

// 发布规则查询接口
func (g *GoodsAPI) GoodsCatRuleGet(notMustparams ...Params) (resp []*RuleProperties, err error) {
	params := NewParamsWithType("pdd.goods.cat.rule.get")
	params.Set("cat_id", 3045)
	r, err := Call(g.Context, params)
	if err != nil {
		return
	}
	bytes, err := GetResponseBytes(r, "cat_rule_get_response", "goods_properties_rule", "properties")
	json.Unmarshal(bytes, &resp)
	return
}
