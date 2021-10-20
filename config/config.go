package config

// Config eg logOut
type Config struct {
	LogStdOut bool
}

var config = &Config{
	LogStdOut: false,
}

func init() {
	logInit()
}
