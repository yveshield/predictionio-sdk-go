package pio

import (
	"errors"

	"github.com/valyala/fasthttp"
)

func requestPIO(URL string, JSON []byte) ([]byte, error) {
	req := &fasthttp.Request{}
	req.SetRequestURI(URL)

	req.SetBody(JSON)

	req.Header.SetContentType("application/json")
	req.Header.SetMethod("POST")

	resp := &fasthttp.Response{}

	client := &fasthttp.Client{}

	err := client.Do(req, resp)
	if err != nil {
		return nil, err
	}

	statusCode := resp.StatusCode()
	body := resp.Body()
	if statusCode >= 200 && statusCode < 226 {
		return body, nil
	}
	return nil, errors.New(string(body))
}
