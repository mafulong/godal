package model

import (
	"fmt"
	"github.com/mafulong/godal/utils"
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
	t.Log(path.Join(utils.GetPWD(), "../../test/gen_model.sql"))
	err := gen("../../test/gen_model.sql")
	if err != nil {
		t.Fatal(err)
		return
	}
	utils.RemoveDirectory("model")
}