package todo

type Todo struct {
	Id   int
	Name string
}

var todos = []Todo{}
var id = 1

func AddTodo(name string) {
	todos = append(todos, Todo{Id: id, Name: name})
	id += 1
}

func GetTodos() []Todo {
	return todos
}

func DeleteTodo(id int) {
	newTodos := []Todo{}
	for _, todo := range todos {
		if todo.Id != id {
			newTodos = append(newTodos, todo)
		}
	}
	todos = newTodos
}
