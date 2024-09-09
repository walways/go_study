package wework_sdk

import (
	"fmt"
	"testing"
)

func GetClient() *CallbackDecode {
	token := "QDG6eK"
	receiverId := "wx5823bf96d3bd56c7"
	encodingAeskey := "jWmYm7qr5nMoAUwZRjGtBxmz3KA1tkAj3ykkR6q2B2C"
	client := New(receiverId, &Config{
		EncodeAesKey: encodingAeskey,
		Token:        token,
	})
	return client
}

func TestCallbackDecode_VerifyURL(t *testing.T) {

	client := GetClient()
	// verifyMsgSign := HttpUtils.ParseUrl("msg_signature")
	verifyMsgSign := "5c45ff5e21c57e6ad56bac8758b79b1d9ac89fd3"
	// verifyTimestamp := HttpUtils.ParseUrl("timestamp")
	verifyTimestamp := "1409659589"
	// verifyNonce := HttpUtils.ParseUrl("nonce")
	verifyNonce := "263014780"
	// verifyEchoStr := HttpUtils.ParseUrl("echoStr")
	verifyEchoStr := "P9nAzCzyDtyTWESHep1vC5X9xho/qYX3Zpb4yKa9SKld1DsH3Iyt3tP3zNdtp+4RPcs8TgAE7OaBO+FZXvnaqQ=="
	str, err := client.VerifyURL(verifyMsgSign, verifyTimestamp, verifyNonce, verifyEchoStr)
	fmt.Println(str, err)
}

func TestCallbackDecode_DecryptMsg(t *testing.T) {
	client := GetClient()
	// reqMsgSign := HttpUtils.ParseUrl("msg_signature")
	reqMsgSign := "0623cbc5a8cbee5bcc137c70de99575366fc2af3"
	// reqTimestamp := HttpUtils.ParseUrl("timestamp")
	reqTimestamp := "1409659813"
	// reqNonce := HttpUtils.ParseUrl("nonce")
	reqNonce := "1372623149"
	type Verify struct {
		ToUsername   string `json:"ToUserName"`
		FromUsername string `json:"FromUserName"`
		CreateTime   uint32 `json:"CreateTime"`
		MsgType      string `json:"MsgType"`
		Content      string `json:"Content"`
		Msgid        uint64 `json:"MsgId"`
		Agentid      uint32 `json:"AgentId"`
	}
	var verify Verify
	//reqData := "{\"tousername\":\"wx5823bf96d3bd56c7\",\"encrypt\":\"CZWs4CWRpI4VolQlvn4dlEC1alN2MUEY2VklGehgBVLBrlVF7SyT+SV+Toj43l4ayJ9UMGKphktKKmP7B2j/P1ey67XB8PBgS7Wr5/8+w/yWriZv3Vmoo/MH3/1HsIWZrPQ3N2mJrelStIfI2Y8kLKXA7EhfZgZX4o+ffdkZDM76SEl79Ib9mw7TGjZ9Aw/x/A2VjNbV1E8BtEbRxYYcQippYNw7hr8sFfa3nW1xLdxokt8QkRX83vK3DFP2F6TQFPL2Tu98UwhcUpPvdJBuu1/yiOQIScppV3eOuLWEsko=\",\"agentid\":\"218\"}"
	reqData := `{"tousername":"wx5823bf96d3bd56c7","encrypt":"CZWs4CWRpI4VolQlvn4dlEC1alN2MUEY2VklGehgBVLBrlVF7SyT+SV+Toj43l4ayJ9UMGKphktKKmP7B2j/P1ey67XB8PBgS7Wr5/8+w/yWriZv3Vmoo/MH3/1HsIWZrPQ3N2mJrelStIfI2Y8kLKXA7EhfZgZX4o+ffdkZDM76SEl79Ib9mw7TGjZ9Aw/x/A2VjNbV1E8BtEbRxYYcQippYNw7hr8sFfa3nW1xLdxokt8QkRX83vK3DFP2F6TQFPL2Tu98UwhcUpPvdJBuu1/yiOQIScppV3eOuLWEsko=","agentid":"218"}`
	err := client.DecryptMsg(reqMsgSign, reqTimestamp, reqNonce, reqData, &verify)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(verify)
	}
}

func TestCallbackDecode_EncryptMsg(t *testing.T) {

}
