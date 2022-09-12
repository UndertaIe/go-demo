package logger

import (
	"context"
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/sirupsen/logrus"
)

func (l L) TestLogrus() {
	log := logrus.WithFields(logrus.Fields{
		"animal": "walrus",
	})
	log.Info("A walrus appears")
}

// 默认格式、标准输出流
func (l L) NewOfLogrus() {
	log := logrus.New()
	log.Trace("Trace msg")
	log.Debug("Debug msg") // 不会输出，默认日志等级为info,  info大于debug级别则不会输出
	log.Info("Info msg")
	log.Warn("Warn msg")
	log.Error("Error msg")
	log.Fatal("Fatal msg")
	log.Panic("Panic msg") // 触发panic强制结束
}

func (l L) StandardLoggerOfLogrus() {
	log := logrus.StandardLogger() // 返回默认全局变量std
	fmt.Println("default log level:", log.Level.String())
	fmt.Printf("default Formatter: %T\n", log.Formatter)
	fmt.Printf("default out: %T\n", log.Out) // 默认标准输出流
}

type hook int

func (h hook) Levels() []logrus.Level {
	return []logrus.Level{logrus.InfoLevel, logrus.WarnLevel}
}

func (h hook) Fire(*logrus.Entry) error {
	fmt.Print("hook msg!  ===>>>  ")
	return nil
}

func (l L) HookOfLogrus() {
	log := logrus.StandardLogger() // 返回默认全局变量std
	log.AddHook(new(hook))
	log.Info("Info msg")
	log.Warn("Warn msg")
	log.Error("Error msg") // 在Levels方法中没有定义Error级别的回调方法
	log.Debug("Debug msg")

}

func (l L) CustomLogrus() {
	hooks := map[logrus.Level][]logrus.Hook{logrus.InfoLevel: {new(hook)}}
	nLog := logrus.Logger{
		Out:          os.Stdout,
		Formatter:    new(logrus.TextFormatter),
		Hooks:        hooks,
		Level:        logrus.DebugLevel,
		ExitFunc:     os.Exit,
		ReportCaller: true, // 打印调用文件、行号及函数
	}
	nLog.Debug("debug msg at debug level")

	nLog.ReportCaller = false
	nLog.Info("Info msg at debug level")

	nLog.Out = os.Stderr
	nLog.Info("Info msg at debug level")

	nLog.Formatter = new(logrus.JSONFormatter)
	nLog.Info("Info msg at debug level")
}

func (l L) EntryOfLogrus() {
	log := logrus.New()
	entry := log.WithContext(context.Background())
	entry.Info("Info msg")
	entry.Infof("Info msg: %d", 5)
	entry.Infoln("Info msg ")

	entry.Logger.Log(logrus.InfoLevel, "info msg", " msg2 ")
	entry.Logger.Logf(logrus.InfoLevel, "info msg: %s", "msg2 ")
	cb := func() []interface{} {
		return []interface{}{"msg1, ", "msg2"}
	}

	entry.Logger.Level = logrus.DebugLevel
	entry.Logger.LogFn(logrus.DebugLevel, cb)
}

func (l L) WithOfLogrus() {
	log := logrus.New()
	entry := log.WithTime(time.Now())
	entry.Info("WithTime msg")
	fmt.Println(entry.Time.Format("2006-01-02 15:04:05"))

	e2 := log.WithError(errors.New("some error"))
	e2.Error("error msg")
	e2.Info("error msg")

	e3 := log.WithField("k1", "v1")
	e3.Error("WithField msg")

	e4 := log.WithFields(map[string]any{"api": "/api/v1", "time": time.Now()})
	e4.Info("WithFields msg")
}
