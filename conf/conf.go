package conf

import _ "embed"

//go:embed version.txt

var CodeVersion string
