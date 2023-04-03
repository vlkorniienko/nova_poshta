package config

import (
	"errors"
	"fmt"

	"github.com/iamolegga/enviper"
	"github.com/joho/godotenv"
	"github.com/mitchellh/mapstructure"
	"github.com/spf13/viper"
)

var ErrUnsupportedConfigType = errors.New("can't read configuration")

func ReadFromFile(filepath string, dst interface{}) error {
	if filepath[len(filepath)-4:] == ".env" {
		return readFromEnv(filepath, dst)
	}

	return ErrUnsupportedConfigType
}

func readFromEnv(filepath string, dst interface{}) error {
	if err := godotenv.Load(filepath); err != nil {
		return fmt.Errorf("loading env file: %w", err)
	}

	const tagName = "config"
	v := enviper.New(viper.GetViper()).WithTagName(tagName)

	confOption := func(c *mapstructure.DecoderConfig) {
		c.Squash = true
	}

	if err := v.Unmarshal(dst, confOption); err != nil {
		return fmt.Errorf("unmarshaling: %w", err)
	}

	return nil
}
