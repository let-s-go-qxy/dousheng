package config

type Logger struct {
	SavePath     string `mapstructure:"savePath" yaml:"savePath"`
	EncoderType  string `mapstructure:"encoderType" yaml:"encoderType"`
	EncodeLevel  string `mapstructure:"encodeLevel" yaml:"encodeLevel"`
	EncodeCaller string `mapstructure:"encodeCaller" yaml:"encodeCaller"`
}
