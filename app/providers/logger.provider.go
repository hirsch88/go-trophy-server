package providers

import (
	"github.com/hirsch88/go-trophy-server/config"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"strings"
)

func NewLoggerProvider(config *config.AppConfig) *zap.SugaredLogger {
	var log *zap.Logger

	atom := zap.NewAtomicLevel()
	atom.SetLevel(parseLevel(config.LogLevel))

	logConfig := zap.NewDevelopmentConfig()
	logConfig.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	logConfig.Level = atom
	log, _ = logConfig.Build()

	return log.Sugar()
}

func parseLevel(level string) zapcore.Level {
	switch strings.ToLower(level) {

	case "debug":
		return zap.DebugLevel

	case "info":
		return zap.InfoLevel

	case "warn":
		return zap.WarnLevel

	default:
		return zap.ErrorLevel
	}

}
