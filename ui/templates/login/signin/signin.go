package signin

import (
	g "github.com/maragudk/gomponents"
	h "github.com/maragudk/gomponents/html"
	"github.com/valentinRog/sba-todo/ui/utils"

	_ "embed"
)

var (
	div    = utils.AddClass(Id, h.Div)
	form   = utils.AddClass(Id, h.Form)
	input  = utils.AddClass(Id, h.Input)
	button = utils.AddClass(Id, h.Button)
	a      = utils.AddClass(Id, h.A)
)

func SigninForm() g.Node {
	return div(
		h.ID("signin-form"),
		form(
			g.Attr("action", "/auth/signin"),
			g.Attr("method", "POST"),
			input(h.Type("text"), h.Name("username")),
			input(h.Type("text"), h.Name("password")),
			button(h.Type("submit"), g.Text("signin")),
		),
		a(
			g.Text("signup"),
			g.Attr("hx-get", "/login/signup-form"),
			g.Attr("hx-target", "#signin-form"),
			g.Attr("hx-swap", "outerHTML"),
		),
	)
}