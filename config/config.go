package config

// Config eg logOut
type Config struct {
	LogStdOut bool
}

var config = &Config{
	LogStdOut: false,
}

// Init config init
func Init() {
	logInit()
}

func init() {
	Init()
}
