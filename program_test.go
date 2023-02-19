package wrapper

import (
	"errors"
	"os/exec"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_program_cmdError(t *testing.T) {

	cmdErr := errors.New("test error")

	program := NewProgram("")
	getCombinedOutputFunc = func(cmd *exec.Cmd) ([]byte, error) {
		return nil, cmdErr
	}
	defer func() { getCombinedOutputFunc = getCombinedOutput }()
	_, err := program.Run()
	require.Equal(t, cmdErr, err)
}
