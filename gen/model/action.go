package model

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"path"
	"strings"
	"text/template"

	"github.com/mafulong/go_utils"
	log "github.com/sirupsen/logrus"
	"github.com/xwb1989/sqlparser"
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
	log.Info("gen file: ", file)
	content, err := ioutil.ReadFile(file)
	if err != nil {
		log.Error(err)
		return err
	}
	ddls, err := go_utils.ParseSQLs(string(content))
	if err != nil {
		return err
	}
	pkg := "model"
	if len(cmdArgs.Database) > 0 {
		pkg = go_utils.ToSnakeCase(cmdArgs.Database)
	}
	for _, ddl := range ddls {
		fp := getFilePath(ddl.NewName.Name.String())
		log.Debug(fp)
		err := go_utils.WriteToFile(fp, []byte(genTable(pkg, ddl)))
		if err != nil {
			log.Error(err)
			return err
		}
		go_utils.GoFmt(fp)
	}
	return nil
}

func getFilePath(tableName string) string {
	if len(cmdArgs.Database) > 0 {
		return path.Join(go_utils.GetPWD(), fmt.Sprintf("/model/%+v/%+v.go", cmdArgs.Database, tableName))
	}
	return path.Join(go_utils.GetPWD(), fmt.Sprintf("/model/%+v.go", tableName))
}

func genTable(pkg string, ddl *sqlparser.DDL) string {
	var imports string
	if needTimeImport(ddl) {
		imports = `import "time"` + "\n"
	}

	tableNameStr := ddl.NewName.Name.String()
	tableName := go_utils.ToCamelFirstUpper(tableNameStr)

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
	_ = template.Must(template.New("header").Parse(tableTemplate)).Execute(&buf, params)

	return buf.String()
}

func needTimeImport(ddl *sqlparser.DDL) bool {
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
		columns = append(columns, go_utils.GenColumn(c))
	}
	return columns
}
