package logger

import (
	"github.com/inagacky/go_gin_api/src/api/util"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
)
const (
	logFileName = "go_sample_output.log"
)

// ロガーの設定
func Init() (*zap.Logger, error) {

	// 環境変数から取得
	logPath := util.Getenv("go_gin_api_LOG_PATH", "/Users/d-inagaki/go/src/github.com/inagacky/go_gin_api/log")

	logFileName := logPath + "/" + logFileName
	// ログファイルの作成・追記の準備
	if _, err := os.OpenFile(logFileName, os.O_WRONLY|os.O_CREATE, 0664); err != nil {
		return nil, err
	}

	level := zap.NewAtomicLevel()
	level.SetLevel(zapcore.InfoLevel)

	myConfig := zap.Config{
		Level: level,
		Encoding: "json",
		EncoderConfig: zapcore.EncoderConfig{
			TimeKey:        "Time",
			LevelKey:       "Level",
			NameKey:        "Name",
			CallerKey:      "Caller",
			MessageKey:     "Msg",
			StacktraceKey:  "St",
			EncodeLevel:    zapcore.CapitalLevelEncoder,
			EncodeTime:     zapcore.ISO8601TimeEncoder,
			EncodeDuration: zapcore.StringDurationEncoder,
			EncodeCaller:   zapcore.ShortCallerEncoder,
		},
		OutputPaths:      []string{logFileName},
		ErrorOutputPaths: []string{logFileName},
	}

	logger, lErr := myConfig.Build()
	zap.ReplaceGlobals(logger)

	defer logger.Sync()

	return logger, lErr
}

// Loggerを返却
func GetLogger() *zap.Logger {
	return zap.L()
}