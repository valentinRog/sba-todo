package layout

import (
	g "github.com/maragudk/gomponents"
	c "github.com/maragudk/gomponents/components"
	h "github.com/maragudk/gomponents/html"
	"github.com/valentinRog/sba-todo/ui/utils"

	_ "embed"
)

var (
	H1  = utils.AddClass(id, h.H1)
	Ul  = utils.AddClass(id, h.Ul)
	Li  = utils.AddClass(id, h.Li)
	A   = utils.AddClass(id, h.A)
	Div = utils.AddClass(id, h.Div)
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
			Ul(
				Li(A(
					h.Href("/"),
					g.Text("home"),
					g.Attr("hx-get", "/"),
					g.Attr("hx-target", "#content-div"),
					g.Attr("hx-swap", "innerHTML"),
					g.Attr("hx-push-url", "true"),
				)),
				Li(A(
					h.Href("/login"),
					g.Text("login"),
					g.Attr("hx-get", "/login"),
					g.Attr("hx-target", "#content-div"),
					g.Attr("hx-swap", "innerHTML"),
					g.Attr("hx-push-url", "true"),
				)),
			),
			Div(h.ID("content-div"), g.Group(children)),
		},
	})
}
