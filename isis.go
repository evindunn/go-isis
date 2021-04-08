package isis

import (
	"bytes"
	"errors"
	"fmt"
	"os/exec"
	"reflect"
	"strings"
)

func Isis(cmd string, args map[string]string) error {
	argsArr := make([]string, 0)

	for argKey, argVal := range args {
		arg := fmt.Sprintf("%s=%s", argKey, argVal)
		argsArr = append(argsArr, arg)
	}

	var stderr bytes.Buffer
	isisCommandPath, err := exec.LookPath(cmd)
	if err != nil {
		return errors.New(fmt.Sprintf("%s is not in PATH", cmd))
	}

	isisCommand := exec.Command(strings.TrimSpace(isisCommandPath), argsArr[:]...)
	isisCommand.Stderr = &stderr
	// fmt.Println(isisCommand.String())

	err = isisCommand.Run()
	if err != nil {
		if reflect.TypeOf(err).String() == fmt.Sprintf("%T", err) {
			return errors.New(fmt.Sprintf("%s failed: %s", cmd, stderr.String()))
		} else {
			return err
		}
	}

	return nil
}
