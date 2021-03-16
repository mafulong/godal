package config

import (
	"fmt"
	nested "github.com/antonfisher/nested-logrus-formatter"
	log "github.com/sirupsen/logrus"
	"github.com/zput/zxcTool/ztLog/zt_formatter"
	"os"
	"path"
	"runtime"
)

func logInit() {
	// Log as JSON instead of the default ASCII formatter.
	formatter := &zt_formatter.ZtFormatter{
		CallerPrettyfier: func(f *runtime.Frame) (string, string) {
			filename := path.Base(f.File)
			return fmt.Sprintf("%s()", f.Function), fmt.Sprintf("%s:%d", filename, f.Line)
			//return "", fmt.Sprintf("%s:%d", filename, f.Line)
		},
		Formatter: nested.Formatter{
			//HideKeys: true,
			FieldsOrder: []string{"component", "category"},
		},
	}
	log.SetFormatter(formatter)
	log.SetReportCaller(true)

	// Only log the warning severity or above.
	log.SetLevel(log.DebugLevel)

	// Output to stdout instead of the default stderr
	// Can be any io.Writer, see below for File example
	//log.SetOutput(os.Stdout)
	stdout := config.LogStdOut
	if stdout {
		log.SetOutput(os.Stdout)
	} else {
		file, err := os.OpenFile("/tmp/godal.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
		if err == nil {
			log.SetOutput(file)
		} else {
			log.Info("Failed to log to file, using default stderr")
		}
	}
}
