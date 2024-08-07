package config

import (
	"dm/db"
	"dm/internal/gateway"
	"dm/internal/server"
)

type Config struct {
	Server   server.Config  `toml:"server"`
	Database db.Config      `toml:"database"`
	Gateway  gateway.Config `toml:"gateway"`
}
