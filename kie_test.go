package kieclient

import (
	"testing"

	"github.com/apache/servicecomb-kie/pkg/model"
)

func newClient() (Client, error) {
	return New(Config{
		Endpoint: testEndpoint,
	})

}

func TestCreate(t *testing.T) {
	client, err := newClient()
	if err != nil {
		t.Error(err.Error())
		t.FailNow()
	}
	resp, err := client.Create(testProject, &model.KVRequest{
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
	if resp.Key != testKey && resp.Value != testValue {
		t.Error(err.Error())
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
	resp, err := client.Get(testProject, "not_exist")
	if err == nil {
		t.Error("get an not exist key, but return exist")
		t.FailNow()
	}
	t.Log(resp)
}

func TestGetAll(t *testing.T) {
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
	}
}
