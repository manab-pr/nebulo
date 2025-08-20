package config

import (
	"log"
	"os"
	"strconv"
	"time"

	"github.com/joho/godotenv"
)

type Config struct {
	Server    ServerConfig
	Database  DatabaseConfig
	Redis     RedisConfig
	JWT       JWTConfig
	Storage   StorageConfig
	Device    DeviceConfig
}

type ServerConfig struct {
	Port string
	Env  string
	Name string
}

type DatabaseConfig struct {
	URI      string
	Database string
}

type RedisConfig struct {
	Addr     string
	Password string
	DB       int
}

type JWTConfig struct {
	Secret    string
	ExpiresIn time.Duration
}

type StorageConfig struct {
	Path        string
	MaxFileSize int64
}

type DeviceConfig struct {
	ServerPort        string
	HeartbeatInterval time.Duration
	TransferTimeout   time.Duration
}

func LoadConfig() *Config {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, using environment variables")
	}

	redisDB, _ := strconv.Atoi(getEnv("REDIS_DB", "0"))
	jwtExpiresIn, _ := time.ParseDuration(getEnv("JWT_EXPIRES_IN", "24h"))
	heartbeatInterval, _ := time.ParseDuration(getEnv("HEARTBEAT_INTERVAL", "30s"))
	transferTimeout, _ := time.ParseDuration(getEnv("TRANSFER_TIMEOUT", "300s"))
	
	maxFileSize := parseFileSize(getEnv("MAX_FILE_SIZE", "100MB"))

	return &Config{
		Server: ServerConfig{
			Port: getEnv("PORT", "8080"),
			Env:  getEnv("APP_ENV", "development"),
			Name: getEnv("APP_NAME", "Nebulo"),
		},
		Database: DatabaseConfig{
			URI:      getEnv("MONGO_URI", "mongodb://localhost:27017"),
			Database: getEnv("MONGO_DATABASE", "nebulo"),
		},
		Redis: RedisConfig{
			Addr:     getEnv("REDIS_ADDR", "localhost:6379"),
			Password: getEnv("REDIS_PASSWORD", ""),
			DB:       redisDB,
		},
		JWT: JWTConfig{
			Secret:    getEnv("JWT_SECRET", "default-secret-key"),
			ExpiresIn: jwtExpiresIn,
		},
		Storage: StorageConfig{
			Path:        getEnv("STORAGE_PATH", "./storage"),
			MaxFileSize: maxFileSize,
		},
		Device: DeviceConfig{
			ServerPort:        getEnv("DEVICE_SERVER_PORT", "8081"),
			HeartbeatInterval: heartbeatInterval,
			TransferTimeout:   transferTimeout,
		},
	}
}

func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}

func parseFileSize(sizeStr string) int64 {
	// Simple parser for sizes like "100MB", "1GB", etc.
	if len(sizeStr) < 3 {
		return 100 * 1024 * 1024 // Default 100MB
	}
	
	unit := sizeStr[len(sizeStr)-2:]
	sizeNumStr := sizeStr[:len(sizeStr)-2]
	sizeNum, err := strconv.ParseInt(sizeNumStr, 10, 64)
	if err != nil {
		return 100 * 1024 * 1024 // Default 100MB
	}
	
	switch unit {
	case "KB":
		return sizeNum * 1024
	case "MB":
		return sizeNum * 1024 * 1024
	case "GB":
		return sizeNum * 1024 * 1024 * 1024
	default:
		return sizeNum // Assume bytes
	}
}