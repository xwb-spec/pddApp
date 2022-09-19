package sdk

import (
	"encoding/json"
)

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

func (g *GoodsAPI) GoodsLogisticsTemplateListGet(page, pageSize int, mustParams ...Params) (resp *GoodsLogisticsTemplateListResponse, err error) {
	params := NewParamsWithType("pdd.goods.logistics.template.get", mustParams...)
	params.Set("page", page)
	params.Set("page_size", pageSize)
	r, err := Call(g.Context, params)
	if err != nil {
		return
	}
	bytes, err := GetResponseBytes(r, "goods_logistics_template_get_response")
	err = json.Unmarshal(bytes, resp)
	return
}
