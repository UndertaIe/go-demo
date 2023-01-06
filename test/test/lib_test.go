package lib

import "testing"

// 测试文件名必须以”_test.go”结尾；
// 测试函数名必须以“TestXxx”开始；
// 命令行下使用”gotest”即可启动测试；

func AppendTest(t *testing.T) {
	s1 := "1"
	s2 := "2"
	expect := "12"
	if expect != Append(s1, s2) {
		t.Errorf("Append(\"1\", \"2\") != \"12\"")
	}
}

func AddTest(t *testing.T) {
	i1 := 1
	i2 := 2
	expect := 2
	if expect != Add(i1, i2) {
		t.Errorf("Add(1, 2) != 12")
	}
}
