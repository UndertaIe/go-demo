package logger

import (
	"fmt"
	"log"

	"github.com/sirupsen/logrus"
	"gopkg.in/natefinch/lumberjack.v2"
)

var rollog = "./logs/lumberjack.log"

func (l L) DescOfLumberjack() {
	fmt.Println("Lumberjack is log rolling package")
}

func (l L) TestOfLumberjack() {
	log.SetOutput(&lumberjack.Logger{
		Filename:   rollog,
		MaxSize:    1, // megabytes
		MaxBackups: 5,
		MaxAge:     28,   //days
		Compress:   true, // disabled by default
		LocalTime:  true,
	})
	log.Println("rolling log msg")
}
func (l L) LumberjackForLog() {
	log.SetOutput(&lumberjack.Logger{
		Filename:   rollog,
		MaxSize:    1, // megabytes
		MaxBackups: 5,
		MaxAge:     28,   //days
		Compress:   true, // disabled by default
		LocalTime:  true,
	})
	for i := 0; i < 100000; i++ {
		log.Printf("rolling log msg: %d \n", i)
	}
}

func (l L) LumberjackForLogrus() {
	out := &lumberjack.Logger{
		Filename:   rollog,
		MaxSize:    2, // megabytes
		MaxBackups: 4,
		MaxAge:     30,    //days
		Compress:   false, // disabled by default
		LocalTime:  true,
	}
	nlog := logrus.New()
	nlog.SetOutput(out)
	nlog.SetFormatter(&logrus.JSONFormatter{PrettyPrint: true})
	for i := 0; i < 100000; i++ {
		nlog.Infof("rolling msg for logrus: %d \n", i)
	}
}






