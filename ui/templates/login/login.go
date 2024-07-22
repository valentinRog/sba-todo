package login

import (
	g "github.com/maragudk/gomponents"
	h "github.com/maragudk/gomponents/html"
	"github.com/valentinRog/sba-todo/ui/utils"
)

var (
	h1     = utils.AddClass(Id, h.H1)
	div    = utils.AddClass(Id, h.Div)
)

func LoginPage(form g.Node) g.Node {
	return div(
		h1(g.Text("Page de login")),
		form,
	)
}
