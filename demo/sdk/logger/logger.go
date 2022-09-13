package logger

import (
	"log"
	"os"
	"path/filepath"
)

type L struct{}

func (l L) Desc() string {
	return "日志示例包括: logrus, lumberjack, rotatedlogs"
}

func (l L) Name() string {
	return "logger"
}

func (l L) DefaultOfLogger() {
	log := log.Default()
	log.Println("std log msg")
}

var filename = "./logs/log.log"

func (l L) NewOfLogger() {
	dirname := filepath.Dir(filename)
	var err error
	if err = os.MkdirAll(dirname, 0755); err != nil {
		log.Panicf("failed to create directory %s\n", dirname)
	}
	fh, err := os.OpenFile(filename, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		log.Panicf("failed to open file %s\n", dirname)
	}

	clog := log.New(fh, "[Custom Logger]", log.LstdFlags)
	clog.Println("println msg")
	clog.SetOutput(os.Stdout) // 切换标准输出流
	clog.Println("println msg")
}
