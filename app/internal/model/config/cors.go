package config

type CORS struct {
	Mode      string          `mapstructure:"mode" yaml:"mode"`
	Whitelist []CORSWhitelist `mapstructure:"whitelist" yaml:"whitelist"`
}

type CORSWhitelist struct {
	AllowOrigin      string `mapstructure:"allowOrigin" yaml:"allowOrigin"`
	AllowMethods     string `mapstructure:"allowMethods" yaml:"allowMethods"`
	AllowHeaders     string `mapstructure:"allowHeaders" yaml:"allowHeaders"`
	ExposeHeaders    string `mapstructure:"exposeHeaders" yaml:"exposeHeaders"`
	AllowCredentials bool   `mapstructure:"allowCredentials" yaml:"allowCredentials"`
}
