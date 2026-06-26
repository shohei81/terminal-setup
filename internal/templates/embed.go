package templates

import _ "embed"

//go:embed ghostty.tmpl
var GhosttyTmpl string

//go:embed starship.toml
var StarshipToml string

//go:embed zshrc.tmpl
var ZshrcTmpl string
