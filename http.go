package main

import (
	"fmt"
	"freewheel/servicekit/pkg/errors"
	"io"
	"io/ioutil"
	"net/http"
)

var Header = map[string]string{
	"Content-Type": "application/json",
	"User-Agent":   "Freewheel",
}

func main() {
	netClient := http.DefaultClient
	req, _ := CreateRequest("http://localhost:4567/counter?command=get&type=ad&scope=local&index=1&id=500881", "Get", nil)
	resp, err := netClient.Do(req)
	respBody, err := ioutil.ReadAll(resp.Body)
	fmt.Println("respBody: ", string(respBody))
	fmt.Println("err: ", err)
}

func CreateRequest(url string, method string, body io.Reader) (*http.Request, error) {
	var (
		req *http.Request
		err error
	)
	switch method {
	case "Get":
		req, err = http.NewRequest(http.MethodGet, url, nil)
	case "Post":
		req, err = http.NewRequest(http.MethodPost, url, body)
	default:
		return nil, errors.NewInternalServerError("invalid HTTP method")
	}
	if err != nil {
		return nil, err
	}

	for key, value := range Header {
		req.Header.Set(key, value)
	}

	return req, nil
}
