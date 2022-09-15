package gorm

type G struct{}

func (g G) Name() string {
	return "gorm"
}

func (g G) Desc() string {
	return "gorm 示例"
}
