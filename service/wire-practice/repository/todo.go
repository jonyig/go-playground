package repository

type TodoRepository struct {
}

func NewTodoRepository() *TodoRepository {
	return &TodoRepository{}
}

func (t *TodoRepository) Get() string {
	return "789"
}
