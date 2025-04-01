Nodejs compatibility library for sobek
====

This is a collection of sobek modules that provide nodejs compatibility.

Example:

```go
package main

import (
    "github.com/grafana/sobek"
    "github.com/ohayocorp/sobek_nodejs/require"
)

func main() {
    registry := new(require.Registry) // this can be shared by multiple runtimes

    runtime := sobek.New()
    req := registry.Enable(runtime)

    runtime.RunString(`
    var m = require("./m.js");
    m.test();
    `)

    m, err := req.Require("./m.js")
    _, _ = m, err
}
```

More modules will be added. Contributions welcome too.
