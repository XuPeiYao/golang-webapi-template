package app

import (
	"github.com/go-resty/resty/v2"
	"time"
)

func NewHttpClient() *resty.Client {
	return resty.New().SetHeaders(map[string]string{
		"application": "edward-testing",
	}).SetRetryCount(3).SetRetryWaitTime(time.Second * 5)
}
