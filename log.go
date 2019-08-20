package logger

import (
	"strings"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var logger *zap.Logger

var defaultConfig = zap.Config{
	Encoding:         "json",
	Level:            atom,
	OutputPaths:      []string{"stderr"},
	ErrorOutputPaths: []string{"stderr"},
	EncoderConfig: zapcore.EncoderConfig{
		MessageKey: "message",

		LevelKey:    "level",
		EncodeLevel: zapcore.CapitalLevelEncoder,

		TimeKey:    "time",
		EncodeTime: zapcore.ISO8601TimeEncoder,

		CallerKey:    "caller",
		EncodeCaller: zapcore.ShortCallerEncoder,
	},
}

// Zap .
type Zap struct {
	Level            string   `json:"level"`
	OutputPaths      []string `json:"output_paths"`
	ErrorOutputPaths []string `json:"error_output_paths"`
}

var atom = zap.NewAtomicLevel()

// Init 初始化
func Init(zapCfg Zap) (err error) {
	atom.SetLevel(parseLevel(zapCfg.Level))
	if len(zapCfg.OutputPaths) > 0 {
		defaultConfig.OutputPaths = zapCfg.OutputPaths
	}
	if len(zapCfg.ErrorOutputPaths) > 0 {
		defaultConfig.ErrorOutputPaths = zapCfg.ErrorOutputPaths
	}
	logger, err = defaultConfig.Build()
	return err
}

func parseLevel(lstr string) zapcore.Level {
	l := strings.TrimSpace(lstr)
	switch l {
	case "", "debug":
		return zap.DebugLevel
	case "info":
		return zap.InfoLevel
	default:
		return zap.DebugLevel
	}
}

func checkInit() {
	if logger == nil {
		logger, _ = defaultConfig.Build()
	}
}

// Info .
func Info(msg string, fields ...zapcore.Field) {
	checkInit()
	logger.Info(msg, fields...)
}

// Debug .
func Debug(msg string, fields ...zapcore.Field) {
	checkInit()
	logger.Debug(msg, fields...)
}

// Error .
func Error(msg string, fields ...zapcore.Field) {
	checkInit()
	logger.Error(msg, fields...)
}

// Warn .
func Warn(msg string, fields ...zapcore.Field) {
	checkInit()
	logger.Warn(msg, fields...)
}

// DPanic .
func DPanic(msg string, fields ...zapcore.Field) {
	checkInit()
	logger.DPanic(msg, fields...)
}

// Panic .
func Panic(msg string, fields ...zapcore.Field) {
	checkInit()
	logger.Panic(msg, fields...)
}

// SetLevel dynamic change level
func SetLevel(lstr string) {
	level := parseLevel(lstr)
	atom.SetLevel(level)
}

// Level show level
func Level() string {
	return atom.String()
}
