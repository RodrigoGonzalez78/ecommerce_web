package config

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Config struct {
	Port        string
	DatabaseURL string
	SecretKeyQR string
	JWTSecret   string
}

var Cnf *Config

func getEnv(key, fallback string, required bool) string {
	value := os.Getenv(key)
	if value == "" {
		if required {
			log.Fatalf("❌ La variable de entorno %s es obligatoria pero no está definida", key)
		}
		log.Printf("⚠️  La variable de entorno %s no está definida. Usando valor por defecto: %s", key, fallback)
		return fallback
	}
	return value
}

func getEnvBool(key string, fallback bool) bool {
	valStr := os.Getenv(key)
	if valStr == "" {
		return fallback
	}
	val, err := strconv.ParseBool(valStr)
	if err != nil {
		log.Printf("⚠️  No se pudo parsear %s como booleano. Usando valor por defecto: %v", key, fallback)
		return fallback
	}
	return val
}

func LoadConfig() {
	if err := godotenv.Load(); err != nil {
		log.Println("📄 Archivo .env no encontrado. Usando variables de entorno del sistema.")
	}

	Cnf = &Config{
		Port:        getEnv("PORT", "8080", false),
		DatabaseURL: getEnv("DATABASE_URL", "", true),
		SecretKeyQR: getEnv("SECRET_KEY_QR", "default_qr_key", false),
		JWTSecret:   getEnv("JWT_SECRET", "", true),
	}
}
