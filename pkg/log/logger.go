package log

import (
	"github.com/google/wire"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"io"
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
	LoggerName string
	LogPath    string
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
	core := zapcore.NewCore(
		getEncoder(),
		getLogWriter(o.LogPath),
		zapcore.DebugLevel,
	)
	logger := zap.New(core, zap.AddCaller(), zap.AddCallerSkip(1))
	defer logger.Sync() // flushes buffer, if any
	sugarLogger := logger.Sugar()
	return sugarLogger, nil
}

func getEncoder() zapcore.Encoder {
	ec := zap.NewProductionEncoderConfig()
	ec.EncodeTime = zapcore.ISO8601TimeEncoder
	ec.EncodeLevel = zapcore.CapitalLevelEncoder
	return zapcore.NewConsoleEncoder(ec)
}

func getLogWriter(lp string) zapcore.WriteSyncer {
	file, err := os.Create(lp)
	if err != nil {
		return nil
	}
	ws := io.MultiWriter(file, os.Stdout)
	return zapcore.AddSync(ws)
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
