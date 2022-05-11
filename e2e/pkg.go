package e2e

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

const (
	CONST_AUTH_HEADER = "Authorization"
)

type TestingClient struct {
	httpClient *http.Client
	host       string
}

func NewTestingClient(host string) *TestingClient {
	return &TestingClient{
		&http.Client{},
		host,
	}
}

func (self *TestingClient) GET(path string, dst interface{}, headers map[string]string) error {
	return self.DO("GET", path, nil, dst, headers)
}

func (self *TestingClient) POST(path string, src, dst interface{}, headers map[string]string) error {
	return self.DO("POST", path, src, dst, headers)
}

func (self *TestingClient) PUT(path string, src, dst interface{}, headers map[string]string) error {
	return self.DO("PUT", path, src, dst, headers)
}

func (self *TestingClient) DO(method, path string, src, dst interface{}, headers map[string]string) error {

	var buf *bytes.Buffer
	if src != nil {
		b, err := json.Marshal(src)
		if err != nil {
			panic(err)
		}
		buf = bytes.NewBuffer(b)
	}

	var err error
	var req *http.Request
	if buf == nil {
		req, err = http.NewRequest(
			method,
			fmt.Sprintf("%s%s", self.host, path),
			nil,
		)
	} else {
		req, err = http.NewRequest(
			method,
			fmt.Sprintf("%s%s", self.host, path),
			buf,
		)
	}
	if err != nil {
		panic(err)
	}

	for k, v := range headers {
		req.Header.Add(k, v)
		log.Println(method, path, "USING HEADER", k, v)
	}

	resp, err := self.httpClient.Do(req)
	if err != nil {
		panic(err)
	}

	b, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	println(path + " <<<<< " + string(b))

	if resp.StatusCode != 200 {
		panic(fmt.Sprintf("invalid status code: %v", resp.StatusCode))
	}

	return json.Unmarshal(b, dst)
}
