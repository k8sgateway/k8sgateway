package main

import (
	"context"

	"github.com/solo-io/gloo/projects/ingress/pkg/setup"
	"github.com/solo-io/go-utils/log"
)

func main() {
	if err := setup.Main(context.Background()); err != nil {
		log.Fatalf("err in main: %v", err.Error())
	}
}
