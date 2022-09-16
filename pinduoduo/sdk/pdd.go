package sdk

type Pdd struct {
	Context *Context
}

type Config struct {
	ClientId     string
	ClientSecret string
	EndPoint     string
	RetryTimes   int
	Debug        bool
}

func NewPdd(c *Config) *Pdd {
	return &Pdd{
		Context: NewContext(c),
	}
}

func (p *Pdd) TokenAPI() *TokenAPI {
	return newTokenAPIWithContext(p.Context)
}

func (p *Pdd) GoodAPI() *GoodsAPI {
	return newGoodsAPIWithContext(p.Context)
}
