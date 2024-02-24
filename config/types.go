package config

type Config struct {
	Application ApplicationConfig
	Database    DatabaseConfig
	Redis       DatabaseConfig
	Auth        AuthConfig
}

type ApplicationConfig struct {
	ServiceName    string
	ServiceVersion string
	Env            string
	ServerPort     string
	ServerHost     string
}

type DatabaseConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	Database string
}

type AuthConfig struct {
	Lifespan  int
	ApiSecret string
}
