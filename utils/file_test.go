package utils

import (
	"testing"
)

func TestGetPWD(t *testing.T) {
	CreateDir("d1")
	_ = IsDirExists("d1")
	var err error
	_, err = IsDirectory("d1")
	if err != nil {
		t.Fatal(err)
		return
	}
	err = CopyDirectory("d1", "d2")
	if err != nil {
		t.Fatal(err)
		return
	}
	err = WriteToFile("a", []byte("data"))
	if err != nil {
		t.Fatal(err)
		return
	}
	_ = IsFileExists("a")
	err = CopyFile("a", "b")
	if err != nil {
		t.Fatal(err)
		return
	}
	_, err = ReadFile("b")
	err = CallCommand("rm a b")
	if err != nil {
		t.Fatal(err)
		return
	}
	err = RemoveDirectory("d1")
	if err != nil {
		t.Fatal(err)
		return
	}
	err = RemoveDirectory("d2")
	if err != nil {
		t.Fatal(err)
		return
	}
	if err != nil {
		t.Fatal(err)
		return
	}
}
