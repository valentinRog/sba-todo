package templates

import (
	"github.com/valentinRog/sba-todo/ui/templates/layout"
	"github.com/valentinRog/sba-todo/ui/templates/login"
	"github.com/valentinRog/sba-todo/ui/templates/todos"
)

type Templates struct {
	Layout layout.Layout
	Login  login.Login
	Todos  todos.Todos
}
