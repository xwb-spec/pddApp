package sdk

type Context struct {
	ClientId     string
	ClientSecret string
	EndPoint     string
	RetryTimes   int
	Debug        bool
}

func NewContext(cfg *Config) *Context {
	return &Context{
		ClientId:     cfg.ClientId,
		ClientSecret: cfg.ClientSecret,
		EndPoint:     cfg.EndPoint,
		RetryTimes:   cfg.RetryTimes,
		Debug:        cfg.Debug,
	}
}
