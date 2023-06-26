package entities

import "github.com/msegeya56/artery.go.module/pkg/tools/commons"

type Form struct {
	commons.FoundationEntity
}

type Art struct {
	commons.FoundationEntity
	FormID   uint
	FormList []Form
}

type Task struct {
	commons.FoundationEntity
	taskID   uint
	TodoList []Todo
}
type Todo struct {
	commons.FoundationEntity
}

// INCLUDE ToJSON and FromJSON for all models and entienties
