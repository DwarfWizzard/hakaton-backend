package domain

import "gorm.io/datatypes"

type User struct {
	Id              uint64 `gorm:"primaryKey;column:id;not null"`
	Name            string `gorm:"type:varchar(255);not null"`
	FirstName       string `gorm:"column:first_name"`
	SecondName      string `gorm:"column:second_name"`
	Email           string `gorm:"type:varchar(255);uniqueIndex;not null"`
	EmailVerifiedAt datatypes.Time
	Password        string `gorm:"type:varchar(255);not null"`
	RememberToken   string `gorm:"type:varchar(100)"`
	CreatedAt       datatypes.Time
	UpdatedAt       datatypes.Time

	ClassId      uint16 `gorm:"column:class_id"`
	StudentClass *Class `gorm:"foreignKey:class_id;references:id"`

	TeacherClasses []*Class `gorm:"many2many:teacher2class;joinForeignKey:teacher_id;joinReferences:id"`
}

func (User) TableName() string {
	return "users"
}

type StudentToTask struct {
	Id        uint64 `gorm:"primaryKey;column:id"`
	StudentId uint64 `gorm:"column:student_id"`
	TaskId    uint64 `gorm:"column:task_id"`
	IsDone    *bool  `gorm:"type:tinyint(1);column:is_done"`
}

func (StudentToTask) TableName() string {
	return "student2task"
}

type StudentToLecture struct {
	Id        uint64 `gorm:"primaryKey;column:id"`
	StudentId uint64 `gorm:"column:student_id"`
	LectureId uint64 `gorm:"column:lecture_id"`
}

func (StudentToLecture) TableName() string {
	return "student2lecture"
}
