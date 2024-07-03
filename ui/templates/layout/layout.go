package layout

import (
	g "github.com/maragudk/gomponents"
	c "github.com/maragudk/gomponents/components"
	h "github.com/maragudk/gomponents/html"
	"github.com/valentinRog/sba-todo/ui/utils"

	_ "embed"
)

var (
	H1 = utils.AddClass(id, h.H1)
)

func Layout(children ...g.Node) g.Node {
	return c.HTML5(c.HTML5Props{
		Title: "ma page",
		Head: []g.Node{
			h.Link(h.Href("/style"), h.Rel("stylesheet")),
			h.Script(h.Src("/htmx")),
		},
		Body: []g.Node{
			H1(g.Text("ma page de fou")),
			g.Group(children),
		},
	})
}
