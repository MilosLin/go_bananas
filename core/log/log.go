package log

import (
	"os"
	"path"

	"github.com/MilosLin/go_bananas/core/config"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

var logs = make(map[string]*zap.Logger)

// Configuration for logging
type LogConfig struct {
	// EncodeLogsAsJson makes the log framework log JSON
	EncodeLogsAsJson bool
	// FileLoggingEnabled makes the framework log to a file
	// the fields below can be skipped if this value is false!
	FileLoggingEnabled bool
	// Directory to log to to when filelogging is enabled
	Directory string
	// Filename is the name of the logfile which will be placed inside the directory
	Filename string
	// MaxSize the max size in MB of the logfile before it's rolled
	MaxSize int
	// MaxBackups the max number of rolled files to keep
	MaxBackups int
	// MaxAge the max age in days to keep a logfile
	MaxAge int
}

// How to log, by example:
// logger.Info("Importing new file, zap.String("source", filename), zap.Int("size", 1024))
// To log a stacktrace:
// logger.Error("It went wrong, zap.Stack())

// DefaultZapLogger is the default logger instance that should be used to log
// It's assigned a default value here for tetypests (which do not call log.Configure())
var DefaultZapLogger = newZapLogger(false, os.Stdout)

// Debug Log a message at the debug level. Messages include any context that's
// accumulated on the logger, as well as any fields added at the log site.
//
// Use zap.String(key, value), zap.Int(key, value) to log fields. These fields
// will be marshalled as JSON in the logfile and key value pairs in the console!
func Debug(msg string, fields ...zapcore.Field) {
	DefaultZapLogger.Debug(msg, fields...)
}

// Info log a message at the info level. Messages include any context that's
// accumulated on the logger, as well as any fields added at the log site.
//
// Use zap.String(key, value), zap.Int(key, value) to log fields. These fields
// will be marshalled as JSON in the logfile and key value pairs in the console!
func Info(msg string, fields ...zapcore.Field) {
	DefaultZapLogger.Info(msg, fields...)
}

// Warn log a message at the warn level. Messages include any context that's
// accumulated on the logger, as well as any fields added at the log site.
//
// Use zap.String(key, value), zap.Int(key, value) to log fields. These fields
// will be marshalled as JSON in the logfile and key value pairs in the console!
func Warn(msg string, fields ...zapcore.Field) {
	DefaultZapLogger.Warn(msg, fields...)
}

// Error Log a message at the error level. Messages include any context that's
// accumulated on the logger, as well as any fields added at the log site.
//
// Use zap.String(key, value), zap.Int(key, value) to log fields. These fields
// will be marshalled as JSON in the logfile and key value pairs in the console!
func Error(msg string, fields ...zapcore.Field) {
	DefaultZapLogger.Error(msg, fields...)
}

// Panic Log a message at the Panic level. Messages include any context that's
// accumulated on the logger, as well as any fields added at the log site.
//
// Use zap.String(key, value), zap.Int(key, value) to log fields. These fields
// will be marshalled as JSON in the logfile and key value pairs in the console!
func Panic(msg string, fields ...zapcore.Field) {
	DefaultZapLogger.Panic(msg, fields...)
}

// Fatal Log a message at the fatal level. Messages include any context that's
// accumulated on the logger, as well as any fields added at the log site.
//
// Use zap.String(key, value), zap.Int(key, value) to log fields. These fields
// will be marshalled as JSON in the logfile and key value pairs in the console!
func Fatal(msg string, fields ...zapcore.Field) {
	DefaultZapLogger.Fatal(msg, fields...)
}

// AtLevel logs the message at a specific log level
func AtLevel(level zapcore.Level, msg string, fields ...zapcore.Field) {
	switch level {
	case zapcore.DebugLevel:
		Debug(msg, fields...)
	case zapcore.PanicLevel:
		Panic(msg, fields...)
	case zapcore.ErrorLevel:
		Error(msg, fields...)
	case zapcore.WarnLevel:
		Warn(msg, fields...)
	case zapcore.InfoLevel:
		Info(msg, fields...)
	case zapcore.FatalLevel:
		Fatal(msg, fields...)
	default:
		Warn("Logging at unkown level", zap.Any("level", level))
		Warn(msg, fields...)
	}
}

func getConfig(log_name string) *LogConfig {
	c := config.Instance()
	l := LogConfig{
		EncodeLogsAsJson: c.GetBool("log.root.encode_logs_as_json"),
		// FileLoggingEnabled makes the framework log to a file
		// the fields below can be skipped if this value is false!
		FileLoggingEnabled: c.GetBool("log.root.file_logging_enabled"),
		// Directory to log to to when filelogging is enabled
		Directory: c.GetString("log.root.directory"),
		// Filename is the name of the logfile which will be placed inside the directory
		Filename: c.GetString("log.root.filename"),
		// MaxSize the max size in MB of the logfile before it's rolled
		MaxSize: c.GetInt("log.root.max_size"),
		// MaxBackups the max number of rolled files to keep
		MaxBackups: c.GetInt("log.root.max_backups"),
		// MaxAge the max age in days to keep a logfile
		MaxAge: c.GetInt("log.root.max_age"),
	}
	return &l
}

func Inst(log_name string) *zap.Logger {
	if _, ok := logs[log_name]; !ok {
		logs[log_name] = n(log_name)
	}
	return logs[log_name]
}

// Configure sets up the logging framework
//
// In production, the container logs will be collected and file logging should be disabled. However,
// during development it's nicer to see logs as text and optionally write to a file when debugging
// problems in the containerized pipeline
//
// The output log file will be located at /var/log/auth-service/auth-service.log and
// will be rolled when it reaches 20MB with a maximum of 1 backup.
func n(log_name string) *zap.Logger {
	config := getConfig(log_name)
	writers := []zapcore.WriteSyncer{os.Stdout}
	if config.FileLoggingEnabled {
		writers = append(writers, newRollingFile(config))
	}

	Default := newZapLogger(config.EncodeLogsAsJson, zapcore.NewMultiWriteSyncer(writers...))
	zap.RedirectStdLog(Default)
	Info("logging configured",
		zap.Bool("fileLogging", config.FileLoggingEnabled),
		zap.Bool("jsonLogOutput", config.EncodeLogsAsJson),
		zap.String("logDirectory", config.Directory),
		zap.String("fileName", config.Filename),
		zap.Int("maxSizeMB", config.MaxSize),
		zap.Int("maxBackups", config.MaxBackups),
		zap.Int("maxAgeInDays", config.MaxAge))
	return Default
}

func newRollingFile(config *LogConfig) zapcore.WriteSyncer {
	if err := os.MkdirAll(config.Directory, 0); err != nil {
		Error("failed create log directory", zap.Error(err), zap.String("path", config.Directory))
		return nil
	}

	return zapcore.AddSync(&lumberjack.Logger{
		Filename:   path.Join(config.Directory, config.Filename),
		MaxSize:    config.MaxSize,    //megabytes
		MaxAge:     config.MaxAge,     //days
		MaxBackups: config.MaxBackups, //files
	})
}

func newZapLogger(encodeAsJSON bool, output zapcore.WriteSyncer) *zap.Logger {
	encCfg := zapcore.EncoderConfig{
		TimeKey:        "@timestamp",
		LevelKey:       "level",
		NameKey:        "logger",
		CallerKey:      "caller",
		MessageKey:     "msg",
		StacktraceKey:  "stacktrace",
		EncodeLevel:    zapcore.LowercaseLevelEncoder,
		EncodeTime:     zapcore.ISO8601TimeEncoder,
		EncodeDuration: zapcore.NanosDurationEncoder,
	}

	encoder := zapcore.NewConsoleEncoder(encCfg)
	if encodeAsJSON {
		encoder = zapcore.NewJSONEncoder(encCfg)
	}

	return zap.New(zapcore.NewCore(encoder, output, zap.NewAtomicLevel()))
}
