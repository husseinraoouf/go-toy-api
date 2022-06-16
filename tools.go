//go:build tools
// +build tools

package main

import (
	// for swagger
	_ "github.com/go-swagger/go-swagger/cmd/swagger"

	// for golangci-lint
	_ "github.com/golangci/golangci-lint/cmd/golangci-lint"
)
