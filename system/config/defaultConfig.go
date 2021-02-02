package config

type DefaultConfig struct {
	Server Server `mapstructure:"server"`
}

type Server struct {
	Port        int
	PortEcho    int    `mapstructure:"port-echo"`
	RootName    string `mapstructure:"root-name"`
	HttpHandler string `mapstructure:"http-handler"`
}
