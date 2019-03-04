# Changelog - v{{ .Version }}
#### Released {{ .Date }}
{{ if gt (len .BreakingChanges) 0}}
## Breaking Changes
{{ range .BreakingChanges }}
    {{ . }}
{{ end }}
{{ end }}
{{ if gt (len .Features) 0}}
## Features
{{ range .Features }}
{{ . }}
{{ end }}
{{ end }}
{{ if gt (len .Fixes) 0}}
## Fixes / Minor Updates
{{ range .Fixes }}
{{ . }}
{{ end }}
{{ end }}
## Release Contributors
{{ range .Contributors }}
{{ . }}
{{ end }}