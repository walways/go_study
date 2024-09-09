package util

import (
	"fmt"
	"testing"
)

func TestRestClient_Get(t *testing.T) {
	client := New()
	res, err := client.Get("https://qyapi.weixin.qq.com/cgi-bin/service/get_provider_token")
	type TestStruct struct {
		Errcode     int    `json:"errcode"`
		Errmsg      string `json:"errmsg"`
		AccessToken string `json:"access_token"`
		ExpiresIn   int    `json:"expires_in"`
	}
	test := &TestStruct{}
	testErr := Unmarshal(res.Body, &test)
	fmt.Printf("res:%+v, err:%+v, struct:%+v , struct-err:%+v", res, err, test, testErr)
}

func TestRestClient_Post(t *testing.T) {
	client := New()
	res, err := client.Post("https://qyapi.weixin.qq.com/cgi-bin/service/get_provider_token", map[string]interface{}{
		"corpid":          "wwb0e6b6e3d0e0e0e0",
		"provider_secret": "8e0e0e0e0e0e0e0e0e0e0e0e0e",
	})
	fmt.Printf("res:%+v, err:%+v", res, err)
}
