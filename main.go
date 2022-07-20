package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"

	"k8s.io/component-base/logs"
	"k8s.io/kubernetes/cmd/kube-scheduler/app"
	"rds-dynamic-scheduler/plugin"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	command := app.NewSchedulerCommand(
		app.WithPlugin(plugin.Name, plugin.NewAllInOnePlugin),
	)

	logs.InitLogs()
	defer logs.FlushLogs()
	if err := command.Execute(); err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}
}
