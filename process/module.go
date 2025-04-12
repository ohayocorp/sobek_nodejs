package process

import (
	"os"
	"strings"

	"github.com/grafana/sobek"
	"github.com/ohayocorp/sobek_nodejs/require"
)

const ModuleName = "process"

type Process struct {
	runtime *sobek.Runtime
	env     map[string]string
	argv    []string
}

func (p *Process) cwd(call sobek.FunctionCall) sobek.Value {
	wd, err := os.Getwd()
	if err != nil {
		panic(p.runtime.ToValue(err))
	}

	return p.runtime.ToValue(wd)
}

func (p *Process) chdir(call sobek.FunctionCall) sobek.Value {
	if len(call.Arguments) == 0 {
		panic(p.runtime.ToValue("chdir requires a path argument"))
	}

	path := call.Arguments[0].String()
	if err := os.Chdir(path); err != nil {
		panic(p.runtime.ToValue(err))
	}

	return sobek.Undefined()
}

func Require(runtime *sobek.Runtime, module *sobek.Object) {
	p := &Process{
		runtime: runtime,
		env:     make(map[string]string),
	}

	for _, e := range os.Environ() {
		envKeyValue := strings.SplitN(e, "=", 2)
		p.env[envKeyValue[0]] = envKeyValue[1]
	}

	o := module.Get("exports").(*sobek.Object)
	o.Set("env", p.env)
	o.Set("argv", p.argv)
	o.Set("cwd", p.cwd)
	o.Set("chdir", p.chdir)
}

func Enable(runtime *sobek.Runtime) {
	runtime.Set("process", require.Require(runtime, ModuleName))
}

func init() {
	require.RegisterCoreModule(ModuleName, Require)
}
