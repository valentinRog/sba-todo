package login

import (
	g "github.com/maragudk/gomponents"
	h "github.com/maragudk/gomponents/html"
	"github.com/valentinRog/sba-todo/ui/utils"

	_ "embed"
)

var (
	H1 = utils.AddClass(id, h.H1)
)

func LoginPage() g.Node {
	return H1(g.Text("Page de login"))
}
