package model

import (
	"fmt"
	"io/ioutil"
	"path"
	"testing"

	"github.com/mafulong/go_utils"
	"github.com/mafulong/godal/config"
	log "github.com/sirupsen/logrus"
)

func TestGenName(t *testing.T) {
	t.Log(go_utils.GetPWD())
	testFile := "../../test/gen_model.sql"
	content, err := ioutil.ReadFile(testFile)
	if err != nil {
		t.Fatal(err)
		return
	}
	ddls, err := go_utils.ParseSQLs(string(content))
	if err != nil {
		t.Fatal(err)
		return
	}
	for _, ddl := range ddls {
		fp := path.Join(go_utils.GetPWD(), fmt.Sprintf("/model/%+v.go", ddl.NewName.Name))
		fmt.Println(fp)
	}
	t.Log(go_utils.ToJSON(ddls))
}

func Test_gen(t *testing.T) {
	config.Init()
	log.Info("test")
	//t.Log(path.Join(go_utils.GetPWD(), "../../test/gen_model.sql"))
	err := gen(path.Join(go_utils.GetPWD(), "../../test/gen_model.sql"))
	if err != nil {
		t.Fatal(err)
		return
	}
	if !go_utils.IsFileExists(path.Join(go_utils.GetPWD(), "/model/test_tb1.go")) || !go_utils.IsFileExists(path.Join(go_utils.GetPWD(), "/model/test_tb2.go")) {
		t.Fatal("no file generated")
		return
	}
	go_utils.RemoveDirectory("model")
}
