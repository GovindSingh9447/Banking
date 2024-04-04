package logger

import (
    "go.uber.org/zap"
    "go.uber.org/zap/zapcore"
)

var log *zap.Logger

func init() {
    // Initialize logger configuration
    config := zap.NewProductionConfig()
    
    // Modify encoder configuration
    config.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
    config.EncoderConfig.TimeKey = "timestamp"
	config.EncoderConfig.StacktraceKey=""
    
    // Build the logger
    var err error
    log, err = config.Build(zap.AddCallerSkip(1))
    if err != nil {
        panic(err)
    }
}

// Info logs an informational message.
func Info(message string, fields ...zap.Field) {
    log.Info(message, fields...)
}

// Debug logs an informational message.
func Debug(message string, fields ...zap.Field) {
    log.Debug(message, fields...)
}

// Error logs an informational message.
func Error(message string, fields ...zap.Field) {
    log.Error(message, fields...)
}
