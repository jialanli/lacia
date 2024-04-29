package lacia

import (
	"bytes"
	"fmt"
	"os/exec"
)

func ExecCommand(name string, args ...string) (outString string, err error) {
	var out bytes.Buffer
	var stderr bytes.Buffer

	cmd := exec.Command(name, args...)
	cmd.Stdout = &out
	cmd.Stderr = &stderr

	err = cmd.Run()
	outString = out.String()
	if err != nil {
		err = fmt.Errorf("exec failed: %v, stderr=%s, name=%s, args=%v", err, stderr.String(), name, args)
	}

	return
}
