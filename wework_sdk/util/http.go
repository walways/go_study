package util

// Import resty into your code and refer it as `resty`.
import (
	"github.com/go-resty/resty/v2"
)

type HttpClient interface {
	Get(url string) (resp *Response, err error)
	Post(url string, body map[string]interface{}) (resp *Response, err error)
}

type Response struct {
	BodyString string `json:"body_string"`
	Body       []byte `json:"body"`
}

type RestClient struct {
	restyClient *resty.Client
}

func New() (client *RestClient) {
	return &RestClient{
		restyClient: resty.New(),
	}
}

func (c *RestClient) Get(url string) (resp *Response, err error) {
	restResp, err := c.restyClient.R().Get(url)
	if err != nil {
		return nil, err
	}
	return &Response{
		Body:       restResp.Body(),
		BodyString: string(restResp.Body()),
	}, nil
}

func (c *RestClient) Post(url string, body map[string]interface{}) (resp *Response, err error) {
	restResp, err := c.restyClient.R().SetBody(body).Post(url)
	if err != nil {
		return nil, err
	}
	return &Response{
		Body:       restResp.Body(),
		BodyString: string(restResp.Body()),
	}, nil
}
