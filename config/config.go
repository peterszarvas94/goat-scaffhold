package config

import (
	"log/slog"
)

var (
	AppName  = "scaffhold"
	LogLevel = slog.LevelDebug
)

type envT struct {
	DbPath string
	Env    string
	Port   string
}

var Vars envT
