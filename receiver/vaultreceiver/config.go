// Copyright 2020, OpenTelemetry Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package vaultreceiver

import (
	"errors"
	"time"

	"go.opentelemetry.io/collector/config"
)

var _ config.Receiver = (*Config)(nil)

type Config struct {
	config.ReceiverSettings `mapstructure:",squash"`
	// The URL of the docker server.  Default is "unix:///var/run/docker.sock"
	Endpoint string `mapstructure:"endpoint"`
	// The time between each collection event.  Default is 10s.
	CollectionInterval time.Duration `mapstructure:"collection_interval"`

	// The maximum amount of time to wait for docker API responses.  Default is 5s
	Timeout time.Duration `mapstructure:"timeout"`
}

func (config Config) Validate() error {
	if config.Endpoint == "" {
		return errors.New("config.Endpoint must be specified")
	}
	if config.CollectionInterval == 0 {
		return errors.New("config.CollectionInterval must be specified")
	}
	return nil
}
