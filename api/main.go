package main

import (
	"context"
	"fmt"
	"os"

	"github.com/setalid/snapcar/api/pkg/api"
)

func main() {
	ctx := context.Background()
	if err := api.Run(ctx); err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(1)
	}
}
