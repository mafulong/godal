package utils

import (
	"io"
	"io/ioutil"
	"os"
	"path"
	"strings"
)

func IsExist(p string) bool {
	_, err := os.Stat(p)
	return os.IsExist(err)
}

func IsDirExists(dir string) bool {
	stat, err := os.Stat(dir)
	if err != nil {
		return false
	}
	return stat.IsDir()
}
func IsFileExists(path string) bool {
	_, err := os.Stat(path) //os.Stat获取文件信息
	if err != nil {
		if os.IsExist(err) {
			return true
		}
		return false
	}
	return true
}

func RemoveDirectory(dir string) error {
	err := os.RemoveAll(dir)
	if err != nil {
		return err
	} else {
		return nil
	}
}

func IsDirectory(p string) (bool, error) {
	stat, err := os.Stat(p)
	if err != nil {
		return false, err
	}
	return stat.IsDir(), nil
}

func CopyFile(from string, to string) error {

	fromFd, err := os.Open(from)
	if err != nil {
		return err
	}
	defer func() {
		if err := fromFd.Close(); err != nil {
		}
	}()

	toFd, err := os.Create(to)
	if err != nil {
		return err
	}
	defer func() {
		if err := toFd.Close(); err != nil {
		}
	}()

	if _, err = io.Copy(toFd, fromFd); err != nil {
		return err
	}

	fromInfo, err := os.Stat(from)
	if err != nil {
		return err
	}
	return os.Chmod(to, fromInfo.Mode())
}

func CopyDirectory(from string, to string) error {
	fromInfo, err := os.Stat(from)
	if err != nil {
		return err
	}

	if err = os.MkdirAll(to, fromInfo.Mode()); err != nil {
		return err
	}

	fds, err := ioutil.ReadDir(from)
	if err != nil {
		return err
	}
	for _, fd := range fds {
		f := path.Join(from, fd.Name())
		t := path.Join(to, fd.Name())

		if fd.IsDir() {
			if err = CopyDirectory(f, t); err != nil {
				return err
			}
		} else {
			if err = CopyFile(f, t); err != nil {
				return err
			}
		}
	}
	return nil
}
func WriteToFile(filename string, data []byte) (err error) {
	if strings.HasSuffix(filename, ".sh") {
		err = ioutil.WriteFile(filename, data, 0755)
	} else {
		err = ioutil.WriteFile(filename, data, 0644)
	}
	return
}

func ReadFile(file string) (dataBytes []byte, err error) {
	content, err := ioutil.ReadFile(file)
	if err != nil {
		return nil, err
	}
	return content, nil
}

func GetPWD() (string) {
	wd, _:= os.Getwd()
	return wd
}
