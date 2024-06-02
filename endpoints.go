package huggingface

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

func (c *Client) ListEndpoints() ([]Endpoint, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/%s", c.HostURL, c.Namespace), nil)
	if err != nil {
		return nil, err
	}

	body, err := c.DoRequest(req)
	if err != nil {
		return nil, err
	}

	response := ListEndpointResponse{}
	err = json.Unmarshal(body, &response)
	if err != nil {
		return nil, err
	}

	return response.Endpoints, nil
}

func (c *Client) CreateEndpoint(endpoint Endpoint) (Endpoint, error) {
	response := Endpoint{}

	reqBody, err := json.Marshal(endpoint)
	if err != nil {
		return response, err
	}

	req, err := http.NewRequest("POST", fmt.Sprintf("%s/%s", c.HostURL, c.Namespace), bytes.NewBuffer(reqBody))
	if err != nil {
		return response, err
	}

	body, err := c.DoRequest(req)
	if err != nil {
		return response, err
	}

	err = json.Unmarshal(body, &response)
	if err != nil {
		return response, err
	}

	return response, nil
}

func (c *Client) GetEndpoint(name string) (Endpoint, error) {
	response := Endpoint{}

	req, err := http.NewRequest("GET", fmt.Sprintf("%s/%s/%s", c.HostURL, c.Namespace, name), nil)
	if err != nil {
		return response, err
	}

	body, err := c.DoRequest(req)
	if err != nil {
		return response, err
	}

	err = json.Unmarshal(body, &response)
	if err != nil {
		return response, err
	}

	return response, nil
}

func (c *Client) DeleteEndpoint(name string) error {
	req, err := http.NewRequest("DELETE", fmt.Sprintf("%s/%s/%s", c.HostURL, c.Namespace, name), nil)
	if err != nil {
		return err
	}

	_, err = c.DoRequest(req)
	if err != nil {
		return err
	}

	return nil
}
