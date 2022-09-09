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

type ShowCondition struct {
	ParentRefPid int   `json:"parent_ref_pid"`
	ParentVids   []int `json:"parent_vids"`
}
type Properties struct {
	CanNote           bool            `json:"can_note"`
	ChooseMaxNum      int             `json:"choose_max_num"`
	InputMaxNum       int             `json:"input_max_num"`
	IsImportant       bool            `json:"is_important"`
	IsSale            bool            `json:"is_sale"`
	IsSku             bool            `json:"is_sku"`
	MaxValue          string          `json:"max_value"`
	MinValue          string          `json:"min_value"`
	Name              string          `json:"name"` // 品牌
	ParentSpecId      int             `json:"parent_spec_id"`
	PropertyValueType int             `json:"property_value_type"`
	RefPid            int             `json:"ref_pid"`
	Required          bool            `json:"required"`
	RequiredRule      string          `json:"required_rule"`
	RequiredRuleType  int             `json:"required_rule_type"`
	ShowCondition     []ShowCondition `json:"show_condition"`
}
type GoodsPropertiesRule struct {
	ChooseAllQualifySpec bool         `json:"choose_all_qualify_spec"`
	InputMaxSpecNum      int          `json:"input_max_spec_num"`
	Properties           []Properties `json:"properties"`
}

// 添加商品接口
type GoodsAddResponse struct {
	GoodsCommitId int `json:"goods_commit_id"`
	GoodsId       int `json:"goods_id"`
	MatchedSpuId  int `json:"matched_spu_id"`
}
type GoodsResultResponse struct {
	GoodsAddResponse GoodsAddResponse `json:"goods_add_response"`
}

func (g *GoodsAPI) GoodsAdd(goodsName, goodsDesc string) (resp GoodsResultResponse, err error) {
	params := NewParamsWithType("pdd.goods.add")
	params.Set("goods_name", goodsName)        // 商品标题
	params.Set("goods_desc", goodsDesc)        // 商品描述
	params.Set("carousel_gallery", []string{}) // 商品主图/轮播图
	params.Set("detail_gallery", []string{})   //商品详情图
	r, err := Call(g.Context, params)
	if err != nil {
		return
	}
	bytes, err := GetResponseBytes(r, "goods_auth_cats_get_response", "goods_cats_list")
	json.Unmarshal(bytes, &resp)
	return
}
