package config

//config
type Config struct {
	LogStdOut bool
}

var config = &Config{
	LogStdOut: false,
}

//init
func Init() {
	logInit()
}

func init() {
	Init()
}
