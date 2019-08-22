# http/cors module

Module provides a middleware to wrap any **http.Handler** instance with CORS middleware that automatically presets customizable headers.

**Usage:**

```go
import (
    "net/http"

    "github.com/KWRI/demo-service/core/http/cors"
)

func process(h http.Handler, cfg Configer) {
    // Configer must implement cors.OptionsGetter
    h := cors.WrapHTTPHandler(h, cfg)
}
```
