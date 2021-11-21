package http

import (
	"io/ioutil"
	"net/http"
	"sync"
	"time"
)

type HttpClient interface {
	GET(url string) ([]byte, error)
}

var once sync.Once
var defaultHttp DefaultHttpClient

type DefaultHttpClient struct {
	Client http.Client
}

func NewDefaultHttpClient() HttpClient {
	once.Do(func() {
		defaultHttp = DefaultHttpClient{Client: http.Client{
			Timeout: 3 * time.Second,
		}}
	})
	return &defaultHttp
}

func (d *DefaultHttpClient) GET(url string) ([]byte, error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	resp, err := d.Client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return data, nil
}
