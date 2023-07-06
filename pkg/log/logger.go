package log

import (
	"github.com/google/wire"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
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
	file, err := os.Create(o.LogPath)
	if err != nil {
		return nil, err
	}
	core := zapcore.NewCore(
		zapcore.NewJSONEncoder(zap.NewProductionEncoderConfig()),
		zapcore.AddSync(file),
		zapcore.DebugLevel,
	)
	logger := zap.New(core)
	defer logger.Sync() // flushes buffer, if any
	sugarLogger := logger.Sugar()
	return sugarLogger, nil
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
