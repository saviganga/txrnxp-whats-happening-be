package configs

import (
	// "errors"

	"os"
	// "github.com/joho/godotenv"
)

type Config struct {
	DBPort             string
	DBName             string
	DBUser             string
	DBPassword         string
	AllowedOrigins     string
	Environment        string
	Version            string
	ImageKitPublicKey  string
	ImageKitPrivateKey string
	ImageKitURL        string
}

func NewConfig() *Config {
	return &Config{}
}

// LoadConfig loads and validates environment variables, returning any errors encountered
func (c *Config) LoadConfig() error {
	// load the .env file
	// err := godotenv.Load()
	// if err != nil {
	// 	// This is not a critical error; we can rely on system environment variables
	// 	return errors.New("no .env file found, relying on system environment variables")
	// }

	// Load environment variables
	c.DBName = getEnv("DB_NAME_LOCAL", "")
	c.DBUser = getEnv("DB_USER_LOCAL", "")
	c.DBPassword = getEnv("DB_PASSWORD_LOCAL", "")
	c.AllowedOrigins = getEnv("ALLOWED_ORIGINS", "*")
	c.Environment = getEnv("ENVIRONMENT", "local")
	c.Version = getEnv("VERSION", "1")
	c.ImageKitPrivateKey = mustGetEnv("IMAGEKIT_PRIVATE_KEY")
	c.ImageKitPublicKey = mustGetEnv("IMAGEKIT_PUBLIC_KEY")
	c.ImageKitURL = mustGetEnv("IMAGEKIT_URL")

	return nil
}

// getEnv retrieves the value of the environment variable or returns a default value
func getEnv(key, defaultValue string) string {
	value, exists := os.LookupEnv(key)
	if !exists {
		return defaultValue
	}
	return value
}

// mustGetEnv retrieves the value of the environment variable or returns an error if not set
func mustGetEnv(key string) string {
	value, exists := os.LookupEnv(key)
	if !exists {
		panic("required environment variable not set: " + key)
	}
	return value
}

// func getEnvAsSlice(key, defaultValue string) []string {
// 	value, exists := os.LookupEnv(key)
// 	if !exists || value == "" {
// 		return strings.Split(defaultValue, ",")
// 	}
// 	return strings.Split(value, ",")
// }
