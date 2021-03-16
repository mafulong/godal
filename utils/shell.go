package utils

import (
	"fmt"
	"github.com/ionrock/procs"
)

func CallCommand(cmd string) error {
	p := procs.NewProcess(cmd)
	p.OutputHandler = func(line string) string {
		return line
	}
	p.ErrHandler = func(line string) string {
		return line
	}
	return p.Run()
}

func GoFmt(file string) {
	CallCommand(fmt.Sprintf("go fmt %q", file))
}
