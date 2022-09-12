package logger

type L struct{}

func (l L) Desc() string {
	return "日志示例包括: logrus, lumberjack, rotatedlogs"
}

func (l L) Name() string {
	return "logger"
}
