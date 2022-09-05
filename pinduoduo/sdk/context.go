package sdk

type Context struct {
	ClientId     string
	ClientSecret string
	RetryTimes   int
	Debug        bool
}

func NewContext() *Context {
	return &Context{
		ClientId:     "cfg.ClientId",
		ClientSecret: "cfg.ClientSecret",
		RetryTimes:   3,
		Debug:        true,
	}
}
