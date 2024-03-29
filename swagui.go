package swagui

import (
	_ "embed"
	"net/http"
	"strings"
)

//go:embed swagger-ui.html
var swaggerUiHtml []byte

type SpecType string

var (
	Json SpecType = "application/json"
	Yaml SpecType = "application/yaml"
)

// mount this handler at a documentation endpoint, make sure its a wild card '*' endpoint
//
//	func main(){
//		http.Handle("/swagger_doc/*",swagui.Handle(spec,swagui.Json))
//	}
//
// spec is the swagger file data, specType is whether its json or yaml
func Handle(swaggerSpec []byte, specType SpecType) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if strings.HasSuffix(r.URL.Path, "swagger_spec") {
			w.Header().Set("Content-Type", string(specType))
			w.Write(swaggerSpec)
			return
		}
		w.Header().Set("Content-Type", "text/html")
		w.Write(swaggerUiHtml)
	}
}
