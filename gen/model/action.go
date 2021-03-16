package model

import (
	"fmt"
	"github.com/mafulong/godal/utils"
	"io/ioutil"
	"path"
)

const tableTemplate = `
package {{.Package}}

{{.Imports}}

type {{.TableName}} struct {
{{.Columns}}
}

func ({{.TableName}}) TableName() string {
	return "{{.TableNameStr}}"
}
`

func gen(file string) error {
	content, err := ioutil.ReadFile(file)
	if err != nil {
		return err
	}
	ddls, err := utils.ParseSQLs(string(content))
	if err != nil {
		return err
	}
	for _, ddl := range ddls {
		//fp := getModelFilePath(ddl)
		fp := path.Join(utils.GetPWD(), fmt.Sprintf("/model/%+v.go", ddl.NewName))
		fmt.Println(fp)
		//util.CallCommand(fmt.Sprintf("go fmt %q", fp))
	}
	return nil
}
