package huggingface

type ListEndpointResponse struct {
	Endpoints []EndpointDetails `json:"items"`
}

type EndpointDetails struct {
	AccountId *string  `json:"accountId,omitempty"`
	Compute   Compute  `json:"compute"`
	Model     Model    `json:"model"`
	Name      string   `json:"name"`
	Provider  Provider `json:"provider"`
	Status    Status   `json:"status"`
	Type      string   `json:"type"`
}

type Status struct {
	CreatedAt     string  `json:"createdAt"`
	CreatedBy     User    `json:"createdBy"`
	UpdatedAt     string  `json:"updatedAt"`
	UpdatedBy     User    `json:"updatedBy"`
	Private       Private `json:"private"`
	State         string  `json:"state"`
	Message       string  `json:"message"`
	ReadyReplica  int     `json:"readyReplica"`
	TargetReplica int     `json:"targetReplica"`
}

type CreateEndpointRequest struct {
	AccountId *string  `json:"accountId,omitempty"`
	Compute   Compute  `json:"compute"`
	Model     Model    `json:"model"`
	Name      string   `json:"name"`
	Provider  Provider `json:"provider"`
	Type      string   `json:"type"`
}

type UpdateEndpointRequest struct {
	Compute *Compute `json:"compute,omitempty"`
	Model   *Model   `json:"model,omitempty"`
	Type    *string  `json:"type,omitempty"`
}

type Compute struct {
	Accelerator  string  `json:"accelerator"`
	InstanceSize string  `json:"instanceSize"`
	InstanceType string  `json:"instanceType"`
	Scaling      Scaling `json:"scaling"`
}

type Scaling struct {
	MaxReplica         int  `json:"maxReplica"`
	MinReplica         int  `json:"minReplica"`
	ScaleToZeroTimeout *int `json:"scaleToZeroTimeout,omitempty"`
}

type Model struct {
	Framework  string  `json:"framework"`
	Image      Image   `json:"image"`
	Repository string  `json:"repository"`
	Revision   *string `json:"revision,omitempty"`
	Task       *string `json:"task,omitempty"`
}

type Image struct {
	Huggingface *Huggingface `json:"huggingface,omitempty"`
	Custom      *Custom      `json:"custom,omitempty"`
}

type Custom struct {
	Credentials *Credentials      `json:"credentials,omitempty"`
	Env         map[string]string `json:"env,omitempty"`
	HealthRoute *string           `json:"health_route,omitempty"`
	Port        *int              `json:"port,omitempty"` // Constraints: Min 0, Default: 80
	URL         string            `json:"url"`
}

type Credentials struct {
	Password string `json:"password"`
	Username string `json:"username"`
}

type Huggingface struct {
	Env map[string]string `json:"env,omitempty"`
}

type Provider struct {
	Region string `json:"region"`
	Vendor string `json:"vendor"`
}

type User struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type Private struct {
	ServiceName string `json:"serviceName"`
}
