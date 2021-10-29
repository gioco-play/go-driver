package logrusz

import (
	"fmt"
	"github.com/rifflock/lfshook"
	"github.com/sirupsen/logrus"
)

type Logger struct {
	Log    *logrus.Logger
	Level  string
	Prefix string
	Path   string
}

type Formatter struct {
	logrus.TextFormatter
}

func (f *Formatter) Format(entry *logrus.Entry) ([]byte, error) {
	_, e := f.TextFormatter.Format(entry)

	str := fmt.Sprintf("[%s]: %s - %s File:%s\n", entry.Level, entry.Time.Format(f.TimestampFormat), entry.Message, entry.Caller.File)

	return []byte(str), e
	// return append([]byte("prefix: "), l...), e
}

func New() *Logger {
	return &Logger{
		Log:    logrus.New(),
		Level:  "error",
		Prefix: "",
		Path:   "./logs",
	}
}

func (l *Logger) SetPrefix(prefix string) *Logger {
	l.Prefix = prefix
	return l
}

func (l *Logger) SetPath(path string) *Logger {
	l.Path = path
	return l
}

func (l *Logger) SetLevel(level string) *Logger {
	l.Level = level
	return l
}

func (l *Logger) Writer() *logrus.Logger {

	logger := l.Log

	lv, err := logrus.ParseLevel(l.Level)
	if err != nil {
		panic(err)
	}

	logger.SetLevel(lv)
	logger.SetReportCaller(true)

	pathMap := lfshook.PathMap{
		logrus.DebugLevel: fmt.Sprintf("%s/%s%s", l.Path, l.Prefix, "debug.log"),
		logrus.InfoLevel:  fmt.Sprintf("%s/%s%s", l.Path, l.Prefix, "info.log"),
		logrus.WarnLevel:  fmt.Sprintf("%s/%s%s", l.Path, l.Prefix, "warn.log"),
		logrus.ErrorLevel: fmt.Sprintf("%s/%s%s", l.Path, l.Prefix, "error.log"),
	}

	logger.AddHook(lfshook.NewHook(pathMap, &Formatter{
		logrus.TextFormatter{
			TimestampFormat: "2006-01-02 15:04:05",
			DisableColors:   true,
		},
	}))

	return logger
}
