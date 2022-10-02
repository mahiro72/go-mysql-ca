package entity

// TaskはTodoリストのタスクに関する情報を保持する構造体です
type Task struct {
	Id   int
	Name string
	Status string
}

type Tasks []*Task
