package kieclient

import (
	"testing"
)

func newClient() (Client, error) {
	return New(Config{
		Endpoint: testEndpoint,
	})

}

var createkeyId string = ""
var existKeys []string = make([]string, 10)

func TestCreate(t *testing.T) {
	client, err := newClient()
	if err != nil {
		t.Error(err.Error())
		t.FailNow()
	}
	resp, err := client.Create(testProject, &KVRequest{
		Key:   testKey,
		Value: testValue,
		Labels: map[string]string{
			"environment": "onebox",
		},
	})
	if err != nil {
		t.Error(err.Error())
		t.FailNow()
	}
	if resp.Data.Key != testKey && resp.Data.Value != testValue {
		t.Error(err.Error())
		t.FailNow()
	}
	createkeyId = resp.Data.ID
	t.Log(resp.Data)
}

func TestGetWithNotExist(t *testing.T) {
	client, err := newClient()
	if err != nil {
		t.Error(err.Error())
		t.FailNow()
	}
	resp, err := client.Get(testProject, "notexist")
	if err == nil {
		t.Error("get an not exist key, but return success")
		t.FailNow()
	}
	t.Log(resp)
}

func TestGet(t *testing.T) {
	client, err := newClient()
	if err != nil {
		t.Error(err.Error())
		t.FailNow()
	}
	resp, err := client.Get(testProject, createkeyId)
	if err != nil {
		t.Error(err.Error())
		t.FailNow()
	}
	t.Log(resp.Data)
	if resp.Data.Key != testKey && resp.Data.Value != testValue {
		t.Error("get config not right")
		t.FailNow()
	}
}
func TestGetAll(t *testing.T) {
	// TestCreate(t)
	client, err := newClient()
	if err != nil {
		t.Error(err.Error())
		t.FailNow()
	}
	resp, err := client.GetAll(testProject, "", "", -1, map[string]string{
		"environment": "onebox",
	})
	if err != nil {
		t.Error(err.Error())
		t.FailNow()
	}
	if resp.Total <= 0 {
		t.Error("config not found")
		t.FailNow()
	}
	for _, data := range resp.Data {
		t.Log(data.Key, "=", data.Value)
		existKeys = append(existKeys, data.ID)
	}
}

func TestDeleteKeys(t *testing.T) {
	client, err := newClient()
	if err != nil {
		t.Error(err.Error())
		t.FailNow()
	}
	resp, err := client.DeleteKeys(testProject, existKeys)
	if err != nil {
		t.Error(err.Error())
		t.FailNow()
	}
	t.Log(resp.StatusCode)
}
