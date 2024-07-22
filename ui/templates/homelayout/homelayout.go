package homelayout

import (
	g "github.com/maragudk/gomponents"
	h "github.com/maragudk/gomponents/html"
	"github.com/valentinRog/sba-todo/ui/utils"
)

var (
	p     = utils.AddClass(Id, h.H1)
	div   = utils.AddClass(Id, h.Div)
	form  = utils.AddClass(Id, h.Form)
	input = utils.AddClass(Id, h.Input)
)

func Layout(username string, content g.Node) g.Node {
	return div(
		p(g.Text(username)),
		form(
			h.Action("/auth/logout"),
			h.Method("POST"),
			input(h.Type("submit"), h.Value("logout")),
		),
		content,
	)
}
