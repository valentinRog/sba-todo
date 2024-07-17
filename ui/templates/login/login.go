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
)

func signupForm() g.Node {
	return Form(
		g.Attr("hx-post", "/login"),
		g.Attr("hx-redirect", "/"),
		Input(h.Type("text"), h.Name("text")),
		Button(h.Type("submit"), g.Text("signup")),
	)
}

func signinForm() g.Node {
	return Form()
}

func (Login) LoginPage() g.Node {
	return Div(
		H1(g.Text("Page de login")),
		signupForm(),
	)
}
