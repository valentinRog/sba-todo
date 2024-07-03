package todos

import (
	g "github.com/maragudk/gomponents"
	h "github.com/maragudk/gomponents/html"
	"github.com/valentinRog/sba-todo/store/todo"
	"github.com/valentinRog/sba-todo/ui/utils"
)

var (
	Ul     = utils.AddClass(id, h.Ul)
	Li     = utils.AddClass(id, h.Li)
	Form   = utils.AddClass(id, h.Form)
	Input  = utils.AddClass(id, h.Input)
	Button = utils.AddClass(id, h.Button)
)

func addTodoForm() g.Node {
	return Form(
		g.Attr("hx-post", "/addtodo"),
		g.Attr("hx-target", "#todo-list"),
		g.Attr("hx-swap", "outerHTML"),
		Input(h.Type("text"), h.Name("text")),
		Button(h.Type("submit"), g.Text("add todo")),
	)
}

func TodoList() g.Node {
	todos := todo.GetTodos()
	return Ul(h.ID("todo-list"),
		g.Group(g.Map(todos, func(todo todo.Todo) g.Node {
			return Li(g.Text(todo.Name))
		})))
}

func Todos() g.Node {
	return g.Group([]g.Node{TodoList(), addTodoForm()})
}
