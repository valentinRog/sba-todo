package login

import (
	g "github.com/maragudk/gomponents"
	h "github.com/maragudk/gomponents/html"
	"github.com/valentinRog/sba-todo/ui/utils"

	_ "embed"
)

var (
	h1     = utils.AddClass(id, h.H1)
	div    = utils.AddClass(id, h.Div)
	form   = utils.AddClass(id, h.Form)
	input  = utils.AddClass(id, h.Input)
	button = utils.AddClass(id, h.Button)
	a      = utils.AddClass(id, h.A)
)

func SigninForm() g.Node {
	return div(
		h.ID("signin-form"),
		form(
			g.Attr("hx-post", "/auth/signin"),
			g.Attr("hx-target", "#content-div"),
			g.Attr("hx-swap", "outerHTML"),
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
func SignupForm() g.Node {
	return div(
		h.ID("signup-form"),
		form(
			g.Attr("hx-post", "/auth/signup"),
			g.Attr("hx-target", "#content-div"),
			g.Attr("hx-swap", "innerHTML"),
			input(h.Type("text"), h.Name("username")),
			input(h.Type("text"), h.Name("password")),
			button(h.Type("submit"), g.Text("signup")),
		),
		a(
			g.Text("signin"),
			g.Attr("hx-get", "/login/signin-form"),
			g.Attr("hx-target", "#signup-form"),
			g.Attr("hx-swap", "outerHTML"),
		),
	)
}

func LoginPage() g.Node {
	return div(
		h1(g.Text("Page de login")),
		SigninForm(),
	)
}
