package main

import (
	"context"
	"github.com/spf13/cobra"
	"os/signal"
	"projectIraUw/cmd"
	"projectIraUw/infra"
	"syscall"
)

func main() {
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGTERM, syscall.SIGINT, syscall.SIGKILL)
	defer stop()

	config := infra.NewConfig()
	containerDI := infra.NewContainerDI(config)
	defer containerDI.ShutDown()

	cmdHttpServer := &cobra.Command{
		Use:   "httpserver",
		Short: "Run httpserver",
		Run: func(cli *cobra.Command, args []string) {
			cmd.StartHttp(ctx, containerDI)
		},
	}

	var rootCmd = &cobra.Command{Use: "APP"}
	rootCmd.AddCommand(cmdHttpServer)
	rootCmd.Execute()
}
