package config

type Config struct {
	Database DatabaseConfig
	Server   ServerConfig
}

type DatabaseConfig struct {
	Host     string `json:"host"`
	Port     string `json:"port"`
	User     string `json:"user"`
	Password string `json:"password"`
	DBname   string `json:"dbname"`
}

type ServerConfig struct {
	Port string `json:"port"`
}
