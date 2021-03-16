package config

import (
	"github.com/mafulong/godal/utils"
	log "github.com/sirupsen/logrus"
	"testing"
)

func TestLog(t *testing.T) {
	logInit()
	log.Debug("wahhh")
	log.Info(utils.ToJSON(config))
}
