package config

import (
	"log/slog"
)

var (
	AppName  = "scaffhold"
	Port     = "9999"
	LogLevel = slog.LevelDebug
)

type envT struct {
	DbPath string
	Env    string
}

var Vars envT
