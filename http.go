package pio

import (
	"bytes"
	"errors"
	"io/ioutil"
	"net/http"
)

func requestPIO(URL string, JSON []byte) ([]byte, error) {
	req, err := http.NewRequest("POST", URL, bytes.NewBuffer(JSON))
	if err != nil {
		return nil, errors.New("make http req error")
	}
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode == 400 {
		return nil, errors.New("invalid access key")

	} else if resp.StatusCode == 401 {
		return nil, errors.New("invalid format, content error")

	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}
