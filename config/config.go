package config

import "time"

type TypeUserIDctx string

var (
	Conf Config

	ValueUserIDctx TypeUserIDctx = "ValueUserIDctx"
)

type Config struct {
	App          Application       `mapstructure:"app"`
	ConnDatabase DatabaseConfig    `mapstructure:"postgres"`
	Redis        RedisConfig       `mapstructure:"redis"`
	RateLimiter  RateLimiterConfig `mapstructure:"rateLimiter"`
	JwtSecret    string            `mapstructure:"jwtSecret"`
}

type Application struct {
	Env       string        `mapstructure:"env"`
	Name      string        `mapstructure:"name"`
	Port      int           `mapstructure:"port"`
	LogFormat string        `mapstructure:"logFormat"`
	Timeout   time.Duration `mapstructure:"timeout"`
}

type DatabaseConfig struct {
	DbConnectionInfo
	Slave  DBInfo `mapstructure:"slave"`
	Master DBInfo `mapstructure:"master"`
}

type DBInfo struct {
	DBName   string `mapstructure:"dbName"`
	Host     string `mapstructure:"host"`
	Password string `mapstructure:"password"`
	Port     int    `mapstructure:"port"`
	Schema   string `mapstructure:"schema"`
	User     string `mapstructure:"user"`
	Debug    bool   `mapstructure:"debug"`
}

type DbConnectionInfo struct {
	SetMaxIdleCons    int `mapstructure:"maxIdleConnections"`
	SetMaxOpenCons    int `mapstructure:"maxOpenConnections"`
	SetConMaxIdleTime int `mapstructure:"setConMaxIdleTime"`
	SetConMaxLifetime int `mapstructure:"connectTimeout"`
}

type RedisConfig struct {
	Host                    string        `mapstructure:"host"`
	DB                      int           `mapstructure:"db"`
	Password                string        `mapstructure:"password"`
	FailedLoginAttemptTTL   time.Duration `mapstructure:"failedLoginAttemptTTL"`
	FailedLoginIssuspendTTL time.Duration `mapstructure:"failedLoginIssuspendTTL"`
	LastPageCandidateTTL    time.Duration `mapstructure:"lastPageCandidateTTL"`
	LockingSwipeActionTTL   time.Duration `mapstructure:"lockingSwipeActionTTL"`
	CountSwipeActionTTL     time.Duration `mapstructure:"countSwipeActionTTL"`
}

type RateLimiterConfig struct {
	MaxSwipeAction  int `mapstructure:"maxSwipeAction"`
	MaxAttemptLogin int `mapstructure:"maxAttemptLogin"`
}
