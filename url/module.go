package url

import (
	"github.com/grafana/sobek"
	"github.com/ohayocorp/sobek_nodejs/require"
)

const ModuleName = "url"

type urlModule struct {
	r *sobek.Runtime

	URLSearchParamsPrototype         *sobek.Object
	URLSearchParamsIteratorPrototype *sobek.Object
}

func Require(runtime *sobek.Runtime, module *sobek.Object) {
	exports := module.Get("exports").(*sobek.Object)
	m := &urlModule{
		r: runtime,
	}
	exports.Set("URL", m.createURLConstructor())
	exports.Set("URLSearchParams", m.createURLSearchParamsConstructor())
	exports.Set("domainToASCII", m.domainToASCII)
	exports.Set("domainToUnicode", m.domainToUnicode)
}

func Enable(runtime *sobek.Runtime) {
	m := require.Require(runtime, ModuleName).ToObject(runtime)
	runtime.Set("URL", m.Get("URL"))
	runtime.Set("URLSearchParams", m.Get("URLSearchParams"))
}

func init() {
	require.RegisterCoreModule(ModuleName, Require)
}
