package models

import (
	"github.com/msegeya56/artery.go.module/pkg/domain/entities"
	"github.com/msegeya56/artery.go.module/pkg/tools/commons"
)

type Form struct {
	commons.Foundation
}

type Art struct {
	commons.Foundation
	FormID uint `json:"form_id,omitempty"`
}

type Task struct {
	commons.Foundation
	TodoID uint `json:"todo_id,omitempty"`
}

type Todo struct {
	commons.Foundation
}

// INCLUDE ToJSON and FromJSON for all models and entienties

type ArtReply struct {
	Data        *entities.Art
	Collection  []entities.Art
	Stream      <-chan entities.Art
	Error       error
	ErrorStream <-chan error
}
type TodoReply struct {
	Data        *entities.Todo
	Collection  []entities.Todo
	Stream      <-chan entities.Todo
	Error       error
	ErrorStream <-chan error
}
type TaskReply struct {
	Data        *entities.Task
	Collection  []entities.Task
	Stream      <-chan entities.Task
	Error       error
	ErrorStream <-chan error
}

type FormReply struct {
	Data        *entities.Form
	Collection  []entities.Form
	Stream      <-chan entities.Form
	Error       error
	ErrorStream <-chan error
}
