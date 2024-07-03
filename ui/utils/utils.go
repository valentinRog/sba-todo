package utils

import (
	g "github.com/maragudk/gomponents"
	"github.com/maragudk/gomponents/html"
)

func AddClass(className string, f func(...g.Node) g.Node) func(...g.Node) g.Node {
	return func(children ...g.Node) g.Node {
		children = append(children, html.Class(className))
		return f(children...)
	}
}
