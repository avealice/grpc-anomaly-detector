package config

type Config struct {
	UserName string
	DBName   string
	Port     string
	SSLMode  string
}

func NewConfig() *Config {
	return &Config{
		UserName: "postgres",
		DBName:   "report",
		Port:     "5432",
		SSLMode:  "disable",
	}
}
