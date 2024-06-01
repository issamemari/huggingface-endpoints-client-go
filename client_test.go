package huggingface

import (
	"fmt"
	"testing"
)

func TestListEndpoints(t *testing.T) {
	host := "https://api.endpoints.huggingface.cloud/v2/endpoint"
	namespace := ""
	token := ""

	client, err := NewClient(&host, &namespace, &token)
	if err != nil {
		panic(err)
	}

	endpoints, err := client.ListEndpoints()
	if err != nil {
		panic(err)
	}

	// print the results
	for _, endpoint := range endpoints {
		fmt.Printf("Endpoint: %s\n", endpoint.Name)
		fmt.Printf("  Account ID: %s\n", endpoint.AccountId)
		fmt.Printf("  Type: %s\n", endpoint.Type)
		fmt.Printf("  Status: %s\n", endpoint.Status.State)
		fmt.Printf("  Created At: %s\n", endpoint.Status.CreatedAt)
		fmt.Printf("  Framework: %s\n", endpoint.Model.Framework)
		fmt.Printf("  Task: %s\n", endpoint.Model.Task)
		fmt.Printf("  Image: %s\n", endpoint.Model.Image.Huggingface)
		fmt.Printf("  Region: %s\n", endpoint.Provider.Region)
		fmt.Printf("  Vendor: %s\n", endpoint.Provider.Vendor)
		fmt.Printf("  Instance Type: %s\n", endpoint.Compute.InstanceType)
		fmt.Printf("  Instance Size: %s\n", endpoint.Compute.InstanceSize)
		fmt.Printf("  Accelerator: %s\n", endpoint.Compute.Accelerator)
		fmt.Printf("  Max Replica: %d\n", endpoint.Compute.Scaling.MaxReplica)
		fmt.Printf("  Min Replica: %d\n", endpoint.Compute.Scaling.MinReplica)
		fmt.Printf("  Scale To Zero Timeout: %d\n", endpoint.Compute.Scaling.ScaleToZeroTimeout)
		fmt.Println()
	}
}
