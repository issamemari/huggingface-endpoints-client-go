package huggingface

import (
	"testing"
)

var client *Client

func TestMain(m *testing.M) {
	host := "https://api.endpoints.huggingface.cloud/v2/endpoint"
	namespace := ""
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

func TestCreateEndpoint(t *testing.T) {
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
			Repository: "GorgiasML/article_reranker",
			Revision:   "696d7548fe4bad7a9da6b846b05ce4416ab89f07",
			Task:       "custom",
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
}
