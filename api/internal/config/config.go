package config

import "os"

type Config struct {
	DatabaseURL         string
	JWTSecret           string
	ClaudeAPIKey        string
	LinkedInClientID    string
	LinkedInClientSecret string
	LinkedInRedirectURI string
	APIPort             string
	FrontendURL         string
}

func Load() Config {
	return Config{
		DatabaseURL:         getEnv("DATABASE_URL", "postgres://coverletter:coverletter_secret@localhost:5432/coverletter_hub?sslmode=disable"),
		JWTSecret:           getEnv("JWT_SECRET", "dev-secret"),
		ClaudeAPIKey:        getEnv("CLAUDE_API_KEY", ""),
		LinkedInClientID:    getEnv("LINKEDIN_CLIENT_ID", ""),
		LinkedInClientSecret: getEnv("LINKEDIN_CLIENT_SECRET", ""),
		LinkedInRedirectURI: getEnv("LINKEDIN_REDIRECT_URI", "http://localhost:8080/api/v1/auth/linkedin/callback"),
		APIPort:             getEnv("API_PORT", "8080"),
		FrontendURL:         getEnv("FRONTEND_URL", "http://localhost:3000"),
	}
}

func getEnv(key, fallback string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return fallback
}
