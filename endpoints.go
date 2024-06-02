package huggingface

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) DoRequest(method, path string, body interface{}) ([]byte, error) {
	var reqBody *bytes.Buffer
	if body != nil {
		reqBody = new(bytes.Buffer)
		err := json.NewEncoder(reqBody).Encode(body)
		if err != nil {
			return nil, fmt.Errorf("failed to encode request body: %w", err)
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
		return nil, fmt.Errorf("failed to create new request: %w", err)
	}

	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", c.Token))
	req.Header.Set("Content-Type", "application/json")

	res, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	respBody, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("status: %d, url: %s, body: %s", res.StatusCode, url, body)
	}

	return respBody, err
}

func (c *Client) ListEndpoints() ([]Endpoint, error) {
	body, err := c.DoRequest("GET", "", nil)
	if err != nil {
		return nil, fmt.Errorf("failed to list endpoints: %w", err)
	}

	response := ListEndpointResponse{}
	err = json.Unmarshal(body, &response)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal response: %w", err)
	}

	return response.Endpoints, nil
}

func (c *Client) CreateEndpoint(endpoint Endpoint) (Endpoint, error) {
	body, err := c.DoRequest("POST", "", endpoint)
	if err != nil {
		return Endpoint{}, fmt.Errorf("failed to create endpoint: %w", err)
	}

	var response Endpoint
	err = json.Unmarshal(body, &response)
	if err != nil {
		return Endpoint{}, fmt.Errorf("failed to unmarshal response: %w", err)
	}

	return response, nil
}

func (c *Client) GetEndpoint(name string) (Endpoint, error) {
	body, err := c.DoRequest("GET", name, nil)
	if err != nil {
		return Endpoint{}, fmt.Errorf("failed to get endpoint: %w", err)
	}

	var response Endpoint
	err = json.Unmarshal(body, &response)
	if err != nil {
		return Endpoint{}, fmt.Errorf("failed to unmarshal response: %w", err)
	}

	return response, nil
}

func (c *Client) DeleteEndpoint(name string) error {
	_, err := c.DoRequest("DELETE", name, nil)
	if err != nil {
		return fmt.Errorf("failed to delete endpoint: %w", err)
	}
	return nil
}

func (c *Client) UpdateEndpoint(name string, endpoint Endpoint) (Endpoint, error) {
	body, err := c.DoRequest("PUT", name, endpoint)
	if err != nil {
		return Endpoint{}, fmt.Errorf("failed to update endpoint: %w", err)
	}

	var response Endpoint
	err = json.Unmarshal(body, &response)
	if err != nil {
		return Endpoint{}, fmt.Errorf("failed to unmarshal response: %w", err)
	}

	return response, nil
}
