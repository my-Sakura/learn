package embed

import (
	_ "embed"
	"strings"
)

var (
	Version string = strings.TrimSpace(version)
	//go:embed version.txt
	version string

	//go:embed embed.go
	Src string
)
