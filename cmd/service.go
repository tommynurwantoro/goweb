package cmd

import (
	"github.com/spf13/cobra"
	"github.com/tommynurwantoro/goweb/config"
	"github.com/tommynurwantoro/goweb/internal/bootstrap"
)

func RunService() *cobra.Command {
	var configFile string
	command := &cobra.Command{
		Use:     "service",
		Aliases: []string{"svc"},
		Short:   "Run the service",
		Run: func(cmd *cobra.Command, args []string) {
			conf := &config.Model{}
			conf.LoadConfig(configFile)

			bstp := bootstrap.New()
			bstp.Run(conf)
		},
	}

	command.Flags().StringVarP(&configFile, "config", "c", "config.yaml", "config file")

	return command
}
