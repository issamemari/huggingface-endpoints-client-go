package huggingface

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

var client *Client

func TestMain(m *testing.M) {
	host := "https://api.endpoints.huggingface.cloud/v2/endpoint"
	namespace := "issamemari"
	token := ""

	var err error
	client, err = NewClient(&host, &namespace, &token)
	if err != nil {
		panic(err)
	}

	m.Run()
}

func randomString(n int) string {
	const letters = "abcdefghijklmnopqrstuvwxyz"
	rng := rand.New(rand.NewSource(time.Now().UnixNano()))
	b := make([]byte, n)
	for i := range b {
		b[i] = letters[rng.Intn(len(letters))]
	}
	return string(b)
}

func newCreateEndpointRequest() CreateEndpointRequest {
	name := fmt.Sprintf("test-endpoint-%s", randomString(4))
	scaleToZeroTimeout := 15
	revision := "main"
	task := "sentence-embeddings"
	return CreateEndpointRequest{
		AccountId: nil,
		Compute: Compute{
			Accelerator:  "cpu",
			InstanceSize: "x4",
			InstanceType: "intel-icl",
			Scaling: Scaling{
				MinReplica:         0,
				MaxReplica:         1,
				ScaleToZeroTimeout: &scaleToZeroTimeout,
			},
		},
		Model: Model{
			Framework: "pytorch",
			Image: Image{
				Huggingface: &Huggingface{
					Env: map[string]string{},
				},
			},
			Repository: "sentence-transformers/all-MiniLM-L6-v2",
			Revision:   &revision,
			Task:       &task,
		},
		Name: name,
		Provider: Provider{
			Region: "us-east-1",
			Vendor: "aws",
		},
		Type: "protected",
	}
}

func newCreateEndpointRequestWithCustomImage() CreateEndpointRequest {
	endpoint := newCreateEndpointRequest()
	endpoint.Model.Image.Custom = &Custom{
		Credentials: &Credentials{
			Password: "password",
			Username: "username",
		},
		Env: map[string]string{
			"key": "value",
		},
		HealthRoute: nil,
		Port:        nil,
		URL:         "https://example.com",
	}
	endpoint.Model.Image.Huggingface = nil
	return endpoint
}

func TestCustomImage(t *testing.T) {
	endpoint := newCreateEndpointRequestWithCustomImage()

	_, err := client.CreateEndpoint(endpoint)
	if err != nil {
		panic(err)
	}

	err = client.DeleteEndpoint(endpoint.Name)
	if err != nil {
		panic(err)
	}
}

func TestNilCredentials(t *testing.T) {
	endpoint := newCreateEndpointRequestWithCustomImage()
	endpoint.Model.Image.Custom.Credentials = nil

	_, err := client.CreateEndpoint(endpoint)
	if err != nil {
		panic(err)
	}

	err = client.DeleteEndpoint(endpoint.Name)
	if err != nil {
		panic(err)
	}
}

func TestEmptyEnv(t *testing.T) {
	endpoint := newCreateEndpointRequest()
	endpoint.Model.Image.Huggingface.Env = map[string]string{}

	_, err := client.CreateEndpoint(endpoint)
	if err != nil {
		panic(err)
	}

	err = client.DeleteEndpoint(endpoint.Name)
	if err != nil {
		panic(err)
	}
}

func TestListEndpoints(t *testing.T) {
	_, err := client.ListEndpoints()
	if err != nil {
		panic(err)
	}
}

func TestCreateAndDeleteEndpoint(t *testing.T) {
	endpoint := newCreateEndpointRequest()

	_, err := client.CreateEndpoint(endpoint)
	if err != nil {
		panic(err)
	}

	err = client.DeleteEndpoint(endpoint.Name)
	if err != nil {
		panic(err)
	}
}

func TestGetEndpoint(t *testing.T) {
	endpoint := newCreateEndpointRequest()

	_, err := client.CreateEndpoint(endpoint)
	if err != nil {
		panic(err)
	}

	_, err = client.GetEndpoint(endpoint.Name)
	if err != nil {
		panic(err)
	}

	err = client.DeleteEndpoint(endpoint.Name)
	if err != nil {
		panic(err)
	}
}

func TestUpdateEndpoint(t *testing.T) {
	endpoint := newCreateEndpointRequest()

	_, err := client.CreateEndpoint(endpoint)
	if err != nil {
		panic(err)
	}

	updateEndpointRequest := UpdateEndpointRequest{
		Compute: &Compute{
			Accelerator:  "cpu",
			InstanceSize: "x8",
			InstanceType: "intel-icl",
			Scaling: Scaling{
				MinReplica: 0,
				MaxReplica: 1,
			},
		},
		Model: &endpoint.Model,
		Type:  nil,
	}

	_, err = client.UpdateEndpoint(endpoint.Name, updateEndpointRequest)
	if err != nil {
		panic(err)
	}

	err = client.DeleteEndpoint(endpoint.Name)
	if err != nil {
		panic(err)
	}
}

func TestOptionalFields(t *testing.T) {
	endpoint := newCreateEndpointRequest()
	endpoint.Model.Revision = nil
	endpoint.Compute.Scaling.ScaleToZeroTimeout = nil

	_, err := client.CreateEndpoint(endpoint)
	if err != nil {
		panic(err)
	}

	err = client.DeleteEndpoint(endpoint.Name)
	if err != nil {
		panic(err)
	}
}
