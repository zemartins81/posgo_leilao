package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var (
    log *zap.Logger
)

func init() {
    logConfiguration := zap.Config{
        Level: zap.NewAtomicLevelAt(zap.DebugLevel),
        Development: false,
        DisableCaller: false,
        DisableStacktrace: false,
        Sampling: nil,
        Encoding: "json",
        EncoderConfig: zapcore.EncoderConfig{
            MessageKey: "message",
            LevelKey: "level",
            TimeKey: "time",
            EncodeLevel: zapcore.LowercaseColorLevelEncoder,
            EncodeTime: zapcore.ISO8601TimeEncoder,
            EncodeCaller: zapcore.ShortCallerEncoder,
        },
    }

    log, _ = logConfiguration.Build()
}
