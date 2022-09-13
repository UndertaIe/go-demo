package logger

import (
	"fmt"
	"log"
	"net/http"
	"time"

	apachelog "github.com/lestrrat-go/apache-logformat"
	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"github.com/sirupsen/logrus"
)

func (l L) DescOfRotatedLog() {
	fmt.Println("file-rotatelogs is log rotate package")
}

var RotatedLink = "./logs/rotatelogs"
var RotatedFn = "./logs/rotatelogs.%Y%m%d%H%M"

func (l L) ReadMeOfRotatedLog() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("string"))
	})

	logf, err := rotatelogs.New(
		RotatedFn,
		rotatelogs.WithLinkName(RotatedLink),
		rotatelogs.WithMaxAge(5*time.Minute),
		rotatelogs.WithRotationSize(1024*1024*5),
		rotatelogs.WithRotationTime(10*time.Second),
	)
	if err != nil {
		log.Printf("failed to create rotatelogs: %s", err)
		return
	}

	// Now you must write to logf. apache-logformat library can create
	// a http.Handler that only writes the approriate logs for the request
	// to the given handle
	if err2 := http.ListenAndServe(":8080", apachelog.CombinedLog.Wrap(mux, logf)); err2 != nil {
		fmt.Println(err2)
	}
}

var RotatedLinkForLog = "./logs/rotatelogs_in_log"
var RotatedFnForLog = "./logs/rotatelogs_in_log.%Y%m%d%H%M"

func (l L) RotatedLogForLog() {
	out, err := rotatelogs.New(
		RotatedFnForLog,
		rotatelogs.WithLinkName(RotatedLinkForLog),
		rotatelogs.WithMaxAge(5*time.Minute),
		rotatelogs.WithRotationSize(1024*1024*2),
		rotatelogs.WithRotationTime(10*time.Second),
	)
	if err != nil {
		log.Printf("failed to create rotatelogs: %s", err)
		return
	}
	log.Default().SetOutput(out)

	for i := 0; i < 100000; i++ {
		log.Printf("rotated msg for log: %d \n", i)
	}
}

var RotatedLinkForLogrus = "./logs/rotatelogs_in_logrus"
var RotatedFnForLogrus = "./logs/rotatelogs_in_logrus.%Y%m%d%H%M"

func (l L) RotatedLogForLogrus() {
	out, err := rotatelogs.New(
		RotatedFnForLogrus,
		rotatelogs.WithLinkName(RotatedLinkForLogrus),
		rotatelogs.WithMaxAge(5*time.Minute),
		rotatelogs.WithRotationSize(1024*1024*2),
		rotatelogs.WithRotationTime(10*time.Second),
	)
	if err != nil {
		log.Printf("failed to create rotatelogs: %s", err)
		return
	}
	nlog := logrus.New()
	nlog.SetOutput(out)
	nlog.SetFormatter(&logrus.JSONFormatter{})
	for i := 0; i < 100000; i++ {
		nlog.Infof("rotated msg for logrus: %d \n", i)
	}
}
