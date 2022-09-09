package sdk

import (
	"encoding/json"
	"log"
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
type GoodsResultResponse struct {
	GoodsAddResponse GoodsAddResponse `json:"goods_add_response"`
}

func (g *GoodsAPI) getGoodsProperties() {
	goodsCatRuleGet, err := g.GoodsCatRuleGet()
	if err != nil {
		log.Println("获取商品属性失败")
	}
	var goodsProperties []*GoodsProperties
	for _, r := range goodsCatRuleGet {
		refPid := r.RefPid
		for _, i := range r.Values {
			goodsProperties = append(goodsProperties, &GoodsProperties{
				RefPid: refPid,
				Vid:    i.Vid,
				Value:  i.Value,
				SpecId: i.SpecId,
			})
		}
	}
}
func (g *GoodsAPI) GoodsAdd(goodsName, goodsDesc string, goodsPropertiesList []*GoodsProperties) (resp GoodsResultResponse, err error) {
	params := NewParamsWithType("pdd.goods.add")
	params.Set("goods_name", goodsName)        // 商品标题
	params.Set("goods_desc", goodsDesc)        // 商品描述
	params.Set("carousel_gallery", []string{}) // 商品主图/轮播图
	params.Set("detail_gallery", []string{})   //商品详情图
	goodsProperties, _ := json.Marshal(&goodsPropertiesList)
	params.Set("goods_properties", goodsProperties)
	r, err := Call(g.Context, params)
	if err != nil {
		return
	}
	bytes, err := GetResponseBytes(r, "goods_auth_cats_get_response", "goods_cats_list")
	json.Unmarshal(bytes, &resp)
	return
}
