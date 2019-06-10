package storage

import (
	"os"
	"path"
	"time"

	"github.com/kataras/golog"
	rotatelogs "github.com/lestrrat/go-file-rotatelogs"
	"github.com/pkg/errors"
	"github.com/rifflock/lfshook"
	"github.com/sirupsen/logrus"
)

type persist struct {
	log *logrus.Logger
}
type plainFormatter struct {
}

func (p *plainFormatter) Format(entry *logrus.Entry) ([]byte, error) {
	return []byte(entry.Message), nil
}
func newPersist(logPath, logname string, maxAge, rotationTime time.Duration) *persist {
	p := persist{
		log: logrus.New(),
	}
	p.log.SetLevel(logrus.InfoLevel)
	p.configLocalFilesystemLogger(logPath, logname, maxAge, rotationTime)
	return &p
}
func (p *persist) configLocalFilesystemLogger(logPath string, logFileName string, maxAge time.Duration, rotationTime time.Duration) {
	_, err := os.Stat(logPath)
	if err != nil {
		err = os.Mkdir(logPath, os.ModePerm)
		if err != nil {
			golog.Fatal(err)
		}
	}
	baseLogPaht := path.Join(logPath, logFileName)
	writer, err := rotatelogs.New(
		baseLogPaht+".%Y%m%d%H%M%S",
		rotatelogs.WithLinkName(baseLogPaht),
		rotatelogs.WithMaxAge(maxAge),
		rotatelogs.WithRotationTime(rotationTime),
	)
	if err != nil {
		golog.Errorf("config local file system logger error. %+v", errors.WithStack(err))
	}
	lfHook := lfshook.NewHook(lfshook.WriterMap{
		logrus.DebugLevel: writer,
		logrus.InfoLevel:  writer,
		logrus.WarnLevel:  writer,
		logrus.ErrorLevel: writer,
		logrus.FatalLevel: writer,
		logrus.PanicLevel: writer,
	}, &plainFormatter{})
	p.log.AddHook(lfHook)
}

func (p *persist) append(arg ...interface{}) {
	p.log.Println(arg...)
}
