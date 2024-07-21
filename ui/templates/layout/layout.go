package layout

import (
	g "github.com/maragudk/gomponents"
	c "github.com/maragudk/gomponents/components"
	h "github.com/maragudk/gomponents/html"
	"github.com/valentinRog/sba-todo/ui/utils"

	_ "embed"
)

var (
	H1  = utils.AddClass(Id, h.H1)
	Ul  = utils.AddClass(Id, h.Ul)
	Li  = utils.AddClass(Id, h.Li)
	A   = utils.AddClass(Id, h.A)
	Div = utils.AddClass(Id, h.Div)
)

func Layout(children ...g.Node) g.Node {
	return c.HTML5(c.HTML5Props{
		Title: "ma page",
		Head: []g.Node{
			h.Link(h.Href("/static/style"), h.Rel("stylesheet")),
			h.Script(h.Src("/static/htmx")),
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
					utils.HxHeader(Id),
				)),
				Li(A(
					h.Href("/login-content"),
					g.Text("login"),
					g.Attr("hx-get", "/login-content"),
					g.Attr("hx-target", "#content-div"),
					g.Attr("hx-swap", "innerHTML"),
					g.Attr("hx-push-url", "/login"),
					utils.HxHeader(Id),
				)),
			),
			Div(h.ID("content-div"), g.Group(children)),
		},
	})
}
