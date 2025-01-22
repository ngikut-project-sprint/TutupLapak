package config

type DatabaseConfig struct {
	Username string `env:"DB_USER"`
	Password string `env:"DB_PASSWORD"`
	Host     string `env:"DB_HOST"`
	Port     uint16 `env:"DB_PORT"`
	Database string `env:"DB_NAME"`
	SslMode  string `env:"DB_SSLMODE"`
}

type JWTConfig struct {
	Secret string `env:"JWT_SECRET"`
}

type AWSConfig struct {
	Region     string `env:"AWS_REGION"`
	AccessKey  string `env:"AWS_ACCESS_KEY_ID"`
	SecretKey  string `env:"AWS_SECRET_ACCESS_KEY"`
	BucketName string `env:"AWS_BUCKET_NAME"`
}

type Config struct {
	Database DatabaseConfig
	JWT      JWTConfig
	AWS      AWSConfig
}
