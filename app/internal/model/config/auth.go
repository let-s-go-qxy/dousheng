package config

import "net/http"

type Auth struct {
	Jwt    Jwt    `mapstructure:"jwt" yaml:"jwt"`
	Cookie Cookie `mapstructure:"cookie" yaml:"cookie"`
}

type Jwt struct {
	SecretKey   string `mapstructure:"secretKey" yaml:"secretKey"`
	ExpiresTime int64  `mapstructure:"expiresTime" yaml:"expiresTime"`
	BufferTime  int64  `mapstructure:"bufferTime" yaml:"bufferTime"`
	Issuer      string `mapstructure:"issuer" yaml:"issuer"`
}

type Cookie struct {
	Secret      string `mapstructure:"secret" yaml:"secret"`
	http.Cookie `mapstructure:",squash"`
}
