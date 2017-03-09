package logger

import (
	"os"
	"path"

	"github.com/MilosLin/go_bananas/core/config"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

// Configuration for logging
type LogConfig struct {
	// Logger Level
	Level string
	// Print log on console
	PrintLog bool
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

var logs = make(map[string]*zap.Logger)

func init() {
	Forge("default")
}

// How to log, by example:
// log.Info("Importing new file, zap.String("source", filename), zap.Int("size", 1024))
// To log a stacktrace:
// logger.Error("It went wrong, zap.Stack())

// Debug Log a message at the debug level. Messages include any context that's
// accumulated on the logger, as well as any fields added at the log site.
//
// Use zap.String(key, value), zap.Int(key, value) to log fields. These fields
// will be marshalled as JSON in the logfile and key value pairs in the console!
func Debug(msg string, fields ...zapcore.Field) {
	logs["default"].Debug(msg, fields...)
}

// Info log a message at the info level. Messages include any context that's
// accumulated on the logger, as well as any fields added at the log site.
//
// Use zap.String(key, value), zap.Int(key, value) to log fields. These fields
// will be marshalled as JSON in the logfile and key value pairs in the console!
func Info(msg string, fields ...zapcore.Field) {
	logs["default"].Info(msg, fields...)
}

// Warn log a message at the warn level. Messages include any context that's
// accumulated on the logger, as well as any fields added at the log site.
//
// Use zap.String(key, value), zap.Int(key, value) to log fields. These fields
// will be marshalled as JSON in the logfile and key value pairs in the console!
func Warn(msg string, fields ...zapcore.Field) {
	logs["default"].Warn(msg, fields...)
}

// Error Log a message at the error level. Messages include any context that's
// accumulated on the logger, as well as any fields added at the log site.
//
// Use zap.String(key, value), zap.Int(key, value) to log fields. These fields
// will be marshalled as JSON in the logfile and key value pairs in the console!
func Error(msg string, fields ...zapcore.Field) {
	logs["default"].Error(msg, fields...)
}

// Panic Log a message at the Panic level. Messages include any context that's
// accumulated on the logger, as well as any fields added at the log site.
//
// Use zap.String(key, value), zap.Int(key, value) to log fields. These fields
// will be marshalled as JSON in the logfile and key value pairs in the console!
func Panic(msg string, fields ...zapcore.Field) {
	logs["default"].Panic(msg, fields...)
}

// Fatal Log a message at the fatal level. Messages include any context that's
// accumulated on the logger, as well as any fields added at the log site.
//
// Use zap.String(key, value), zap.Int(key, value) to log fields. These fields
// will be marshalled as JSON in the logfile and key value pairs in the console!
func Fatal(msg string, fields ...zapcore.Field) {
	logs["default"].Fatal(msg, fields...)
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

// forge a logger instance
//
// 依照log name取得log實例，logName需符合設定值
func Forge(logName string) *zap.Logger {
	if _, ok := logs[logName]; !ok {
		logs[logName] = newLogger(logName)
	}
	return logs[logName]
}

// loadConfig from config file
//
// 讀取本地config檔
func loadConfig(logName string) *LogConfig {
	logBaseConfigPath := "log." + logName + "."
	c := config.Instance()
	log := LogConfig{
		Level:              c.GetString(logBaseConfigPath + "level"),
		PrintLog:           c.GetBool(logBaseConfigPath + "printLog"),
		EncodeLogsAsJson:   c.GetBool(logBaseConfigPath + "encodeLogsAsJson"),
		FileLoggingEnabled: c.GetBool(logBaseConfigPath + "fileLoggingEnabled"),
		Directory:          c.GetString(logBaseConfigPath + "directory"),
		Filename:           c.GetString(logBaseConfigPath + "filename"),
		MaxSize:            c.GetInt(logBaseConfigPath + "maxSize"),
		MaxBackups:         c.GetInt(logBaseConfigPath + "maxBackups"),
		MaxAge:             c.GetInt(logBaseConfigPath + "maxAge"),
	}
	return &log
}

// new logger by specific log name
//
// 依照設定值產生log實例
func newLogger(logName string) *zap.Logger {
	config := loadConfig(logName)
	writers := []zapcore.WriteSyncer{}

	if config.PrintLog {
		writers = append(writers, os.Stdout)
	}

	if config.FileLoggingEnabled {
		writers = append(writers, newRollingFile(config))
	}

	zapLogger := newZapLogger(
		config.EncodeLogsAsJson,
		newZapLevel(config.Level),
		zapcore.NewMultiWriteSyncer(writers...),
	)
	zap.RedirectStdLog(zapLogger)
	zapLogger.Info("logging configured",
		zap.String("level", config.Level),
		zap.Bool("fileLogging", config.FileLoggingEnabled),
		zap.Bool("jsonLogOutput", config.EncodeLogsAsJson),
		zap.String("logDirectory", config.Directory),
		zap.String("fileName", config.Filename),
		zap.Int("maxSizeMB", config.MaxSize),
		zap.Int("maxBackups", config.MaxBackups),
		zap.Int("maxAgeInDays", config.MaxAge))
	return zapLogger
}

// new zap level
//
// more detail see https://godoc.org/go.uber.org/zap#Stack
func newZapLevel(levelSetting string) (level zap.AtomicLevel) {
	level = zap.NewAtomicLevel()
	switch levelSetting {
	case "debug": //debug level
		level.SetLevel(zap.DebugLevel)
	case "info": //info level
		level.SetLevel(zap.InfoLevel)
	case "warn": //warn level
		level.SetLevel(zap.WarnLevel)
	case "error": //error level
		level.SetLevel(zap.ErrorLevel)
	case "dpanic": // Using in development. logger panics after writing the message.
		level.SetLevel(zap.DPanicLevel)
	case "panic": // PanicLevel logs a message, then panics.
		level.SetLevel(zap.PanicLevel)
	case "fatal": // FatalLevel logs a message, then calls os.Exit(1).
		level.SetLevel(zap.FatalLevel)
	default:
		level.SetLevel(zap.InfoLevel)
	}
	return
}

// new Rolling File log
//
// 若 FileLoggingEnable為true，則套用寫檔設定
func newRollingFile(config *LogConfig) zapcore.WriteSyncer {
	return zapcore.AddSync(&lumberjack.Logger{
		Filename:   path.Join(config.Directory, config.Filename),
		MaxSize:    config.MaxSize,    //megabytes
		MaxAge:     config.MaxAge,     //days
		MaxBackups: config.MaxBackups, //files
	})
}

// new Zap Logger
//
// 產生zap logger實例
func newZapLogger(encodeAsJSON bool, level zap.AtomicLevel, output zapcore.WriteSyncer) *zap.Logger {
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

	return zap.New(zapcore.NewCore(encoder, output, level))
}
