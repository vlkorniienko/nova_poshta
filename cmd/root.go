package cmd

import (
	"os"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/rs/zerolog/pkgerrors"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "novaposhta",
	Short: "Service for handling request to NovaPoshta inc",
	PersistentPreRun: func(_ *cobra.Command, _ []string) {
		zerolog.ErrorStackMarshaler = pkgerrors.MarshalStack

		log.Logger = log.Output(zerolog.ConsoleWriter{
			Out: os.Stderr,
		})
		log.Info().Msg("Enabled human readable logging!")

		validLogLevel, err := zerolog.ParseLevel("debug")
		if err != nil {
			log.Fatal().Err(err).Msg("Failed to parse log level")
		}

		log.Logger = log.Level(validLogLevel).With().Stack().Logger()

		zerolog.DefaultContextLogger = &log.Logger
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}

var configFile string

func init() {
	rootCmd.PersistentFlags().StringVar(&configFile, "config-file", "", "path to config file; optional")
}
