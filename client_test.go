package huggingface

import (
	"testing"
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

func TestListEndpoints(t *testing.T) {
	_, err := client.ListEndpoints()
	if err != nil {
		panic(err)
	}
}

func TestCreateAndDeleteEndpoint(t *testing.T) {
	endpoint := Endpoint{
		AccountId: nil,
		Compute: Compute{
			Accelerator:  "cpu",
			InstanceSize: "x4",
			InstanceType: "intel-icl",
			Scaling: Scaling{
				MinReplica:         0,
				MaxReplica:         1,
				ScaleToZeroTimeout: 15,
			},
		},
		Model: Model{
			Framework: "pytorch",
			Image: Image{
				Huggingface: Huggingface{
					Env: map[string]interface{}{},
				},
			},
			Repository: "sentence-transformers/all-MiniLM-L6-v2",
			Revision:   "main",
			Task:       "sentence-embeddings",
		},
		Name: "issa-test-endpoint",
		Provider: Provider{
			Region: "us-east-1",
			Vendor: "aws",
		},
		Type: "protected",
	}

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
	endpointId := "issa-test-endpoint"
	_, err := client.GetEndpoint(endpointId)
	if err != nil {
		panic(err)
	}
}
