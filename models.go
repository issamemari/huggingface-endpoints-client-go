package huggingface

import "time"

type ListEndpointResponse struct {
	Endpoints []Endpoint `json:"items"`
}

type Endpoint struct {
	AccountId *string  `json:"accountId"`
	Compute   Compute  `json:"compute"`
	Model     Model    `json:"model"`
	Name      string   `json:"name"`
	Provider  Provider `json:"provider"`
	Status    *Status  `json:"status"`
	Type      string   `json:"type"`
}

type Compute struct {
	Accelerator  string  `json:"accelerator"`
	InstanceSize string  `json:"instanceSize"`
	InstanceType string  `json:"instanceType"`
	Scaling      Scaling `json:"scaling"`
}

type Scaling struct {
	MaxReplica         int `json:"maxReplica"`
	MinReplica         int `json:"minReplica"`
	ScaleToZeroTimeout int `json:"scaleToZeroTimeout"`
}

type Model struct {
	Framework  string `json:"framework"`
	Image      Image  `json:"image"`
	Repository string `json:"repository"`
	Revision   string `json:"revision"`
	Task       string `json:"task"`
}

type Image struct {
	Huggingface Huggingface `json:"huggingface"`
}

type Huggingface struct {
	Env map[string]interface{} `json:"env"`
}

type Provider struct {
	Region string `json:"region"`
	Vendor string `json:"vendor"`
}

type Status struct {
	CreatedAt     time.Time `json:"createdAt"`
	CreatedBy     User      `json:"createdBy"`
	ErrorMessage  string    `json:"errorMessage"`
	Message       string    `json:"message"`
	Private       Private   `json:"private"`
	ReadyReplica  int       `json:"readyReplica"`
	State         string    `json:"state"`
	TargetReplica int       `json:"targetReplica"`
	UpdatedAt     time.Time `json:"updatedAt"`
	UpdatedBy     User      `json:"updatedBy"`
	URL           string    `json:"url"`
}

type User struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type Private struct {
	ServiceName string `json:"serviceName"`
}
