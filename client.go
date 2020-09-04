package ratesapi_go

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
	"time"
)

const baseURL = "https://api.ratesapi.io/api/"

type httpClient struct {
	client *http.Client
}

func (c httpClient) Get(URL string, params map[string]string, model interface{}) error {
	req, err := http.NewRequest("GET", URL, nil)
	if err != nil {
		return err
	}

	var q url.Values

	if params != nil {
		q = req.URL.Query()
		for k, v := range params {
			q.Add(k, v)
		}
		req.URL.RawQuery = q.Encode()
	}

	res, err := c.client.Do(req)
	if err != nil {
		return err
	}

	b, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return err
	}

	return parseResponse(b, model)
}

func NewRatesClient() *httpClient {
	client := &http.Client{
		Timeout: time.Second * 10,
	}

	return &httpClient{
		client: client,
	}
}

func parseResponse(b []byte, model interface{}) error {
	return json.Unmarshal(b, &model)
}
