package login

import (
	g "github.com/maragudk/gomponents"
	h "github.com/maragudk/gomponents/html"
	"github.com/valentinRog/sba-todo/ui/utils"

	_ "embed"
)

var (
	h1     = utils.AddClass(Id, h.H1)
	div    = utils.AddClass(Id, h.Div)
	form   = utils.AddClass(Id, h.Form)
	input  = utils.AddClass(Id, h.Input)
	button = utils.AddClass(Id, h.Button)
	a      = utils.AddClass(Id, h.A)
)

func LoginPage(form g.Node) g.Node {
	return div(
		h1(g.Text("Page de login")),
		form,
	)
}
