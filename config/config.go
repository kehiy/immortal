package config

import (
	"os"

	"github.com/dezh-tech/immortal/client"
	"github.com/dezh-tech/immortal/database"
	"github.com/dezh-tech/immortal/handler"
	"github.com/dezh-tech/immortal/relay/redis"
	"github.com/dezh-tech/immortal/server/grpc"
	"github.com/dezh-tech/immortal/server/websocket"
	"github.com/joho/godotenv"
	"gopkg.in/yaml.v3"
)

// Config represents the configs used by relay and other concepts on system.
type Config struct {
	Environment     string           `yaml:"environment"`
	Kraken          client.Config    `yaml:"kraken"`
	WebsocketServer websocket.Config `yaml:"ws_server"`
	Database        database.Config  `yaml:"database"`
	RedisConf       redis.Config     `yaml:"redis"`
	GRPCServer      grpc.Config      `yaml:"grpc_server"`
	Handler         handler.Config
}

// Load loads config from file and env.
func Load(path string) (*Config, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, Error{
			reason: err.Error(),
		}
	}
	defer file.Close()

	config := &Config{}

	decoder := yaml.NewDecoder(file)

	if err := decoder.Decode(config); err != nil {
		return nil, Error{
			reason: err.Error(),
		}
	}

	if config.Environment != "prod" {
		if err := godotenv.Load(); err != nil {
			return nil, Error{
				reason: err.Error(),
			}
		}
	}

	config.Database.URI = os.Getenv("IMMO_MONGO_URI")
	config.RedisConf.URI = os.Getenv("IMMO_REDIS_URI")

	if err = config.basicCheck(); err != nil {
		return nil, Error{
			reason: err.Error(),
		}
	}

	return config, nil
}

// basicCheck validates the basic stuff in config.
func (c *Config) basicCheck() error {
	return nil
}
