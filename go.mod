module github.com/UndertaIe/go-demo

go 1.18

require (
	github.com/getsentry/sentry-go v0.13.0
	github.com/go-redis/redis/v8 v8.11.5
	github.com/gomodule/redigo v1.8.9
	github.com/ilam01/limits-go v1.0.0
	github.com/lestrrat-go/apache-logformat v2.0.4+incompatible
	github.com/lestrrat-go/file-rotatelogs v2.4.0+incompatible
	github.com/ozgio/strutil v0.4.0
	github.com/sirupsen/logrus v1.9.0
	github.com/spf13/cobra v1.5.0
	go.uber.org/zap v1.23.0
	gopkg.in/natefinch/lumberjack.v2 v2.0.0
	gorm.io/driver/mysql v1.3.6
	gorm.io/driver/sqlite v1.3.2
	gorm.io/gorm v1.23.8
	gorm.io/plugin/opentelemetry v0.1.0
)

require (
	github.com/BurntSushi/toml v1.2.0 // indirect
	github.com/cespare/xxhash/v2 v2.1.2 // indirect
	github.com/dgryski/go-rendezvous v0.0.0-20200823014737-9f7001d12a5f // indirect
	github.com/facebookgo/clock v0.0.0-20150410010913-600d898af40a // indirect
	github.com/go-sql-driver/mysql v1.6.0 // indirect
	github.com/inconshreveable/mousetrap v1.0.0 // indirect
	github.com/jinzhu/inflection v1.0.0 // indirect
	github.com/jinzhu/now v1.1.5 // indirect
	github.com/jonboulle/clockwork v0.3.0 // indirect
	github.com/lestrrat-go/strftime v1.0.6 // indirect
	github.com/mattn/go-sqlite3 v1.14.12 // indirect
	github.com/onsi/gomega v1.20.2 // indirect
	github.com/pkg/errors v0.9.1 // indirect
	github.com/spf13/pflag v1.0.5 // indirect
	go.opentelemetry.io/otel v1.8.0 // indirect
	go.opentelemetry.io/otel/trace v1.8.0 // indirect
	go.uber.org/atomic v1.7.0 // indirect
	go.uber.org/multierr v1.8.0 // indirect
	golang.org/x/sys v0.0.0-20220722155257-8c9f86f7a55f // indirect
)

replace github.com/ozgio/strutil v0.4.0 => github.com/UndertaIe/strutil v0.0.0-20220922025843-22719a4bf7fb
