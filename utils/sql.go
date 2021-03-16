package utils

import (
	"fmt"
	"github.com/xwb1989/sqlparser"
)

func ParseSQLs(content string) ([]*sqlparser.DDL, error) {
	pieces, err := sqlparser.SplitStatementToPieces(content)
	if err != nil {
		return nil, err
	}
	ddls := make([]*sqlparser.DDL, 0, len(pieces))
	for _, piece := range pieces {
		stmt, err := sqlparser.Parse(piece)
		if err != nil {
			continue
		}
		switch stmt.(type) {
		case *sqlparser.DDL:
			ddl := stmt.(*sqlparser.DDL)
			if ddl.Action != "create" {
				continue
			}
			if ddl.TableSpec == nil {
				continue
			}
			ddls = append(ddls, ddl)
		}
	}
	return ddls, nil
}

type Column struct {
	Name    string
	Type    string
	Comment string
}

func (c Column) String() string {
	s := fmt.Sprintf("%s %s `gorm:\"Column:%s\" json:\"%s\"`",
		ToFirstUpperCamel(c.Name), c.Type, c.Name, c.Name)
	if c.Comment == "" {
		return s
	} else {
		return s + "// " + c.Comment
	}
}

func GenColumn(c *sqlparser.ColumnDefinition) string {
	switch c.Type.Type {
	case "bigint":
		return Column{c.Name.String(), "int64", getComment(c)}.String()
	case "int", "smallint", "tinyint":
		return Column{c.Name.String(), "int", getComment(c)}.String()
	case "char", "varchar", "text", "mediumtext", "longtext":
		return Column{c.Name.String(), "string", getComment(c)}.String()
	case "blob":
		return Column{c.Name.String(), "[]byte", getComment(c)}.String()
	case "float", "double", "decimal":
		return Column{c.Name.String(), "float64", getComment(c)}.String()
	case "bit":
		return Column{c.Name.String(), "uint64", getComment(c)}.String()
	case "date", "datetime", "timestamp":
		return Column{c.Name.String(), "time.Time", getComment(c)}.String()
	default:
		panic(fmt.Sprintf("bad Column: %+v", c))
	}
}

func getComment(c *sqlparser.ColumnDefinition) string {
	if c.Type.Comment == nil {
		return ""
	} else {
		return string(c.Type.Comment.Val)
	}
}