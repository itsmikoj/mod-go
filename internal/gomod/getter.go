package gomod

import (
	"os/exec"

	"github.com/itsmikoj/mod-go/pkg/utils"
)

func Get(pkg string, version ...string) error {
	cmd := exec.Command("go", "get", pkg)
	return utils.RunCommand(cmd)
}
