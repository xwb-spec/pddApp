package sdk

import (
	"encoding/json"
)

// 获取当前店铺可发布商品类目id
type Category struct {
	CatId   int    `json:"cat_id"` //
	Leaf    bool   `json:"leaf"`   // 是否为叶子类目
	CatName string `json:"cat_name"`
}

func (g *GoodsAPI) GoodsAuthorizationCatGet(parentCatId int) (resp []*Category, err error) {
	params := NewParamsWithType("pdd.goods.cats.get")
	params.Set("parent_cat_id", parentCatId)
	r, err := Call(g.Context, params)
	if err != nil {
		return
	}
	bytes, err := GetResponseBytes(r, "goods_auth_cats_get_response", "goods_cats_list")
	json.Unmarshal(bytes, &resp)
	return
}
