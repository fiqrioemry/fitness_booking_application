package config

import (
	"os"
	"strings"
)

func GetTrustedProxies() []string {
	raw := os.Getenv("TRUSTED_PROXIES")
	if raw == "" {
		return []string{"127.0.0.1"}
	}
	parts := strings.Split(raw, ",")
	for i := range parts {
		parts[i] = strings.TrimSpace(parts[i])
	}
	return parts
}
