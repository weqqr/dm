package config

import (
	"dm/internal/db"
	"dm/internal/server"
)

type Config struct {
	Server   server.Config `toml:"server"`
	Database db.Config     `toml:"database"`
}
