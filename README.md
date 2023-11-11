# Kemp Load Balancer Go Client

This Go client library provides a convenient way to interact with Kemp Load Balancers, enabling the management of virtual services, real servers, network interfaces, and more.

## Structure

The project is structured as follows:


### Internal

- `commands/`: Contains command implementations for interacting with the Kemp Load Balancer.
- `helpers/`: Provides helper functions used across the project.

### pkg

- `client/`: Core client implementations for the Kemp Load Balancer.
- `types/`: Data structures and types relevant to Kemp Load Balancers.

## Features

### KempClient

- `NewKempClient`: Initialize a new client to interact with a Kemp Load Balancer.
- `ListVirtualServices`: List all virtual services.
- `GetVirtualServiceByID`: Get a virtual service by its ID.
- `GetVirtualServiceByIP`: Retrieve a virtual service by its IP, port, and protocol.
- `AddLayer4VirtualService`: Add a new Layer 4 virtual service.
- `AddLayer7VirtualService`: Add a new Layer 7 virtual service.
- `DeleteVirtualService`: Delete a virtual service.

### Real Servers and Network Interfaces

- Functions to manage real (backend) servers and network interfaces.

### Types

- Data structures representing various entities such as virtual services, real servers, network interfaces, etc.

## Getting Started

1. Clone the repository.
2. Install dependencies: `go get ./...`
3. Import the package in your Go project.

## Usage

Initialize the client and use its methods to interact with your Kemp Load Balancer:

```go
package main

import (
	"fmt"
	"github.com/scotta01-org/kemp-lb-go-client/pkg/client"
)

func main() {
	kempClient := client.NewKempClient("your-kemp-address", "your-api-key", true)

	// List virtual services
	virtualServices, err := kempClient.ListVirtualServices()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%+v\n",virtualServices)
}
```
## Contributing

Contributions are welcome! Please feel free to submit pull requests or open issues.

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.