package wework_sdk

import (
	"errors"
	"fmt"
	"wework_sdk/util"
)

type Config struct {
	Token        string `json:"token"`
	EncodeAesKey string `json:"encode_aes_key"`
}

type CallbackDecode struct {
	client *util.WXBizMsgCrypt
}

// 这里默认用JsonType
func New(receiverId string, config *Config) (client *CallbackDecode) {
	return &CallbackDecode{
		client: util.NewWXBizMsgCrypt(config.Token, config.EncodeAesKey, receiverId, util.JsonType),
	}
}

/*
		------------使用示例一：验证回调URL---------------
		*企业开启回调模式时，企业微信会向验证url发送一个get请求
		假设点击验证时，企业收到类似请求：
		* GET /cgi-bin/wxpush?msg_signature=5c45ff5e21c57e6ad56bac8758b79b1d9ac89fd3&timestamp=1409659589&nonce=263014780&echostr=P9nAzCzyDtyTWESHep1vC5X9xho%2FqYX3Zpb4yKa9SKld1DsH3Iyt3tP3zNdtp%2B4RPcs8TgAE7OaBO%2BFZXvnaqQ%3D%3D
		* HTTP/1.1 Host: qy.weixin.qq.com

		接收到该请求时，企业应
	     1.解析出Get请求的参数，包括消息体签名(msg_signature)，时间戳(timestamp)，随机数字串(nonce)以及企业微信推送过来的随机加密字符串(echostr),
	     这一步注意作URL解码。
	     2.验证消息体签名的正确性
	     3. 解密出echostr原文，将原文当作Get请求的response，返回给企业微信
	     第2，3步可以用企业微信提供的库函数VerifyURL来实现。
*/
func (c *CallbackDecode) VerifyURL(verifyMsgSign, verifyTimestamp, verifyNonce, verifyEchoStr string) (string, error) {
	echoStr, cryptError := c.client.VerifyURL(verifyMsgSign, verifyTimestamp, verifyNonce, verifyEchoStr)
	if nil != cryptError {
		return "", errors.New(fmt.Sprintf("verifyUrl fail, err:%+v", cryptError))
	}
	return string(echoStr), nil
}

/*
		------------使用示例二：对用户回复的消息解密---------------
		用户回复消息或者点击事件响应时，企业会收到回调消息，此消息是经过企业微信加密之后的密文以post形式发送给企业，密文格式请参考官方文档
		假设企业收到企业微信的回调消息如下：
		POST /cgi-bin/wxpush? msg_signature=477715d11cdb4164915debcba66cb864d751f3e6&timestamp=1409659813&nonce=1372623149 HTTP/1.1
		Host: qy.weixin.qq.com
		Content-Length: 613
		{
	     "tousername":"wx5823bf96d3bd56c7",
	     "encrypt":"CZWs4CWRpI4VolQlvn4dlPBlXke6+HgmuI7p0LueFp1fKH40TNL+YHWJZwqIiYV+3kTrhdNU7fZwc+PmtgBvxSczkFeRz+oaVSsomrrtP2Z91LE313djjbWujqInRT+7ChGbCeo7ZzszByf8xnDSunPBxRX1MfX3kAxpKq7dqduW1kpMAx8O8xUzZ9oC0TLuZchbpxaml4epzGfF21O+zyXDwTxbCEiO0E87mChtzuh/VPlznXYbfqVrnyLNZ5pr",
		    "agentid":"218"
	 }

		企业收到post请求之后应该：
	     1.解析出url上的参数，包括消息体签名(msg_signature)，时间戳(timestamp)以及随机数字串(nonce)
	     2.验证消息体签名的正确性。
	     3.将post请求的数据进行json解析，并将"Encrypt"标签的内容进行解密，解密出来的明文即是用户回复消息的明文，明文格式请参考官方文档
	     第2，3步可以用企业微信提供的库函数DecryptMsg来实现。
*/
func (c *CallbackDecode) DecryptMsg(reqMsgSign, reqTimestamp, reqNonce, reqData string, obj interface{}) error {
	msg, cryptError := c.client.DecryptMsg(reqMsgSign, reqTimestamp, reqNonce, []byte(reqData))
	if nil != cryptError {
		return errors.New(fmt.Sprintf("DecryptMsg fail %+v", cryptError))
	}
	err := util.Unmarshal(msg, &obj)
	return err
}

/*
------------使用示例三：企业回复用户消息的加密---------------
企业被动回复用户的消息也需要进行加密，并且拼接成密文格式的json串。
假设企业需要回复用户的明文如下：

	{
	    "ToUserName": "mycreate",
	    "FromUserName":"wx5823bf96d3bd56c7",
	    "CreateTime": 1348831860,
	    "MsgType": "text",
	    "Content": "this is a test",
	    "MsgId": 1234567890123456,
	    "AgentID": 128
	}

为了将此段明文回复给用户，企业应：

	1.自己生成时间时间戳(timestamp),随机数字串(nonce)以便生成消息体签名，也可以直接用从企业微信的post url上解析出的对应值。
	2.将明文加密得到密文。
	3.用密文，步骤1生成的timestamp,nonce和企业在企业微信设定的token生成消息体签名。
	4.将密文，消息体签名，时间戳，随机数字串拼接成json格式的字符串，发送给企业。
	以上2，3，4步可以用企业微信提供的库函数EncryptMsg来实现。
*/
//respData := "{\"ToUserName\":\"wx5823bf96d3bd56c7\",\"FromUserName\":\"mycreate\",\"CreateTime\": 1409659813,\"MsgType\":\"text\",\"Content\":\"hello\",\"MsgId\":4561255354251345929,\"AgentID\": 218}"
func (c *CallbackDecode) EncryptMsg(respData, timestamp, nonce string) (string, error) {
	encryptMsg, cryptErr := c.client.EncryptMsg(respData, timestamp, nonce)
	if nil != cryptErr {
		return "", errors.New(fmt.Sprintf("EncryptMsg fail %+v", cryptErr))
	}
	return string(encryptMsg), cryptErr
}
