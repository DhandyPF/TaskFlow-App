package config

import (
	"os"
	"strings"

	"github.com/joho/godotenv"
)

// Config holds runtime configuration loaded from environment variables.
type Config struct {
	Port           string
	DatabaseDSN    string
	JWTSecret      string
	AllowedOrigins []string
}

// Load reads configuration from a .env file (if present) and environment
// variables, falling back to sensible defaults for local development.
func Load() *Config {
	// Ignore the error: it's fine if no .env file exists, e.g. in
	// production where real environment variables are set directly.
	_ = godotenv.Load()

	return &Config{
		Port:           getEnv("PORT", "8080"),
		DatabaseDSN:    getEnv("DATABASE_DSN", "./taskflow.db"),
		JWTSecret:      getEnv("JWT_SECRET", "change-this-secret-in-production"),
		AllowedOrigins: getEnvList("ALLOWED_ORIGINS", nil),
	}
}

func getEnv(key, fallback string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return fallback
}

// getEnvList reads a comma-separated environment variable into a string
// slice, e.g. ALLOWED_ORIGINS=https://app.vercel.app,https://foo.com
// An unset or empty variable returns fallback (nil means "allow any origin",
// handled by the CORS middleware).
func getEnvList(key string, fallback []string) []string {
	raw := os.Getenv(key)
	if raw == "" {
		return fallback
	}
	parts := strings.Split(raw, ",")
	result := make([]string, 0, len(parts))
	for _, p := range parts {
		p = strings.TrimSpace(p)
		if p != "" {
			result = append(result, p)
		}
	}
	return result
}
