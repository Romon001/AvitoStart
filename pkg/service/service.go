package service

type Repository struct {
	Segments
	UserSegmentPair
}
type Segments interface {
	Create(name string) (int, error)
	GetAll() ([]avitoStartApp.Segment, error)
	GetById(userId, listId int) (avitoStartApp.TodoList, error)
	Delete(userId, listId int) error
	Update(userId, listId int, input avitoStartApp.UpdateListInput) error
}

type UserSegmentPair interface {
	Create(listId int, item todo.TodoItem) (int, error)
	GetAll(userId, listId int) ([]todo.TodoItem, error)
	GetById(userId, itemId int) (todo.TodoItem, error)
	Delete(userId, itemId int) error
	Update(userId, itemId int, input todo.UpdateItemInput) error
}
