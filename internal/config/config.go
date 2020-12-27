package config

import "os"

type Config struct {
	LarkHookId string
}

func NewConfig() *Config {
	return &Config{}
}

func (c *Config) Load() {
	c.LarkHookId = c.GetLarkHookId()
}

func (c *Config) GetLarkHookId() string {
	return os.Getenv("HOOK_ID")
}

var Settings *Config

func init() {
	config := NewConfig()
	config.Load()
	Settings = config
}
