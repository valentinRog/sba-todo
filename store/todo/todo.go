package todo

type Todo struct {
	Name string
}

var todos = []Todo{}

func AddTodo(name string) {
	todos = append(todos, Todo{Name: name})
}

func GetTodos() []Todo {
	return todos
}
