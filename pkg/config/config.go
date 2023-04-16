package config

type Config struct {
	DBHost     string `mapstructure:"DB_HOST"`
	DBName     string `mapstructure:"DB_NAME"`
	DBUser     string `mapstructure:"DB_USER"`
	DBPassword string `mapstructure:"DB_PASSWORD"`
	DBPort     string `mapstructure:"DB_PORT"`
	JWT        string `mapstructure:"JWT_CODE"`
	AUTHTOKEN  string `mapstructure:"AUTH_TOKEN"`
	ACCOUNTSID string `mapstructure:"ACCOUNT_SID"`
	SERVICESID string `mapstructure:"SERVICE_SID"`
}
