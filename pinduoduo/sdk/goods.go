package sdk

import (
	"encoding/json"
)

type GoodsAPI struct {
	Context *Context
}

//func NewGoodsAPI(cfg *Config) *GoodsAPI {
//	return &GoodsAPI{Context: NewContext(cfg)}
//}

func newGoodsAPIWithContext(c *Context) *GoodsAPI {
	return &GoodsAPI{Context: c}
}

//type GoodsPropertiesList struct {
//	GoodsProperties []GoodsProperties `json:"goods_properties"` //商品属性列表
//}
type GoodsProperties struct {
	RefPid int64  `json:"ref_pid"` // 引用属性id
	Vid    int64  `json:"vid"`     //属性值id
	Value  string `json:"value"`   //属性值
	SpecId int64  `json:"spec_id"` //
}

type GoodsAddResponse struct {
	GoodsCommitId int `json:"goods_commit_id"`
	GoodsId       int `json:"goods_id"`
	MatchedSpuId  int `json:"matched_spu_id"`
}

func (g *GoodsAPI) GoodsAdd(mustParams ...Params) (resp GoodsAddResponse, err error) {
	params := NewParamsWithType("pdd.goods.add", mustParams...)
	r, err := Call(g.Context, params)
	if err != nil {
		return
	}
	bytes, err := GetResponseBytes(r, "goods_add_response")
	err = json.Unmarshal(bytes, &resp)
	if err != nil {
		return GoodsAddResponse{}, err
	}
	return
}
