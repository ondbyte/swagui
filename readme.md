## swagui - a golang package to serve that 'petstore.json'
or your own swagger spec (json/yaml)

### example
```go
package main

import (
	_ "embed"
	"fmt"
	"net/http"

	"github.com/ondbyte/swagui"
)

//go:embed petstore.json
var spec []byte

// you can also do spec,err := os.ReadFile("petstore.json")

func main() {
	http.HandleFunc("/swagger_doc/*", swagui.Handle(spec, swagui.Json))
	fmt.Println("see your swagger documention on http://localhost:8080/swagger_doc")
	http.ListenAndServe(":8080", http.DefaultServeMux)
}

```