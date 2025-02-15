package config

import (
	"sync"

	"github.com/caarlos0/env/v11"
)

type Config struct {
	TenantKey         string `env:"TenantKey" envDefault:"73c3a1fb-0abc-453d-a804-565b4a211b34"`
	TenantSecret      string `env:"TenantSecret" envDefault:"fe1e549f-5940-4813-9a6e-6ca943b7cc8c"`
	TenantAPIEndPoint string `env:"TenantAPIEndPoint" envDefault:"https://devapi.lrinternal.com"`
	TokenSignKey      string `env:"TokenSignKey" envDefault:"1234567890123456"`
}

var (
	configInstance     *Config
	configInstanceOnce sync.Once
)

func GetInstance() *Config {
	configInstanceOnce.Do(func() {
		var cfg Config
		_ = env.Parse(&cfg)
		configInstance = &cfg
	})
	return configInstance
}
