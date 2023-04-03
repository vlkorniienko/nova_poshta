package cmd

import (
	"context"

	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"

	"github.com/vlkorniienko/novaposhta/pkg/build"
	"github.com/vlkorniienko/novaposhta/pkg/common/config"
	"github.com/vlkorniienko/novaposhta/pkg/entrypoint/http/novaposhta"
)

// serveCmd represents the serve command.
var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "Serves HTTP handlers",
	Run:   serve,
}

func init() {
	rootCmd.AddCommand(serveCmd)
}

type serveConfig struct {
	HTTP build.HTTPConfig  `config:"http"`
	NP   novaposhta.Config `config:"nova_poshta"`
}

func serve(_ *cobra.Command, _ []string) {
	var conf serveConfig

	err := config.ReadFromFile(configFile, &conf)
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to read configuration")
	}

	novaPoshtaService := novaposhta.NewService(conf.NP.APIKey)

	server, err := build.NewHTTPServer(conf.HTTP, novaPoshtaService)
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to create server")
	}

	err = server.Serve(context.Background())
	if err != nil {
		log.Fatal().Err(err).Msg("listen and serve error")
	}
}
