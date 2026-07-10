
project (
	name {{ .Project.Name }}
	version {{ .Project.Version }}
	module {{ .Project.Module }}
	author {{ .Project.Author }}
)

build main (
	input "./cmd/{{ .Project.Name }}/{{ .Project.Name }}.go"
	output "./dist/{{ .Project.Name }}"
	trimpath true
)
