# mux
Caddy plugin for routing requests to Go http handlers

The actual development of this plugin is in vendor dir of another project.


## Deployment 

Caddy requires modifying its source code to recognize new plugin. Fortunately it's a single line change.

Add "mux" to the end of var "directives" array in https://github.com/mholt/caddy/blob/master/caddyhttp/httpserver/plugin.go

Also declare you want to use the plugin in Caddyfile:

```
localhost:2015
mux
```


## Usage Example

Embed Caddy in a project and route requests to be handled by Go functions. It uses Go http ServeMux below but may be replaced by gorilla mux.


```
package main

import (
    "net/http"

    "github.com/mholt/caddy/caddy/caddymain"
    "github.com/seckiss/mux"
)

func main() {
    mux.HandleFunc("/test", TestHandler)
    caddymain.Run()
}

func TestHandler(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "text/plain")
    w.Write([]byte("test handler success"))
}
```
