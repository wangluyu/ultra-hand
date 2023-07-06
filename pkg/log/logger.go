package log

import (
	"github.com/google/wire"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
	"os"
)

var (
	defaultLoggerName = "Default"
)

type Logger struct {
	name   string
	logger *zap.SugaredLogger
}

type Option struct {
	LoggerName   string
	LogPath      string
	ErrLogPath   string
	LevelEnabler string
	Stdout       bool
	Json         bool
	MaxSize      int
	MaxBackups   int
	MaxAge       int
	Compress     bool
}

func NewLoggerOption(v *viper.Viper) (*Option, error) {
	o := new(Option)
	err := v.UnmarshalKey("log", o)
	if err != nil {
		return &Option{}, err
	}
	return o, nil
}

func NewZapLogger(o *Option) (*zap.SugaredLogger, error) {
	cores := make([]zapcore.Core, 0, 3)
	encoder := getEncoder(o.Json)
	levelEnabler, err := getLevelEnabler(o.LevelEnabler)
	if err != nil {
		return nil, err
	}
	cores = append(cores, zapcore.NewCore(
		encoder,
		getLogWriter(o.LogPath, o),
		levelEnabler,
	))
	if "" != o.ErrLogPath {
		cores = append(cores, zapcore.NewCore(
			encoder,
			getLogWriter(o.ErrLogPath, o),
			zapcore.ErrorLevel,
		))
	}
	if o.Stdout {
		cores = append(cores, zapcore.NewCore(
			getEncoder(false),
			zapcore.Lock(os.Stdout),
			levelEnabler,
		))
	}
	coreTee := zapcore.NewTee(cores...)
	logger := zap.New(coreTee, zap.AddCaller(), zap.AddCallerSkip(1))
	defer logger.Sync() // flushes buffer, if any
	sugarLogger := logger.Sugar()
	return sugarLogger, nil
}

func getLevelEnabler(level string) (zap.LevelEnablerFunc, error) {
	var levelEnabler zap.LevelEnablerFunc
	parseLevel, err := zapcore.ParseLevel(level)
	if err != nil {
		return nil, err
	}
	switch level {
	case "debug":
		levelEnabler = zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
			return lvl >= zapcore.DebugLevel
		})
	default:
		levelEnabler = zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
			return lvl <= parseLevel && lvl != zapcore.DebugLevel
		})
	}
	return levelEnabler, nil
}

func getEncoder(isJson bool) zapcore.Encoder {
	ec := zap.NewProductionEncoderConfig()
	ec.EncodeTime = zapcore.ISO8601TimeEncoder
	ec.EncodeLevel = zapcore.CapitalLevelEncoder
	if isJson {
		return zapcore.NewJSONEncoder(ec)
	} else {
		return zapcore.NewConsoleEncoder(ec)
	}
}

func getLogWriter(lp string, o *Option) zapcore.WriteSyncer {
	return zapcore.AddSync(&lumberjack.Logger{
		Filename:   lp,
		MaxSize:    o.MaxSize,
		MaxAge:     o.MaxAge,
		MaxBackups: o.MaxBackups,
		Compress:   o.Compress,
	})
}

func New(o *Option, logger *zap.SugaredLogger) (Logger, error) {
	name := o.LoggerName
	if name == "" {
		name = defaultLoggerName
	}
	return Logger{logger: logger, name: name}, nil
}

func (l Logger) Debug(msg string, keysAndValues ...interface{}) {
	l.logger.Debugf(msg, keysAndValues...)
}

func (l Logger) Info(msg string, keysAndValues ...interface{}) {
	l.logger.Infow(msg, keysAndValues...)
}

func (l Logger) Warn(msg string, keysAndValues ...interface{}) {
	l.logger.Warnf(msg, keysAndValues...)
}

func (l Logger) Error(msg string, keysAndValues ...interface{}) {
	l.logger.Errorf(msg, keysAndValues...)
}

func (l Logger) With(args ...interface{}) Logger {
	l.logger = l.logger.With(args...)
	return l
}

var ProvideSet = wire.NewSet(New, NewZapLogger, NewLoggerOption)
