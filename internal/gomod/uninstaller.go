package gomod

import (
	"os/exec"

	"github.com/itsmikoj/mod-go/pkg/utils"
)

func Uninstall(pkg string) error {
	cmd := exec.Command("py", "clean_imports.py", "pkg="+pkg)
	return utils.RunCommand(cmd)
}
