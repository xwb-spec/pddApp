package sdk

import (
	"encoding/json"
	"fmt"
	"github.com/parnurzeal/gorequest"
	"log"
	"sync"
)

var clients *sync.Pool

func init() {
	clients = &sync.Pool{
		New: func() interface{} {
			return gorequest.New()
		},
	}
}

func Post(context *Context, query string) (b []byte, err error) {
	b, err = post(context, query)
	if err != nil {
		times := 0
		for times < context.RetryTimes {
			b, err = post(context, query)
			if err != nil {
				log.Printf("第 %d 次重试失败：%s", times+1, err)
			} else {
				return
			}
			times++
		}
	}
	return
}

func post(context *Context, query string) (b []byte, err error) {
	client := clients.Get().(*gorequest.SuperAgent)
	_, b, errors := client.Post(context.EndPoint).
		Type("json").
		Send(query).
		EndBytes()
	// push back
	clients.Put(client)
	if IsBadPddRequest(b) {
		err = new(Error)
		json.Unmarshal(b, err)
		return
	}
	if err = getErrorsError(errors); err != nil {
		return
	}
	return
}

func getErrorsError(errors []error) (err error) {
	if len(errors) == 0 {
		return nil
	}
	for _, e := range errors {
		if e != nil {
			err = fmt.Errorf("%s | %s", err, e)
		}
	}
	return
}

func IsBadPddRequest(body []byte) bool {
	return len(body) >= 16 && string(body[2:16]) == "error_response"
}
