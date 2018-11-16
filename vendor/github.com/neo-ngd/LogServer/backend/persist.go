package backend

import (
	"os"
	"path"
	"time"

	"github.com/kataras/golog"
	"github.com/lestrrat/go-file-rotatelogs"
	"github.com/pkg/errors"
	"github.com/rifflock/lfshook"
	"github.com/sirupsen/logrus"
)

type Persist struct {
	log *logrus.Logger
}
type PlainFormatter struct {
}

func (p *PlainFormatter) Format(entry *logrus.Entry) ([]byte, error) {
	return []byte(entry.Message), nil
}
func NewPersist(logPath, logname string, maxAge, rotationTime time.Duration) *Persist {
	p := Persist{
		log: logrus.New(),
	}
	p.log.SetLevel(logrus.InfoLevel)
	p.configLocalFilesystemLogger(logPath, logname, maxAge, rotationTime)
	return &p
}
func (p *Persist) configLocalFilesystemLogger(logPath string, logFileName string, maxAge time.Duration, rotationTime time.Duration) {
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
	}, &PlainFormatter{})
	p.log.AddHook(lfHook)
}

func (p *Persist) Add(arg ...interface{}) {
	p.log.Println(arg...)
}
