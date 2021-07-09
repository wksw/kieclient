package kieclient

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"time"
)

const (
	// API_VERSION 接口版本
	API_VERSION = "v1"
)

// Client 连接servicecomb-kie客户端
type Client struct {
	// 客户端配置
	config Config
	client *http.Client
}

// Config 客户端配置
type Config struct {
	// servicecomb-kie服务地址
	Endpoint string
	// 接口版本
	ApiVersion string
	// 默认label
	DefaultLables map[string]string
}

// 错误返回
type errResp struct {
	ErrorMsg string `json:"error_msg"`
}

// Resp 请求返回
type Resp struct {
	StatusCode int
	Data       []byte
}

// New 创建连接servicecomb-kie客户端
func New(config Config) (Client, error) {
	transport := &http.Transport{}

	u, err := url.Parse(config.Endpoint)
	if err != nil {
		return Client{}, fmt.Errorf("parse endpoint fail[%s]", err.Error())
	}
	if u.Scheme == "https" {
		transport.TLSClientConfig = &tls.Config{
			InsecureSkipVerify: true,
		}
	}
	config.Endpoint = u.String() + "/"
	client := &http.Client{
		Timeout:   5 * time.Second,
		Transport: transport,
	}
	if config.ApiVersion == "" {
		config.ApiVersion = API_VERSION
	}
	if config.DefaultLables == nil {
		config.DefaultLables = make(map[string]string)
	}
	return Client{
		client: client,
		config: config,
	}, nil
}

// Do 发送http请求
func (c Client) do(method, path string, body interface{}) (*Resp, error) {
	uri, err := url.ParseRequestURI(c.config.Endpoint + c.config.ApiVersion + path)
	if err != nil {
		return &Resp{
			StatusCode: http.StatusBadRequest,
		}, fmt.Errorf("parse request host and request path fail[%s]", err.Error())
	}
	var requestBody []byte
	if body != nil {
		bodyByte, err := json.Marshal(body)
		if err != nil {
			return &Resp{
				StatusCode: http.StatusInternalServerError,
			}, fmt.Errorf("marshal request body fail[%s]", err.Error())
		}
		requestBody = bodyByte
	}
	req, err := http.NewRequest(method, uri.String(), bytes.NewReader(requestBody))
	if err != nil {
		return &Resp{
			StatusCode: http.StatusInternalServerError,
		}, fmt.Errorf("new request fail[%s]", err.Error())
	}
	req.Header.Add("Content-Type", "application/json")
	resp, err := c.client.Do(req)
	if err != nil {
		return &Resp{
			StatusCode: http.StatusInternalServerError,
		}, fmt.Errorf("request to remote fail[%s]", err.Error())
	}
	defer resp.Body.Close()
	result, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return &Resp{
			StatusCode: http.StatusInternalServerError,
		}, fmt.Errorf("read response body fail[%s]", err.Error())
	}
	if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusNoContent {
		var eResp errResp
		if err := json.Unmarshal(result, &eResp); err != nil {
			return &Resp{
				StatusCode: http.StatusInternalServerError,
			}, fmt.Errorf("unmarshal response into errResp fail[%s]", err.Error())
		}
		err = errors.New(eResp.ErrorMsg)
	}

	if os.Getenv("DEBUG") != "" {
		fmt.Println("--- BEGIN ---")
		fmt.Printf("> %s %s %s\n", req.Method, req.URL.RequestURI(), req.Proto)
		for key, header := range req.Header {
			for _, value := range header {
				fmt.Printf("> %s: %s\n", key, value)
			}
		}
		fmt.Println(">")
		fmt.Println(string(requestBody))
		fmt.Printf("< %s %s\n", resp.Proto, resp.Status)
		for key, header := range resp.Header {
			for _, value := range header {
				fmt.Printf("< %s: %s\n", key, value)
			}
		}

		fmt.Println("< ")
		fmt.Println(string(result))
		fmt.Println("< ")
		fmt.Println("--- END ---")
	}

	return &Resp{
		StatusCode: resp.StatusCode,
		Data:       result,
	}, err
}
