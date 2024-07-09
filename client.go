package huggingface

import (
	"bytes"
	"encoding/json"
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

func (c *Client) DoRequest(method, path string, body interface{}) ([]byte, *int, error) {
	var reqBody *bytes.Buffer
	if body != nil {
		reqBody = new(bytes.Buffer)
		err := json.NewEncoder(reqBody).Encode(body)
		if err != nil {
			return nil, nil, fmt.Errorf("%w: %v", ErrEncodingRequestBody, err)
		}
	} else {
		reqBody = bytes.NewBuffer([]byte{})
	}

	url := c.HostURL
	if c.Namespace != "" {
		url = fmt.Sprintf("%s/%s", url, c.Namespace)
	}
	if path != "" {
		url = fmt.Sprintf("%s/%s", url, path)
	}

	req, err := http.NewRequest(method, url, reqBody)
	if err != nil {
		return nil, nil, fmt.Errorf("%w: %v", ErrCreatingRequest, err)
	}

	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", c.Token))
	req.Header.Set("Content-Type", "application/json")

	res, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, nil, err
	}
	defer res.Body.Close()

	respBody, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, nil, fmt.Errorf("%w: %v", ErrReadingResponseBody, err)
	}

	return respBody, &res.StatusCode, nil
}
