package ui

import (
	_ "embed"
)

//go:embed dist/style.css
var CssString string

//go:embed node_modules/htmx.org/dist/htmx.min.js
var HtmxString string
