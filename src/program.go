package programWrapper

import "os/exec"

type Program struct {
	executable string
	params     map[string]interface{}
}

func NewProgram(executable string) *Program {
	return &Program{
		executable: executable,
		params:     make(map[string]interface{}),
	}
}

func (p *Program) Run() (string, error) {
	cmd := exec.Command(p.executable)
	output, err := cmd.CombinedOutput()
	if err != nil {
		return "", err
	}
	return string(output), nil
}
