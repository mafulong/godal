package main

import (
	"github.com/mafulong/go_utils"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_godal(t *testing.T) {
	err := run([]string{"", "gen", "model", "--database", "test_tb1", "./test/gen_model.sql"})
	assert.Nil(t, err)
	isExist := go_utils.IsFileExists("./model/test_tb1/test_tb1.go")
	assert.Equal(t, true, isExist)
	go_utils.RemoveDirectory("./model")
}
