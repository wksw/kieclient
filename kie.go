package kieclient

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
)

// Create 创建配置
func (c Client) Create(project string, in *KVRequest) (*KVDocResp, error) {
	resp, err := c.do(http.MethodPost, fmt.Sprintf("/%s/kie/kv", project), in)
	if err != nil {
		return &KVDocResp{
			StatusCode: resp.StatusCode,
		}, err
	}
	var out KVDoc
	if err := json.Unmarshal(resp.Data, &out); err != nil {
		return &KVDocResp{
			StatusCode: resp.StatusCode,
		}, fmt.Errorf("create config umarshal resp '%s' fail[%s]", string(resp.Data), err.Error())
	}
	return &KVDocResp{
		StatusCode: resp.StatusCode,
		Data:       &out,
	}, nil
}

// Update 更新配置
func (c Client) Update(project, kvID string, in *KVRequest) (*KVDocResp, error) {
	resp, err := c.do(http.MethodPut, fmt.Sprintf("/%s/kie/kv/%s", project, kvID), in)
	if err != nil {
		return &KVDocResp{
			StatusCode: resp.StatusCode,
		}, err
	}
	var out KVDoc
	if err := json.Unmarshal(resp.Data, &out); err != nil {
		return &KVDocResp{
			StatusCode: resp.StatusCode,
		}, fmt.Errorf("update config umarshal resp '%s' fail[%s]", string(resp.Data), err.Error())
	}
	return &KVDocResp{
		StatusCode: resp.StatusCode,
		Data:       &out,
	}, nil
}

// Delete 删除配置
func (c Client) Delete(project, kvID string) (*KVDocResp, error) {
	resp, err := c.do(http.MethodDelete, fmt.Sprintf("/%s/kie/kv/%s", project, kvID), nil)
	if err != nil {
		return &KVDocResp{
			StatusCode: resp.StatusCode,
		}, err
	}
	return &KVDocResp{
		StatusCode: resp.StatusCode,
	}, nil
}

// DeleteKeys 删除所有配置
func (c Client) DeleteKeys(project string, deleteIds []string) (*KVDocResp, error) {
	resp, err := c.do(http.MethodDelete, fmt.Sprintf("/%s/kie/kv", project), &DeleteReq{IDs: deleteIds})
	if err != nil {
		return &KVDocResp{
			StatusCode: resp.StatusCode,
		}, err
	}
	return &KVDocResp{
		StatusCode: resp.StatusCode,
	}, nil
}

// Get 获取配置详情
func (c Client) Get(project, kvID string) (*KVDocResp, error) {
	resp, err := c.do(http.MethodGet, fmt.Sprintf("/%s/kie/kv/%s", project, kvID), nil)
	if err != nil {
		return &KVDocResp{
			StatusCode: resp.StatusCode,
		}, err
	}
	var out KVDoc
	if err := json.Unmarshal(resp.Data, &out); err != nil {
		return &KVDocResp{
			StatusCode: resp.StatusCode,
		}, fmt.Errorf("get config umarshal resp '%s' fail[%s]", string(resp.Data), err.Error())
	}
	return &KVDocResp{
		StatusCode: resp.StatusCode,
		Data:       &out,
	}, nil
}

// GetAll 获取配置列表
func (c Client) GetAll(project, key, match string, revision int, labels map[string]string) (*KVResponse, error) {
	v := url.Values{}
	for key := range labels {
		v.Add("label", fmt.Sprintf("%s:%s", key, labels[key]))
	}
	v.Add("revision", fmt.Sprintf("%d", revision))
	if match != "" {
		v.Add("match", match)
	}
	if key != "" {
		v.Add("key", key)
	}
	resp, err := c.do(http.MethodGet, fmt.Sprintf("/%s/kie/kv?%s", project, v.Encode()), nil)
	if err != nil {
		return &KVResponse{
			StatusCode: resp.StatusCode,
		}, err
	}
	var out KVResponse
	if err := json.Unmarshal(resp.Data, &out); err != nil {
		return &KVResponse{
			StatusCode: resp.StatusCode,
		}, fmt.Errorf("get all configs umarshal resp '%s' fail[%s]", string(resp.Data), err.Error())
	}
	out.StatusCode = resp.StatusCode
	return &out, nil
}

// GetTrack 获取配置修订记录
func (c Client) GetTrack(project, kvID string) (*KVResponse, error) {
	resp, err := c.do(http.MethodGet, fmt.Sprintf("%s/kie/revision/%s", project, kvID), nil)
	if err != nil {
		return &KVResponse{
			StatusCode: resp.StatusCode,
		}, err
	}
	var out KVResponse
	if err := json.Unmarshal(resp.Data, &out); err != nil {
		return &KVResponse{
			StatusCode: resp.StatusCode,
		}, fmt.Errorf("get config '%s' revision umarshal resp '%s' fail[%s]", kvID, string(resp.Data), err.Error())
	}
	out.StatusCode = resp.StatusCode
	return &out, nil
}
