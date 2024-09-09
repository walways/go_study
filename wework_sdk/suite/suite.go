package suite

import (
	"wework_sdk/util"
)

// 获取suite——token
// https://developer.work.weixin.qq.com/document/path/97162
const GET_SUITE_TOKEN = "https://qyapi.weixin.qq.com/cgi-bin/service/get_suite_token"

type Suite struct {
	SuiteId            string `json:"suite_id"`
	SuiteSecret        string `json:"suite_secret"`
	DataCallBackUrl    string `json:"data_call_back_url"`    //数据回调地址
	CommandCallBackUrl string `json:"command_call_back_url"` //指令回调地址
	Token              string `json:"token"`
	EncodeAesKey       string `json:"encode_aes_key"`
}

type Token struct {
	Errcode          int    `json:"errcode"`
	Errmsg           string `json:"errmsg"`
	SuiteAccessToken string `json:"suite_access_token"`
	ExpiresIn        int    `json:"expires_in"`
}

func NewSuite(suiteId, suiteSecret, dataCallBackUrl, commandCallBackUrl, token, encodeAesKey string) *Suite {
	return &Suite{
		SuiteId:            suiteId,
		SuiteSecret:        suiteSecret,
		DataCallBackUrl:    dataCallBackUrl,
		CommandCallBackUrl: commandCallBackUrl,
		Token:              token,
		EncodeAesKey:       encodeAesKey,
	}
}

// 获取suiteToken
func (s *Suite) GetSuiteToken(suiteTicket string) (suiteToken *Token, err error) {
	resp, err := util.New().Post(GET_SUITE_TOKEN, map[string]interface{}{
		"suite_ticket": suiteTicket,
		"suite_id":     s.SuiteId,
		"suite_secret": s.SuiteSecret,
	})
	if err != nil {
		return nil, err
	}
	suiteToken = &Token{}
	err = util.Unmarshal(resp.Body, suiteToken)
	return
}
