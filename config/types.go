package config

type Config struct {
	Deploy  string  `yaml:"deploy" env-default:"local"`
	Storage Storage `yaml:"storage"`
	Server  Server  `yaml:"server"`
	JWT     JWT     `yaml:"jwt"`
}

type Storage struct {
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
	DBName   string `yaml:"dbname"`
	SSLMode  string `yaml:"ssl"`
}

type Server struct {
	Bind string `yaml:"bind"`
	Port int    `yaml:"port"`
}

type JWT struct {
	SecretKey string `yaml:"secret_key"`
}
