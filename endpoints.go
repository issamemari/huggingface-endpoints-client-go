package huggingface

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func (c *Client) ListEndpoints() ([]EndpointDetails, error) {
	body, statusCode, err := c.DoRequest("GET", "", nil)
	if err != nil {
		return nil, fmt.Errorf("%w: %v", ErrExecutingRequest, err)
	}
	if *statusCode != http.StatusOK {
		bodyStr := string(body)
		return nil, &HTTPError{
			StatusCode: *statusCode,
			Body:       &bodyStr,
			Message:    "failed to list endpoints",
		}
	}

	response := ListEndpointResponse{}
	err = json.Unmarshal(body, &response)
	if err != nil {
		return nil, fmt.Errorf("%w: %v", ErrUnmarshalingResponse, err)
	}

	return response.Endpoints, nil
}

func (c *Client) CreateEndpoint(endpoint CreateEndpointRequest) (EndpointDetails, error) {
	body, statusCode, err := c.DoRequest("POST", "", endpoint)
	if err != nil {
		return EndpointDetails{}, fmt.Errorf("%w: %v", ErrExecutingRequest, err)
	}
	if *statusCode != http.StatusOK {
		bodyStr := string(body)
		return EndpointDetails{}, &HTTPError{
			StatusCode: *statusCode,
			Body:       &bodyStr,
			Message:    "failed to create endpoint",
		}
	}

	var response EndpointDetails
	err = json.Unmarshal(body, &response)
	if err != nil {
		return EndpointDetails{}, fmt.Errorf("%w: %v", ErrUnmarshalingResponse, err)
	}

	return response, nil
}

func (c *Client) GetEndpoint(name string) (EndpointDetails, error) {
	body, statusCode, err := c.DoRequest("GET", name, nil)
	if err != nil {
		return EndpointDetails{}, fmt.Errorf("%w: %v", ErrExecutingRequest, err)
	}

	if *statusCode != http.StatusOK {
		bodyStr := string(body)
		return EndpointDetails{}, &HTTPError{
			StatusCode: *statusCode,
			Body:       &bodyStr,
			Message:    "failed to get endpoint",
		}
	}

	var response EndpointDetails
	err = json.Unmarshal(body, &response)
	if err != nil {
		return EndpointDetails{}, fmt.Errorf("%w: %v", ErrUnmarshalingResponse, err)
	}

	return response, nil
}

func (c *Client) DeleteEndpoint(name string) error {
	body, statusCode, err := c.DoRequest("DELETE", name, nil)
	if err != nil {
		return fmt.Errorf("%w: %v", ErrExecutingRequest, err)
	}
	if statusCode != nil && *statusCode != http.StatusOK {
		bodyStr := string(body)
		return &HTTPError{
			StatusCode: *statusCode,
			Body:       &bodyStr,
			Message:    "failed to delete endpoint",
		}
	}
	return nil
}

func (c *Client) UpdateEndpoint(name string, endpoint UpdateEndpointRequest) (EndpointDetails, error) {
	body, statusCode, err := c.DoRequest("PUT", name, endpoint)
	if err != nil {
		return EndpointDetails{}, fmt.Errorf("%w: %v", ErrExecutingRequest, err)
	}
	if *statusCode != http.StatusOK {
		bodyStr := string(body)
		return EndpointDetails{}, &HTTPError{
			StatusCode: *statusCode,
			Body:       &bodyStr,
			Message:    "failed to update endpoint",
		}
	}

	var response EndpointDetails
	err = json.Unmarshal(body, &response)
	if err != nil {
		return EndpointDetails{}, fmt.Errorf("%w: %v", ErrUnmarshalingResponse, err)
	}

	return response, nil
}
