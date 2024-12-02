package config

type Config struct {
	Server ServerConfig `toml:"server"`
}

type ServerConfig struct {
	Port string `toml:"port"`
}

func DefaultConfig() Config {
	return Config{
		Server: ServerConfig{
			Port: ":0",
		},
	}
}
