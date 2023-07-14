package config

var (
	Conf Config
)

type Config struct {
	App          Application `mapstructure:"app"`
	ConnDatabase Database    `mapstructure:"postgres"`
}

type Application struct {
	Env       string `mapstructure:"env"`
	Name      string `mapstructure:"name"`
	Port      int    `mapstructure:"port"`
	LogFormat string `mapstructure:"logFormat"`
}

type Database struct {
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
