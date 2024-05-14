package zaplogger

import (
	"fmt"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
	"os"
)

type Config struct {
	Level string      `yaml:"level" env:"LOGGER_LEVEL" env-default:"info"`
	File  *FileConfig `yaml:"file"`
}

type FileConfig struct {
	Filename   string `yaml:"filename"  env:"LOGGER_FILE_NAME"`
	MaxSize    int    `yaml:"maxsize" env:"LOGGER_FILE_MAXSIZE"`
	MaxAge     int    `yaml:"maxage" env:"LOGGER_FILE_MAXAGE"`
	MaxBackups int    `yaml:"maxbackups" env:"LOGGER__FILE_MAXBACKUPS"`
	LocalTime  bool   `yaml:"localtime" env:"LOGGER_FILE_LOCALTIME"`
	Compress   bool   `yaml:"compress" env:"LOGGER_FILE_COMPRESS"`
}

func (f *FileConfig) toLumberjack() *lumberjack.Logger {
	return &lumberjack.Logger{
		Filename:   f.Filename,
		MaxSize:    f.MaxSize,
		MaxAge:     f.MaxAge,
		MaxBackups: f.MaxBackups,
		LocalTime:  f.LocalTime,
		Compress:   f.Compress,
	}
}

func NewZapLogger(cfg Config) (*zap.Logger, error) {
	level, err := zap.ParseAtomicLevel(cfg.Level)
	if err != nil {
		return nil, fmt.Errorf("parsing level: %w", err)
	}
	encoderCfg := zap.NewProductionEncoderConfig()
	encoderCfg.TimeKey = "timestamp"
	encoderCfg.EncodeTime = zapcore.ISO8601TimeEncoder

	jsonEncoder := zapcore.NewJSONEncoder(encoderCfg)

	var cores []zapcore.Core

	stdout := zapcore.AddSync(os.Stdout)
	cores = append(cores, zapcore.NewCore(jsonEncoder, stdout, level))

	if cfg.File != nil {
		file := zapcore.AddSync(cfg.File.toLumberjack())
		cores = append(cores, zapcore.NewCore(jsonEncoder, file, level))
	}

	core := zapcore.NewTee(cores...)
	return zap.New(core), nil
}

func ReplaceZap(cfg Config) (func(), error) {
	logger, err := NewZapLogger(cfg)
	if err != nil {
		return nil, err
	}
	return zap.ReplaceGlobals(logger.WithOptions(zap.AddCaller())), nil
}
