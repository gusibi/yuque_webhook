package lark

import (
	"context"
	logger "github.com/sirupsen/logrus"
	"github.com/valyala/fasthttp"
	urllib "net/url"
	"time"
)

type Client struct {
	HttpClient *fasthttp.Client
}

func NewHttpClient(maxConns, defaultTimeout int) *Client {
	return &Client{HttpClient: &fasthttp.Client{
		Name:            "larkHttpClient",
		MaxConnsPerHost: maxConns,
		ReadTimeout:     time.Second * time.Duration(defaultTimeout),
		RetryIf:         nil,
	}}
}

func (c *Client) request(ctx context.Context, method, url string, query string, body []byte, timeout time.Duration) (*fasthttp.Response, error) {
	req, resp := fasthttp.AcquireRequest(), fasthttp.AcquireResponse()
	defer func() {
		fasthttp.ReleaseResponse(resp)
		fasthttp.ReleaseRequest(req)
	}()

	req.SetRequestURI(url)

	// 默认是application/x-www-form-urlencoded
	req.Header.SetContentType("application/json")
	req.Header.SetMethod(method)

	req.SetBody(body)
	req.URI().SetQueryString(query)

	if err := c.HttpClient.DoTimeout(req, resp, timeout); err != nil {
		logger.Error("request: %s url:%s fail | err:%+v", method, url, err)
		return nil, err
	}
	return resp, nil
}

func (c *Client) Get(ctx context.Context, url string, queryParams urllib.Values, timeout time.Duration) (*fasthttp.Response, error) {
	query := queryParams.Encode()
	return c.request(ctx, fasthttp.MethodGet, url, query, nil, timeout)
}

func (c *Client) Post(ctx context.Context, url string, body []byte, timeout time.Duration) (*fasthttp.Response, error) {
	return c.request(ctx, fasthttp.MethodPost, url, "", body, timeout)
}
