package homelayout

import (
	g "github.com/maragudk/gomponents"
	h "github.com/maragudk/gomponents/html"
	"github.com/valentinRog/sba-todo/ui/utils"
)

var (
	p   = utils.AddClass(Id, h.H1)
	div = utils.AddClass(Id, h.Div)
)

func Layout(username string, content g.Node) g.Node {
	return div(
		p(g.Text(username)),
		content,
	)
}
