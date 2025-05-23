﻿package amqpexporter

import (
	"errors"

	"go.opentelemetry.io/collector/component"
	"go.uber.org/multierr"
)

var errNoEndpoint = errors.New("no endpoint specified")

type Config struct {
	Address string `mapstructure:"address"`
}

var _ component.Config = (*Config)(nil)

func (c *Config) Validate() error {
	var err error

	if c.Address == "" {
		err = multierr.Append(err, errNoEndpoint)
	}

	return err
}
