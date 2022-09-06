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

// 获取商品运费模板
type GoodsLogisticsTemplateListResponse struct {
	LogisticsTemplateList []*LogisticsTemplate `json:"logistics_template_list"`
	TotalCount            int                  `json:"total_count"`
}

type LogisticsTemplate struct {
	CostType        int   `json:"cost_type"`
	LastUpdatedTime int64 `json:"last_updated_time"`
	TemplateId      int64 `json:"template_id"`
	TemplateName    int   `json:"template_name"`
}

func (p *Pdd) LogisticsTemplateListGet(page, pageSize int) (res *GoodsLogisticsTemplateListResponse, err error) {
	params := NewParamsWithType("pdd.goods.logistics.template.get")
	params.Set("page", page)
	params.Set("page_size", pageSize)

	r, err := Call(p.Context, params)
	if err != nil {
		return
	}
	bytes, err := GetResponseBytes(r, "goods_logistics_template_get_response")
	res = new(GoodsLogisticsTemplateListResponse)
	err = json.Unmarshal(bytes, res)
	return
}

// 获取商品分类
type Category struct {
	Level       int    `json:"level"`  // 层级，1-一级，2-二级，3-三级，4-四级
	CatId       int    `json:"cat_id"` //
	ParentCatId int    `json:"parent_cat_id"`
	CatName     string `json:"cat_name"`
}

func (g *GoodsAPI) GoodsCatGet(parentCatId int) (res []*Category, err error) {
	params := NewParamsWithType("pdd.goods.cats.get")
	params.Set("parent_cat_id", parentCatId)

	r, err := Call(g.Context, params)
	if err != nil {
		return
	}
	bytes, err := GetResponseBytes(r, "goods_cats_get_response", "goods_cats_list")
	json.Unmarshal(bytes, &res)
	return
}

// 图片上传接口
type ImgUploadResponse struct {
	GoodsImgUploadResponse GoodsImgUploadResponse `json:"goods_img_upload_response"`
}

type GoodsImgUploadResponse struct {
	url string
}

func (g *GoodsAPI) GoodsImgUpload() {
	params := NewParamsWithType("pdd.goods.img.upload")
	params.Set("parent_cat_id", "")

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

func (g *GoodsAPI) GoodsAdd() {
	params := NewParamsWithType("pdd.goods.add")
	params.Set("parent_cat_id", "")

}
