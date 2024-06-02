package huggingface

import (
	"encoding/json"
	"fmt"
)

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
