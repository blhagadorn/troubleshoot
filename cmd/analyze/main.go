package main

import (
	"github.com/replicatedhq/troubleshoot/cmd/analyze/cli"
	_ "k8s.io/client-go/plugin/pkg/client/auth/gcp"
)

func main() {
	cli.InitAndExecute()
}
