package config

import "testing"

func Test_logInit(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			t.Fatal("panic")
		}
	}()
	logInit()
}
