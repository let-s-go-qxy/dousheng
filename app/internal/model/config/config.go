package config

type Config struct {
	Logger   *Logger   `mapstructure:"logger" yaml:"logger"`
	Cors     CORS      `mapstructure:"cors" yaml:"cors"`
	DataBase *Database `mapstructure:"database"  yaml:"database"`
	Server   *Server   `mapstructure:"server"  yaml:"server"`
	Auth     Auth      `mapstructure:"auth" yaml:"auth"`
}
