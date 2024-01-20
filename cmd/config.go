package cmd

type Config struct {
	Schedule string `env:"SCHEDULE"`
	Data     string `env:"DATA"`
}
