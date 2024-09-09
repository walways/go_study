package provider

import (
	"wework_sdk/util"
)

// https://developer.work.weixin.qq.com/document/path/96237
const GET_PROVIDER_TOKEN_URL = "https://qyapi.weixin.qq.com/cgi-bin/service/get_provider_token"

type Provider struct {
	CorpId         string `json:"corp_id"`         //服务商企微id
	ProviderSecret string `json:"provider_secret"` //服务商密钥
	CallBackUrl    string `json:"call_back_url"`   //回调地址
	Token          string `json:"token"`
	EncodeAesKey   string `json:"encode_aes_key"`
}

type Token struct {
	AccessToken string `json:"access_token"`
	ExpiresIn   int    `json:"expires_in"`
	ErrCode     int    `json:"errcode"`
	ErrMsg      string `json:"errmsg"`
}

func NewProvider(corpId, providerSecret, callbackUrl, token, encodeAesKey string) *Provider {
	return &Provider{
		CorpId:         corpId,
		ProviderSecret: providerSecret,
		CallBackUrl:    callbackUrl,
		Token:          token,
		EncodeAesKey:   encodeAesKey,
	}
}

// 获取服务商的token
func (p *Provider) GetToken() (providerToken *Token, err error) {
	resp, err := util.New().Post(GET_PROVIDER_TOKEN_URL, map[string]interface{}{
		"corp_id":         p.CorpId,
		"provider_secret": p.ProviderSecret,
	})
	if err != nil {
		return nil, err
	}
	providerToken = &Token{}
	err = until.Unmarshal(resp.Body, providerToken)
	return
}
