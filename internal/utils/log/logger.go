package log

import (
	"os"
	"sync"

	zlog "github.com/rs/zerolog"
)

var logger zlog.Logger
var once sync.Once

func Init() {
	output := zlog.ConsoleWriter{
		Out:     os.Stdout,
		NoColor: false,
	}
	logger = zlog.New(output).With().Timestamp().
		Caller().Stack().Logger()
}

func Info() *zlog.Event {
	once.Do(Init)
	return logger.Info()
}

func Error() *zlog.Event {
	once.Do(Init)
	return logger.Error()
}

func Debug() *zlog.Event {
	once.Do(Init)
	return logger.Debug()
}

func Warn() *zlog.Event {
	once.Do(Init)
	return logger.Warn()
}

func Fatal() *zlog.Event {
	once.Do(Init)
	return logger.Fatal()
}

func Panic() *zlog.Event {
	once.Do(Init)
	return logger.Panic()
}
