package utils

import (
	"testing"
)

func TestParseSQLs(t *testing.T) {
	type args struct {
		content string
	}
	type want struct {
		tableName string
	}
	tests := []struct {
		name    string
		args    args
		want    want
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			args:    args{content: "CREATE TABLE IF NOT EXISTS `test_tb1`(\n   `test_id` INT UNSIGNED AUTO_INCREMENT,\n   `test_title` VARCHAR(100) NOT NULL,\n   `test_author` VARCHAR(40) NOT NULL,\n   `submission_date` DATE,\n   PRIMARY KEY ( `test_id` )\n)ENGINE=InnoDB DEFAULT CHARSET=utf8;"},
			wantErr: false,
			want:    want{tableName: "test_tb1"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ParseSQLs(tt.args.content)
			if (err != nil) != tt.wantErr {
				t.Errorf("ParseSQLs() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if len(got) != 1 {
				t.Error("parse failed")
			}
			if got[0].NewName.Name.String() != tt.want.tableName {
				t.Error("parse failed")
			}
		})
	}
}


func TestColumn(t *testing.T) {
}

