package domain

import "time"

type User struct {
	Id              uint64 `gorm:"primaryKey;column:id;not null"`
	Name            string `gorm:"type:varchar(255);not null"`
	Email           string `gorm:"type:varchar(255);uniqueIndex;not null"`
	EmailVerifiedAt *time.Time
	Password        string `gorm:"type:varchar(255);not null"`
	RememberToken   string `gorm:"type:varchar(100)"`
	CreatedAt       *time.Time
	UpdatedAt       *time.Time

	ClassId      uint16  `gorm:"column:class_id"`
	StudentClass *Class  `gorm:"foreignKey:class_id;references:id;"`
	StudentTasks []*Task `gorm:"many2many:student2task"`

	TeacherClasses []*Class `gorm:"many2many:teacher2class"`
}

func (User) TableName() string {
	return "users"
}

type StudentToTask struct {
	Id        uint64 `gorm:"primaryKey;column:id"`
	StudentId uint64 `gorm:"column:student_id"`
	TaskId    uint64 `gorm:"column:task_id"`
	IsDone    bool   `gorm:"type:tinyint(1);column:is_done"`
}

func (StudentToTask) TableName() string {
	return "student2task"
}

type TeacherToClass struct {
	Id        uint64 `gorm:"primaryKey;column:id"`
	TeacherId uint64 `gorm:"column:teacher_id"`
	ClassId   uint16 `gorm:"column:class_id"`
}

func (TeacherToClass) TableName() string {
	return "teacher2class"
}
