package utils

import (
	"fmt"

	g "github.com/maragudk/gomponents"
	"github.com/maragudk/gomponents/html"
)

func AddClass(className string, f func(...g.Node) g.Node) func(...g.Node) g.Node {
	return func(children ...g.Node) g.Node {
		children = append(children, html.Class(className))
		return f(children...)
	}
}

func HxHeader(id string) g.Node {
	return g.Attr(
		"hx-headers",
		fmt.Sprintf("{\"component-id\": \"%s\"}", id),
	)
}
