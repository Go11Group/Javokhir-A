package config

type Configs struct {
	Database DatabaseConfig
	Server   ServerConfig
}

type ServerConfig struct {
	Port string
}

type DatabaseConfig struct {
	Host     string `json:"host"`
	Port     string `json:"port"`
	User     string `json:"user"`
	Password string `json:"password"`
	Name     string `json:"name"`
	SSLMode  string `json:"sslmode"`
}
