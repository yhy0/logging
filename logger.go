package logging

import (
	"fmt"
	"github.com/rifflock/lfshook"
	"os"

	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"github.com/sirupsen/logrus"
)

var stdFormatter *TextFormatter  // 命令行输出格式
var fileFormatter *TextFormatter // 文件输出格式

var Logger *logrus.Logger

func New(debug bool, logPath string, logFileName string, screenReport bool) *logrus.Logger {
	log := logrus.New()
	log.SetReportCaller(true)
	stdFormatter = &TextFormatter{
		FullTimestamp:   true,
		TimestampFormat: "2006-01-02 15:04:05",
		ForceFormatting: true,
		ForceColors:     true,
		DisableColors:   false,
		ReportCaller:    screenReport,
	}
	fileFormatter = &TextFormatter{
		FullTimestamp:   true,
		TimestampFormat: "2006-01-02 15:04:05",
		ForceFormatting: true,
		ForceColors:     false,
		DisableColors:   true,
		ReportCaller:    true,
	}

	log.SetFormatter(stdFormatter)
	if debug {
		log.SetLevel(logrus.DebugLevel)
	} else {
		log.SetLevel(logrus.InfoLevel)
	}
	if logPath == "" {
		logPath, _ = os.Getwd()
	}

	logName := fmt.Sprintf("%s/logs/%s_", logPath, logFileName)
	writer, _ := rotatelogs.New(logName + "%Y_%m_%d" + ".log")
	lfHook := lfshook.NewHook(lfshook.WriterMap{
		logrus.InfoLevel:  writer,
		logrus.WarnLevel:  writer,
		logrus.DebugLevel: writer,
		logrus.ErrorLevel: writer,
	}, fileFormatter)
	log.SetOutput(os.Stdout)
	log.AddHook(lfHook)
	return log
}
