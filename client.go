package huggingface

import (
	"fmt"
	"io"
	"net/http"
	"time"
)

const HostURL string = "https://api.endpoints.huggingface.cloud/v2/endpoint"

type Client struct {
	HostURL    string
	HTTPClient *http.Client
	Token      string
	Namespace  string
}

func NewClient(host, namespace, token *string) (*Client, error) {
	c := Client{
		HTTPClient: &http.Client{Timeout: 10 * time.Second},
		// Default HuggingFace endpoints URL
		HostURL: HostURL,
	}

	if host != nil {
		c.HostURL = *host
	}

	if token == nil || namespace == nil {
		return &c, nil
	}

	c.Token = *token
	c.Namespace = *namespace

	return &c, nil
}

func (c *Client) DoRequest(req *http.Request) ([]byte, error) {
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", c.Token))

	res, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("status: %d, body: %s", res.StatusCode, body)
	}

	return body, err
}
