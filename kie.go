package kieclient

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"

	"github.com/apache/servicecomb-kie/pkg/model"
)

// Create 创建配置
func (c Client) Create(project string, in *model.KVRequest) (*model.KVDoc, error) {
	resp, err := c.do(http.MethodPost, fmt.Sprintf("/%s/kie/kv", project), in)
	if err != nil {
		return &model.KVDoc{}, err
	}
	var out model.KVDoc
	if err := json.Unmarshal(resp, &out); err != nil {
		return &model.KVDoc{}, fmt.Errorf("create config umarshal resp '%s' fail[%s]", string(resp), err.Error())
	}
	return &out, nil
}

// Update 更新配置
func (c Client) Update(project, kvID string, in *model.KVRequest) (*model.KVDoc, error) {
	resp, err := c.do(http.MethodPut, fmt.Sprintf("/%s/kie/kv/%s", project, kvID), in)
	if err != nil {
		return &model.KVDoc{}, err
	}
	var out model.KVDoc
	if err := json.Unmarshal(resp, &out); err != nil {
		return &model.KVDoc{}, fmt.Errorf("update config umarshal resp '%s' fail[%s]", string(resp), err.Error())
	}
	return &out, nil
}

// Delete 删除配置
func (c Client) Delete(project, kvID string) error {
	_, err := c.do(http.MethodDelete, fmt.Sprintf("/%s/kie/kv/%s", project, kvID), nil)
	if err != nil {
		return err
	}
	return nil
}

// Get 获取配置详情
func (c Client) Get(project, kvID string) (*model.KVDoc, error) {
	resp, err := c.do(http.MethodGet, fmt.Sprintf("/%s/kie/kv/%s", project, kvID), nil)
	if err != nil {
		return &model.KVDoc{}, err
	}
	var out model.KVDoc
	if err := json.Unmarshal(resp, &out); err != nil {
		return &model.KVDoc{}, fmt.Errorf("get config umarshal resp '%s' fail[%s]", string(resp), err.Error())
	}
	return &out, nil
}

// GetAll 获取配置列表
func (c Client) GetAll(project, key, match string, revision int, labels map[string]string) (*model.KVResponse, error) {
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
		return &model.KVResponse{}, err
	}
	var out model.KVResponse
	if err := json.Unmarshal(resp, &out); err != nil {
		return &model.KVResponse{}, fmt.Errorf("get all configs umarshal resp '%s' fail[%s]", string(resp), err.Error())
	}
	return &out, nil
}

// GetTrack 获取配置修订记录
func (c Client) GetTrack(project, kvID string) (*model.KVDoc, error) {
	resp, err := c.do(http.MethodGet, fmt.Sprintf("%s/kie/revision/%s", project, kvID), nil)
	if err != nil {
		return &model.KVDoc{}, err
	}
	var out model.KVDoc
	if err := json.Unmarshal(resp, &out); err != nil {
		return &model.KVDoc{}, fmt.Errorf("get config '%s' revision umarshal resp '%s' fail[%s]", kvID, string(resp), err.Error())
	}
	return &out, nil
}
