package config

import (
	"github.com/EgorTarasov/lct-2024/api/pkg/mongo"
	"github.com/EgorTarasov/lct-2024/api/pkg/postgres"
	"github.com/EgorTarasov/lct-2024/api/pkg/redis"
	"github.com/EgorTarasov/lct-2024/api/pkg/s3"
	"github.com/EgorTarasov/lct-2024/api/pkg/telemetry"
	"github.com/ilyakaznacheev/cleanenv"
)

type server struct {
	Port        int      `yaml:"port"`
	Domain      string   `yaml:"domain"`
	CorsOrigins []string `yaml:"cors-origins"`
}

// TODO: add comments.
type vkAuth struct {
	VkTokenURL     string `yaml:"vk-token-url"`
	VkClientID     string `yaml:"vk-client-id"`
	VkSecureToken  string `yaml:"vk-secure-token"`
	VkServiceToken string `yaml:"vk_service_token"`
	VkRedirectURI  string `yaml:"vk-redirect-uri"`
}

// Config для работы приложения.
type Config struct {
	Server    *server           `yaml:"http-server"`
	Telemetry *telemetry.Config `yaml:"telemetry"`
	Database  *postgres.Config  `yaml:"postgres"`
	Redis     *redis.Config     `yaml:"redis"`
	VkAuth    *vkAuth           `yaml:"vk-users"`
	S3        *s3.Config        `yaml:"s3"`
	Mongo     *mongo.Config     `yaml:"mongo"`
}

// MustNew создает новый конфиг из файла и завершает программу в случае ошибки.
func MustNew(path string) *Config {
	cfg := &Config{}
	if err := cleanenv.ReadConfig(path, cfg); err != nil {
		panic(err)
	}
	return cfg
}
