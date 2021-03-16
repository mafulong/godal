package model

import (
	"fmt"
	"github.com/mafulong/godal/config"
	"github.com/mafulong/godal/utils"
	log "github.com/sirupsen/logrus"
	"io/ioutil"
	"path"
	"testing"
)

func TestGenName(t *testing.T) {
	t.Log(utils.GetPWD())
	testFile := "../../test/gen_model.sql"
	content, err := ioutil.ReadFile(testFile)
	if err != nil {
		t.Fatal(err)
		return
	}
	ddls, err := utils.ParseSQLs(string(content))
	if err != nil {
		t.Fatal(err)
		return
	}
	for _, ddl := range ddls {
		fp := path.Join(utils.GetPWD(), fmt.Sprintf("/model/%+v.go", ddl.NewName.Name))
		fmt.Println(fp)
	}
	t.Log(utils.ToJSON(ddls))
}

func Test_gen(t *testing.T) {
	config.Init()
	log.Info("test")
	//t.Log(path.Join(utils.GetPWD(), "../../test/gen_model.sql"))
	err := gen(path.Join(utils.GetPWD(), "../../test/gen_model.sql"))
	if err != nil {
		t.Fatal(err)
		return
	}
	if !utils.IsFileExists(path.Join(utils.GetPWD(), "/model/test_tb1.go")) || !utils.IsFileExists(path.Join(utils.GetPWD(), "/model/test_tb2.go")) {
		t.Fatal("no file generated")
		return
	}
	utils.RemoveDirectory("model")
}
