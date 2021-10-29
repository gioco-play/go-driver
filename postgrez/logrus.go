package postgrez

import (
	"errors"
	"github.com/sirupsen/logrus"
	"golang.org/x/net/context"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/utils"
	"time"
)

type loggerus struct {
	Logger                *logrus.Logger
	SlowThreshold         time.Duration
	SourceField           string
	SkipErrRecordNotFound bool
}

func NewLogger(log *logrus.Logger) *loggerus {
	return &loggerus{
		Logger:                log,
		SkipErrRecordNotFound: true,
	}
}

func (l *loggerus) LogMode(logger.LogLevel) logger.Interface {
	return l
}

func (l *loggerus) Info(ctx context.Context, s string, args ...interface{}) {
	l.Logger.WithContext(ctx).Infof(s, args)
}

func (l *loggerus) Warn(ctx context.Context, s string, args ...interface{}) {
	l.Logger.WithContext(ctx).Warnf(s, args)
}

func (l *loggerus) Error(ctx context.Context, s string, args ...interface{}) {
	l.Logger.WithContext(ctx).Errorf(s, args)
}

func (l *loggerus) Trace(ctx context.Context, begin time.Time, fc func() (string, int64), err error) {
	elapsed := time.Since(begin)
	sql, _ := fc()
	fields := logrus.Fields{}
	if l.SourceField != "" {
		fields[l.SourceField] = utils.FileWithLineNum()
	}
	if err != nil && !(errors.Is(err, gorm.ErrRecordNotFound) && l.SkipErrRecordNotFound) {
		fields[logrus.ErrorKey] = err
		l.Logger.WithContext(ctx).WithFields(fields).Errorf("%s [%s]", sql, elapsed)
		return
	}

	if l.SlowThreshold != 0 && elapsed > l.SlowThreshold {
		l.Logger.WithContext(ctx).WithFields(fields).Warnf("%s [%s]", sql, elapsed)
		return
	}

	l.Logger.WithContext(ctx).WithFields(fields).Debugf("%s [%s]", sql, elapsed)
}
