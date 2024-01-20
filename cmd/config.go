package cmd

type Config struct {
	Schedule   string            `env:"SCHEDULE"`
	URLs       map[string]string `env:"URLS"`
	Data       string            `env:"DATA"`
	Metube     string            `env:"METUBE"`
	Initialize bool              `env:"INITIALIZE" envDefault:"false"`
}
