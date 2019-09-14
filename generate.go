package main

import (
	"github.com/solo-io/gloo/pkg/version"
	"github.com/solo-io/go-utils/log"
	"github.com/solo-io/solo-kit/pkg/code-generator/cmd"
	"github.com/solo-io/solo-kit/pkg/code-generator/docgen/options"
)

//go:generate go run generate.go

func main() {
	err := version.CheckVersions()
	if err != nil {
		log.Fatalf("generate failed!: %v", err)
	}
	log.Printf("starting generate")

	generateOptions := cmd.GenerateOptions{
		CustomCompileProtos: []string{"projects/gloo/api/grpc"},
		SkipGenMocks:       true,
		CustomImports:      nil,
		SkipGeneratedTests: true,
		SkipDirs:           nil,
		RelativeRoot:       "",
		CompileProtos:      true,
		GenDocs: &cmd.DocsOptions{
			Output: options.Hugo,
		},
	}
	if err := cmd.Generate(generateOptions); err != nil {
		log.Fatalf("generate failed!: %v", err)
	}
}
