package todos

import (
	g "github.com/maragudk/gomponents"
	h "github.com/maragudk/gomponents/html"
	"github.com/valentinRog/sba-todo/store/todo"
	"github.com/valentinRog/sba-todo/ui/utils"
)

var (
	Ul = utils.AddClass(id, h.Ul)
	Li = utils.AddClass(id, h.Li)
)

func Todos() g.Node {
	todos := todo.GetTodos()
	return Ul(g.Group(g.Map(todos, func(todo todo.Todo) g.Node {
		return Li(g.Text(todo.Name))
	})))
}
