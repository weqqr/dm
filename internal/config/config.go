package config

import (
	"dm/db"
	"dm/internal/server"
	"dm/internal/user/userproxy"
)

type Config struct {
	Server   server.Config    `toml:"server"`
	Database db.Config        `toml:"database"`
	Gateway  userproxy.Config `toml:"gateway"`
}
