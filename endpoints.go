package huggingface

import (
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
