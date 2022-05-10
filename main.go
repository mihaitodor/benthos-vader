package main

import (
	"context"

	"github.com/benthosdev/benthos/v4/public/service"

	// Import all standard Benthos components
	_ "github.com/benthosdev/benthos/v4/public/components/all"

	// Import plugins
	_ "github.com/mihaitodor/benthos-vader/bloblang"
)

func main() {
	service.RunCLI(context.Background())
}
