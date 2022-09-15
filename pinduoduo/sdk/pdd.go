package sdk

type Pdd struct {
	Context *Context
}

type Config struct {
	ClientId     string
	ClientSecret string
	RetryTimes   int
	Debug        bool
}

func NewPdd(c *Config) *Pdd {
	return &Pdd{
		Context: NewContext(c),
	}
}

func (p *Pdd) GetGoodsAPI() *GoodsAPI {
	return newGoodsAPIWithContext(p.Context)
}

func (p *Pdd) GetTokenAPI() *TokenAPI {
	return newTokenAPIWithContext(p.Context)
}

func (p *Pdd) GoodAPI() *GoodsAPI {
	return newGoodsAPIWithContext(p.Context)
}

func (p *Pdd) GoodsAuthorizationCat() *GoodsAPI {
	return newGoodsAPIWithContext(p.Context)
}

func (p *Pdd) GoodsImageUpload() *GoodsAPI {
	return newGoodsAPIWithContext(p.Context)
}
