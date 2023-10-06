package domain

import (
	"gorm.io/datatypes"
)

type Task struct {
	Id          uint64            `gorm:"primaryKey;column:id"`
	TopicId     uint64            `gorm:"column:topic_id"`
	TypeId      uint16            `gorm:"column:type_id"`
	Type        *TaskType         `gorm:"foreignKey:type_id;references:id"`
	Title       string            `gorm:"type:varchar(60);column:title"`
	Description string            `gorm:"type:text;colum:title"`
	DoneAt      string            `gorm:"column:done_at"`
	Params      datatypes.JSONMap `gorm:"column:task_params_json"`

	Media []*Media `gorm:"many2many:task2media"`
}

func (Task) TableName() string {
	return "task"
}

type TaskType struct {
	Id    uint16 `gorm:"primaryKey;column:id"`
	Title string `gorm:"type:varchar(10);column:title"`
}

func (TaskType) TableName() string {
	return "task_type"
}