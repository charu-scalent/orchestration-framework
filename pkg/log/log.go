package log

import (
	"fmt"

	zlog "github.com/rs/zerolog/log"
)

type LoggerConfig struct {
	Filename   string
	MaxSize    int
	MaxBackups int
	MaxAge     int
	LogLevel   int
}

func NewLogger(logger *LoggerConfig) *LoggerConfig {
	return &LoggerConfig{
		Filename:   logger.Filename,
		MaxSize:    logger.MaxSize,
		MaxBackups: logger.MaxBackups,
		MaxAge:     logger.MaxAge,
		LogLevel:   logger.LogLevel,
	}
}

func Info(msg string, requestID string) {
	zlog.Info().Str("RequestID", requestID).Caller(1).Msg(msg) //usage : logging.Info("msg", "reqID")
}

func InfoWithData(msg string, data interface{}, requestID string) {
	zlog.Info().Str("RequestID", requestID).Caller(1).Msg(msg + fmt.Sprintf(" %+v", data)) //usage : logging.Info("msg", "reqID")
}

func Error(msg string, requestID string) {
	zlog.Error().Str("RequestID", requestID).Caller(1).Msg(msg) //usage : logging.Error("msg", "reqID")
}

func ErrorWithData(msg string, data interface{}, requestID string) {
	zlog.Error().Str("RequestID", requestID).Caller(1).Msg(msg + fmt.Sprintf(" %+v", data)) //usage : logging.Error("msg",data, "reqID")
}

func Debug(msg string, requestID string) {
	zlog.Debug().Str("RequestID", requestID).Caller(1).Msg(msg) //usage : logging.Debug("msg", "reqID")
}

func DebugWithData(msg string, data interface{}, requestID string) {
	zlog.Debug().Str("RequestID", requestID).Caller(1).Msg(msg + fmt.Sprintf(" %+v", data)) //usage : logging.Debug("msg", "reqID")
}

func Print(v ...interface{}) {
	zlog.Print(v...) //usage : logging.Print("val1", "val2")
}

func Printf(format string, v ...interface{}) {
	zlog.Printf(format, v...)
}

func Debugf(format string, v ...interface{}) {
	zlog.Debug().Msgf(format, v...)
}
