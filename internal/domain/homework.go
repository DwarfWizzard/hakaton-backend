package domain

import "time"

type Homework struct {
	Id          uint64     `gorm:"primaryKey;column:id;not null"`
	TeacherId   uint64     `gorm:"column:teacher_id"`
	Teacher     *User      `gorm:"foreignKey:teacher_id;references:id"`
	ClassId     uint16     `gorm:"column:class_id"`
	Class       *Class     `gorm:"column:class_id"`
	Description string     `gorm:"column:description"`
	DoneAt      *time.Time `gorm:"column:done_at"`
}

func (Homework) TableName() string {
	return "homework"
}
