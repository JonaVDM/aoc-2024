package inputs

import (
	"embed"
	_ "embed"
)

//go:embed *
var Inputs embed.FS
