package config

import (
	"fmt"
)

type Database struct {
	Mysql *Mysql `mapstructure:"mysql" yaml:"mysql"`
	Mongo *Mongo `mapstructure:"mongo" yaml:"mongo"`
	Redis *Redis `mapstructure:"redis" yaml:"redis"`
}

type Mysql struct {
	Addr     string `mapstructure:"addr" yaml:"addr"`
	Port     string `mapstructure:"port" yaml:"port"`
	Db       string `mapstructure:"db" yaml:"db"`
	Username string `mapstructure:"username" yaml:"username"`
	Password string `mapstructure:"password" yaml:"password"`
	Charset  string `mapstructure:"charset" yaml:"charset"`
}

func (m *Mysql) GetDsn() string {
	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=True&loc=Australia%%2FMelbourne",
		m.Username,
		m.Password,
		m.Addr,
		m.Port,
		m.Db,
		m.Charset)
}

type Mongo struct {
	Addr     string `mapstructure:"addr" yaml:"addr"`
	Port     string `mapstructure:"port" yaml:"port"`
	Username string `mapstructure:"username" yaml:"username"`
	Password string `mapstructure:"password" yaml:"password"`
}

func (m *Mongo) GetAddr() string {
	return fmt.Sprintf("mongodb://%s:%s", m.Addr, m.Port)
}

type Redis struct {
	Addr        string `mapstructure:"addr" yaml:"addr"`
	Port        string `mapstructure:"port" yaml:"port"`
	Password    string `mapstructure:"password" yaml:"password"`
	DbVerify    int    `mapstructure:"dbVerify" yaml:"dbVerify"`
	DbUserLike  int    `mapstructure:"dbUserLike" yaml:"dbUserLike"`
	DbVideoLike int    `mapstructure:"dbVideoLike" yaml:"dbVideoLike"`
}
