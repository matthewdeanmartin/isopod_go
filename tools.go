//go:build tools
// +build tools

package tools

import (
	_ "github.com/golangci/golangci-lint/cmd/golangci-lint"
	_ "honnef.co/go/tools/cmd/staticcheck"
)
