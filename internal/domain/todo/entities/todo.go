package entities

type Todo struct {
	ID          int
	Title       string
	Description string
	Completed   bool
}

func NewTodo(title, description string) *Todo {
	return &Todo{
		Title:       title,
		Description: description,
		Completed:   false,
	}
}

func (t *Todo) Complete() {
	t.Completed = true
}
