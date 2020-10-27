package config

import (
	"fmt"
)

// Config is service config.
type Config struct {
	Database DatabaseConfig
	Redis    RedisConfig
}

// DatabaseConfig database config.
type DatabaseConfig struct {
	Dialect      string `envconfig:"DB_DIALECT" yaml:"dialect" default:"mysql"`
	User         string `envconfig:"DB_USER" yaml:"user" default:"root"`
	Password     string `envconfig:"DB_PASSWORD" yaml:"password" default:"password"`
	Host         string `envconfig:"DB_HOST" yaml:"host" default:"127.0.0.1"`
	Port         string `envconfig:"DB_PORT" yaml:"port" default:"3306"`
	Name         string `envconfig:"DB_NAME" yaml:"name" default:"shake"`
	Option       string `envconfig:"DB_OPTION" yaml:"option" default:"?charset=utf8mb4&parseTime=True&loc=Asia%2FTokyo"`
	Logging      bool   `envconfig:"DB_LOGGING" yaml:"logging" default:"true"`
	MaxIdleConns int    `envconfig:"DB_MAX_IDLE_CONNS" yaml:"max_idle_conns" default:"10"`
	MaxOpenConns int    `envconfig:"DB_MAX_OPEN_CONNS" yaml:"max_open_conns" default:"20"`
	AutoMigrate  bool   `default:false`
}

func (a *DatabaseConfig) DSN() string {
	if a.Dialect == "sqlite3" {
		return a.Name
	}
	return fmt.Sprintf(
		"%s:%s@(%s:%s)/%s%s",
		a.User,
		a.Password,
		a.Host,
		a.Port,
		a.Name,
		a.Option,
	)
}

// RedisConfig redis config.
type RedisConfig struct {
	Host     string `envconfig:"REDIS_HOST" yaml:"host" default:"localhost"`
	Port     string `envconfig:"REDIS_PORT" yaml:"port" default:"6379"`
	Password string `envconfig:"REDIS_PASSWORD" yaml:"password" default:""`
	DB       int    `envconfig:"REDIS_DB" yaml:"db" default:"0"`
}

func (a *RedisConfig) Addr() string {
	return fmt.Sprintf("%s:%s", a.Host, a.Port)
}
