package logger

import (
	"fmt"
	"log"
	"net/http"
	"time"

	apachelog "github.com/lestrrat-go/apache-logformat"
	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
)

func (l L) TestOfRotatedLog() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("string"))
	})

	logf, err := rotatelogs.New(
		"./logs/access_log.%Y%m%d%H%M",
		rotatelogs.WithLinkName("./logs/access_log"),
		rotatelogs.WithMaxAge(5*time.Minute),
		rotatelogs.WithRotationSize(1024),
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
