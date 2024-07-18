package login

import (
	g "github.com/maragudk/gomponents"
	h "github.com/maragudk/gomponents/html"
	"github.com/valentinRog/sba-todo/ui/utils"

	_ "embed"
)

type Login struct{}

var (
	H1     = utils.AddClass(id, h.H1)
	Div    = utils.AddClass(id, h.Div)
	Form   = utils.AddClass(id, h.Form)
	Input  = utils.AddClass(id, h.Input)
	Button = utils.AddClass(id, h.Button)
	a      = utils.AddClass(id, h.A)
)

func (Login) SigninForm() g.Node {
	return Div(
		h.ID("signin-form"),
		Form(
			g.Attr("hx-post", "/signin"),
			g.Attr("hx-target", "#content-div"),
			g.Attr("hx-swap", "outerHTML"),
			Input(h.Type("text"), h.Name("username")),
			Input(h.Type("text"), h.Name("password")),
			Button(h.Type("submit"), g.Text("signin")),
		),
		a(
			g.Text("signup"),
			g.Attr("hx-get", "/login/signup-form"),
			g.Attr("hx-target", "#signin-form"),
			g.Attr("hx-swap", "outerHTML"),
		),
	)
}
func (Login) SignupForm() g.Node {
	return Div(
		h.ID("signup-form"),
		Form(
			g.Attr("hx-post", "/signup"),
			g.Attr("hx-target", "#content-div"),
			g.Attr("hx-swap", "innerHTML"),
			Input(h.Type("text"), h.Name("username")),
			Input(h.Type("text"), h.Name("password")),
			Button(h.Type("submit"), g.Text("signup")),
		),
		a(
			g.Text("signin"),
			g.Attr("hx-get", "/login/signin-form"),
			g.Attr("hx-target", "#signup-form"),
			g.Attr("hx-swap", "outerHTML"),
		),
	)
}

func (l Login) LoginPage() g.Node {
	return Div(
		H1(g.Text("Page de login")),
		l.SigninForm(),
	)
}
