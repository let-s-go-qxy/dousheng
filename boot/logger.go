package boot

import (
	"fmt"
	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
	g "tiktok/app/global"
	"tiktok/utils/file"
	"time"
)

// Options logger options
type Options struct {
	SavePath     string // save path
	EncoderType  string // encoder type("json","console")
	EncodeLevel  string // encode style
	EncodeCaller string // caller type
}

const (
	// encoder type
	JsonEncoder   = "json"
	ConsoleEncode = "console"

	// encode level
	LowercaseLevelEncoder      = "LowercaseLevelEncoder"
	LowercaseColorLevelEncoder = "LowercaseColorLevelEncoder"
	CapitalLevelEncoder        = "CapitalLevelEncoder"
	CapitalColorLevelEncoder   = "CapitalColorLevelEncoder"

	// caller option
	ShortCallerEncoder = "ShortCallerEncoder"
	FullCallerEncoder  = "FullCallerEncoder"
)

func LoggerSetup() {
	options := Options{
		SavePath:     g.Config.Logger.SavePath,
		EncoderType:  g.Config.Logger.EncoderType,
		EncodeLevel:  g.Config.Logger.EncodeLevel,
		EncodeCaller: g.Config.Logger.EncodeCaller,
	}
	LoggerSetupWithOptions(options)
}

func LoggerSetupWithOptions(options Options) {
	// create log dir
	err := file.IsNotExistMkDir(options.SavePath)
	if err != nil {
		panic(err)
	}

	dynamicLevel := zap.NewAtomicLevel()
	debugPriority := zap.LevelEnablerFunc(func(lev zapcore.Level) bool {
		return lev == zap.DebugLevel
	})
	infoPriority := zap.LevelEnablerFunc(func(lev zapcore.Level) bool {
		return lev == zap.InfoLevel
	})
	warnPriority := zap.LevelEnablerFunc(func(lev zapcore.Level) bool {
		return lev == zap.WarnLevel
	})
	errorPriority := zap.LevelEnablerFunc(func(lev zapcore.Level) bool {
		return lev >= zap.ErrorLevel
	})
	encoder := getEncoder(options)
	cores := [...]zapcore.Core{
		zapcore.NewCore(encoder, os.Stdout, dynamicLevel), // console output
		// filed on level
		zapcore.NewCore(encoder, getWriteSyncer(fmt.Sprintf("./%s/all/server_all.log", options.SavePath)), zapcore.DebugLevel),
		zapcore.NewCore(encoder, getWriteSyncer(fmt.Sprintf("./%s/debug/server_debug.log", options.SavePath)), debugPriority),
		zapcore.NewCore(encoder, getWriteSyncer(fmt.Sprintf("./%s/info/server_info.log", options.SavePath)), infoPriority),
		zapcore.NewCore(encoder, getWriteSyncer(fmt.Sprintf("./%s/warn/server_warn.log", options.SavePath)), warnPriority),
		zapcore.NewCore(encoder, getWriteSyncer(fmt.Sprintf("./%s/error/server_error.log", options.SavePath)), errorPriority),
	}
	zapLogger := zap.New(zapcore.NewTee(cores[:]...), zap.AddCaller())
	defer func(zapLogger *zap.Logger) {
		_ = zapLogger.Sync()
	}(zapLogger)
	// set current log level to "Debug"
	dynamicLevel.SetLevel(zap.DebugLevel)
	// set global logger
	g.Logger = zapLogger.Sugar()
	g.Logger.Info("initialize logger successfully!")
	//sugar.Debug("test")
	//sugar.Warn("test")
	//sugar.Error("test")
	//sugar.DPanic("test")
	//sugar.Panic("test")
	//sugar.Fatal("test")
}

func getEncoder(options Options) zapcore.Encoder {
	if options.EncoderType == JsonEncoder {
		return zapcore.NewJSONEncoder(getEncoderConfig(options))
	}
	return zapcore.NewConsoleEncoder(getEncoderConfig(options))
}

func getEncoderConfig(options Options) (config zapcore.EncoderConfig) {
	config = zapcore.EncoderConfig{
		MessageKey:     "message",
		LevelKey:       "level",
		TimeKey:        "time",
		NameKey:        "logger",
		CallerKey:      "caller",
		StacktraceKey:  "stacktrace",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.LowercaseLevelEncoder,
		EncodeTime:     CustomTimeEncoder,
		EncodeDuration: zapcore.StringDurationEncoder,
		EncodeCaller:   zapcore.FullCallerEncoder, // format: short（package/filepath.go:line） full (filepath.go:line)
	}
	switch {
	case options.EncodeLevel == LowercaseLevelEncoder: // default
		config.EncodeLevel = zapcore.LowercaseLevelEncoder
	case options.EncodeLevel == LowercaseColorLevelEncoder:
		config.EncodeLevel = zapcore.LowercaseColorLevelEncoder
	case options.EncodeLevel == CapitalLevelEncoder:
		config.EncodeLevel = zapcore.CapitalLevelEncoder
	case options.EncodeLevel == CapitalColorLevelEncoder:
		config.EncodeLevel = zapcore.CapitalColorLevelEncoder
	default:
		config.EncodeLevel = zapcore.LowercaseLevelEncoder
	}
	if options.EncodeCaller == ShortCallerEncoder {
		config.EncodeCaller = zapcore.ShortCallerEncoder
	}
	return config
}

func getWriteSyncer(file string) zapcore.WriteSyncer {
	lumberJackLogger := &lumberjack.Logger{
		Filename:   file, // log file location
		MaxSize:    1,    // log file maximum size (MB)
		MaxBackups: 100,
		MaxAge:     30, // day
		Compress:   true,
	}
	return zapcore.AddSync(lumberJackLogger)
}

func CustomTimeEncoder(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString(t.Format("[2006-01-02 15:04:05.000]"))
}
