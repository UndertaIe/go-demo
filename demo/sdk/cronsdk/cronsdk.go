package cronsdk

type C struct{}

func (c C) Name() string {
	return "cron"
}

func (c C) Desc() string {
	return "定时任务"
}

func (c C) Docs() {
}