package gomod

import (
    "os/exec"
    "github.com/itsmikoj/mod-go/pkg/utils"
)

func Tidy() error {
	cmd := exec.Command("go", "mod", "tidy")
	return utils.RunCommand(cmd)
}