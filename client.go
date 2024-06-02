package huggingface

import (
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
		HostURL:    HostURL,
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
