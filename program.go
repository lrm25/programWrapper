package programWrapper

import "os/exec"

type Program struct {
	executable            string
	params                map[string]interface{}
	GetCombinedOutputFunc func(cmd *exec.Cmd) ([]byte, error)
}

func NewProgram(executable string) *Program {
	program := &Program{
		executable: executable,
		params:     make(map[string]interface{}),
	}
	program.GetCombinedOutputFunc = program.GetCombinedOutput
	return program
}

func (p *Program) GetCombinedOutput(cmd *exec.Cmd) ([]byte, error) {
	return cmd.CombinedOutput()
}

func (p *Program) Run() (string, error) {

	cmd := exec.Command(p.executable)
	output, err := p.GetCombinedOutputFunc(cmd)
	if err != nil {
		return "", err
	}
	return string(output), nil
}
