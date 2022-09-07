package sdk

import (
	"encoding/base64"
	"encoding/json"
	"io/ioutil"
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

// 获取当前店铺可发布商品类目id
type Category struct {
	CatId   int    `json:"cat_id"` //
	Leaf    bool   `json:"leaf"`   // 是否为叶子类目
	CatName string `json:"cat_name"`
}

func (g *GoodsAPI) GoodsAuthorizationCatGet(parentCatId int) (res []*Category, err error) {
	params := NewParamsWithType("pdd.goods.cats.get")
	params.Set("parent_cat_id", parentCatId)
	r, err := Call(g.Context, params)
	if err != nil {
		return
	}
	bytes, err := GetResponseBytes(r, "goods_auth_cats_get_response", "goods_cats_list")
	json.Unmarshal(bytes, &res)
	return
}

//
//var a []int
//
//func (g *GoodsAPI) find(res []*Category) []*Category {
//	for _, r := range res {
//		if r.Leaf {
//			a = append(a, r.CatId)
//			break
//		} else {
//			return g.find(res)
//		}
//	}
//}
//func (g *GoodsAPI) GoodsAuthorizationCatGetId() (id int, err error) {
//	i := 0 // 父级
//
//	res, err := g.GoodsAuthorizationCatGet(i)
//	for _, r := range res {
//		if r.Leaf {
//			a = append(a, r.CatId)
//			break
//		} else {
//			res, err := g.GoodsAuthorizationCatGet(r.CatId)
//
//		}
//	}
//}

// End 获取商品分类
type CatRuleGetResponse struct {
	GoodsPropertiesRule GoodsPropertiesRule `json:"goods_properties_rule"`
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

/* 图片上传接口,
支持格式有：jpg/jpeg、png等图片格式，注意入参图片必须转码为base64编码
*/
type GoodsImageUploadResponse struct {
	ImageUploadResponse ImageUploadResponse `json:"goods_image_upload_response"`
}

type ImageUploadResponse struct {
	url string
}

func (g *GoodsAPI) GoodsImageUpload(imagePath string) (res GoodsImageUploadResponse, err error) {
	srcByte, err := ioutil.ReadFile(imagePath)
	if err != nil {
		log.Fatal(err)
	}
	base64Image := base64.StdEncoding.EncodeToString(srcByte)
	params := NewParamsWithType("pdd.goods.image.upload")
	params.Set("image", base64Image)
	r, err := Call(g.Context, params)
	if err != nil {
		return
	}
	bytes, err := GetResponseBytes(r, "goods_cats_get_response", "goods_cats_list")
	json.Unmarshal(bytes, &res)
	return
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

func (g *GoodsAPI) ImageUpload(imagePathList []string) {
	// 上传图片
	var respImageUrl []string
	for _, image := range imagePathList {
		res, err := g.GoodsImageUpload(image)
		if err != nil {
			log.Printf("上传图片%s失败", image)
		}
		respImageUrl = append(respImageUrl, res.ImageUploadResponse.url)
	}
}
func (g *GoodsAPI) GoodsAdd(goodsName, goodsDesc string) {
	params := NewParamsWithType("pdd.goods.add")
	params.Set("goods_name", goodsName)        // 商品标题
	params.Set("goods_desc", goodsDesc)        // 商品描述
	params.Set("carousel_gallery", []string{}) // 商品主图/轮播图
	params.Set("detail_gallery", []string{})   //商品详情图
}
