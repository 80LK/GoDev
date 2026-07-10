module {{ .Project.Module }}

go {{ .GoVersion }}

tool github.com/80LK/godev/cmd/god
