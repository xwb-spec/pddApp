package client

type SkuProperties struct {
	Punit  string `json:"punit"`   // 属性单位
	RefPid int64  `json:"ref_pid"` // 属性id
	Value  string `json:"value"`   // 属性值
	Vid    string `json:"vid"`     //属性值id
}
type SkuRequestParam struct {
	IsOnsale      int             `json:"is_onsale"`      // sku上架状态
	LimitQuantity int64           `json:"limit_quantity"` // sku购买限制，只入参999
	MultiPrice    int64           `json:"multi_price"`    // 商品团购价格
	SpecIdList    string          `json:"spec_id_list"`   // 商品规格列表，根据pdd.goods.spec.id.get生成的规格属性id
	ThumbUrl      string          `json:"thumb_url"`      // sku 缩略图
	Weight        string          `json:"weight"`         //重量，单位为g
	SkuProperties []SkuProperties `json:"sku_properties"` // sku属性
}
type GoodsPropertiesRequestParam struct {
	GroupId int    `json:"group_id"` // 组id
	ImgUrl  string `json:"img_url"`  // 图片url

}

// 商品添加参数
type GoodsAddRequestParam struct {
	//CommonRequestParam
	CarouselGallery     []string                      `json:"carousel_gallery"`      // 商品轮播图，按次序上传
	CatId               int64                         `json:"cat_id"`                // 叶子类目ID
	CostTemplateId      int64                         `json:"cost_template_id"`      // 物流运费模板ID，可使用pdd.goods.logistics.template.get获取
	CountryId           int                           `json:"country_id"`            // 地区/国家ID，country_id可以通过pdd.goods.country.get获取
	DetailGallery       []string                      `json:"detail_gallery"`        // 商品详情图
	GoodsDesc           string                        `json:"goods_desc"`            // 商品描述
	GoodsName           string                        `json:"goods_name"`            // 商品标题，例如，新疆特产 红满疆枣夹核桃500g
	GoodsType           int                           `json:"goods_type"`            // 商品类型
	IsFolt              bool                          `json:"is_folt"`               // 是否支持假一赔十
	IsPreSale           bool                          `json:"is_pre_sale"`           // 是否预售
	IsRefundable        bool                          `json:"is_refundable"`         // 是否7天无理由退换货
	MarketPrice         int64                         `json:"market_price"`          // 参考价格，单位为分
	SecondHand          bool                          `json:"second_hand"`           // 是否二手商品
	ShipmentLimitSecond int64                         `json:"shipment_limit_second"` // 承诺发货时间（秒
	GoodsProperties     []GoodsPropertiesRequestParam `json:"goods_properties"`      // 商品属性
	SkuList             []SkuRequestParam             `json:"sku_list"`              // sku对象列表
}
