package config

type Config struct {
	LogStdOut bool
}

var config = &Config{
	LogStdOut: false,
}

func Init()  {
	logInit()
}

func init() {
	Init()
}