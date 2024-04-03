package main

import (
	cmd "github.com/kloud-team/kloud/internal/cmd/kloud"
	pkg "github.com/kloud-team/kloud/internal/pkg/kloud"
)

func main() {
	defer pkg.PanicHandler()
	cmd.Execute()
}
