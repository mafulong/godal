package model

import (
	"bytes"
	"fmt"
	"github.com/mafulong/godal/utils"
	"github.com/xwb1989/sqlparser"
	"io/ioutil"
	"path"
	"strings"
	"text/template"
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
		fp := path.Join(utils.GetPWD(), fmt.Sprintf("/model/%+v.go", ddl.NewName.Name))
		err := utils.WriteToFile(fp, []byte(genTable("model", ddl)))
		if err != nil {
			return err
		}
		fmt.Println(fp)
		utils.GoFmt(fp)
	}
	return nil
}

func genTable(pkg string, ddl *sqlparser.DDL) string {
	var imports string
	if needTime(ddl) {
		imports = `import "time"` + "\n"
	}

	tableNameStr := ddl.NewName.Name.String()
	tableName := utils.ToCamelFirstLower(tableNameStr)

	var columns strings.Builder
	for i, c := range genColumns(ddl) {
		if i != 0 {
			columns.WriteString("\n")
		}
		columns.WriteString("\t")
		columns.WriteString(c)
	}

	params := struct {
		Package      string
		Imports      string
		TableName    string
		TableNameStr string
		Columns      string
	}{
		Package:      pkg,
		Imports:      imports,
		TableName:    tableName,
		TableNameStr: tableNameStr,
		Columns:      columns.String(),
	}

	var buf bytes.Buffer
	template.Must(template.New("header").Parse(tableTemplate)).Execute(&buf, params)

	return buf.String()
}

func needTime(ddl *sqlparser.DDL) bool {
	for _, c := range ddl.TableSpec.Columns {
		switch c.Type.Type {
		case "date", "datetime", "timestamp":
			return true
		}
	}
	return false
}

func genColumns(ddl *sqlparser.DDL) []string {
	columns := make([]string, 0, len(ddl.TableSpec.Columns))
	for _, c := range ddl.TableSpec.Columns {
		columns = append(columns, utils.GenColumn(c))
	}
	return columns
}
