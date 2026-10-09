package metric

import (
	"embed"
	"text/template"
)

//go:embed templates/*.html
var templatesFS embed.FS

var metricsListHTML = template.Must(
	template.ParseFS(templatesFS, "templates/*.html"),
)
