package finance_yahoo

import "net/http"

type Client struct {
	httpClient *http.Client
}

func NewClient() *Client {
	return &Client{httpClient: http.DefaultClient}
}

func (c *Client) WithHTTPClient(t http.RoundTripper) {
	c.httpClient.Transport = t
}

func (c *Client) Assets() {

}
