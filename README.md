# HuggingFace Endpoints API Go Client

This is a Go client library for interacting with HuggingFace's endpoints API. It provides methods for creating, retrieving, updating, listing, and deleting endpoints on HuggingFace.

## Features

- List all endpoints
- Create a new endpoint
- Retrieve details of a specific endpoint
- Update an existing endpoint
- Delete an endpoint

## Installation

To install the package, run:

```sh
go get github.com/issamemari/huggingface-endpoints-client-go
```

## Usage

### Initialize the Client

Before using the client, you need to initialize it with your HuggingFace API token, host URL, and a namespace.

```go
package main

import (
    "github.com/yourusername/huggingface"
    "net/http"
)

func main() {
    client := huggingface.Client{
        HostURL:   "https://api.endpoints.huggingface.cloud/v2/endpoint",
        Namespace: "your-namespace",
        Token:     "your-api-token",
        HTTPClient: &http.Client{},
    }

    // Use the client...
}
```

### List Endpoints

```go
endpoints, err := client.ListEndpoints()
if err != nil {
    // Handle error
}
for _, endpoint := range endpoints {
    fmt.Println(endpoint)
}
```

### Create Endpoint

```go
newEndpoint := huggingface.Endpoint{
    // Fill in endpoint details
}

createdEndpoint, err := client.CreateEndpoint(newEndpoint)
if err != nil {
    // Handle error
}
fmt.Println(createdEndpoint)
```

### Get Endpoint

```go
endpoint, err := client.GetEndpoint("endpoint-name")
if err != nil {
    // Handle error
}
fmt.Println(endpoint)
```

### Update Endpoint

```go
updatedEndpoint := huggingface.Endpoint{
    // Fill in updated endpoint details
}

endpoint, err := client.UpdateEndpoint("endpoint-name", updatedEndpoint)
if err != nil {
    // Handle error
}
fmt.Println(endpoint)
```

### Delete Endpoint

```go
err := client.DeleteEndpoint("endpoint-name")
if err != nil {
    // Handle error
}
```

## Error Handling

Each method returns an error as the second return value which should be checked to ensure that the operation was successful.

## Contributing

Contributions are welcome! Please open an issue or submit a pull request for any bugs, improvements, or new features.

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

---

Happy coding!