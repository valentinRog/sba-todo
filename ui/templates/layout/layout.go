package layout

import (
	g "github.com/maragudk/gomponents"
	c "github.com/maragudk/gomponents/components"
	h "github.com/maragudk/gomponents/html"
	"github.com/valentinRog/spa-todo/ui/utils"

	_ "embed"
)

var (
	Div = utils.AddClass(id, h.Div)
	H1  = utils.AddClass(id, h.H1)
	P   = utils.AddClass(id, h.P)
)

func Layout(children ...g.Node) g.Node {
	return c.HTML5(c.HTML5Props{
		Title: "ma page",
		Head: []g.Node{
			h.Link(h.Href("/style"), h.Rel("stylesheet")),
		},
		Body: []g.Node{
			H1(g.Text("ma page de fou")),
		},
	})
}
