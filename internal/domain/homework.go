package domain

type Homework struct {
	Id          uint64   `gorm:"primaryKey;column:id;not null"`
	TeacherId   uint64   `gorm:"column:teacher_id"`
	Teacher     *User    `gorm:"foreignKey:teacher_id;references:id"`
	ClassId     uint16   `gorm:"column:class_id"`
	Class       *Class   `gorm:"foreignKey:class_id;references:id"`
	SubjectId   uint8    `gorm:"column:subject_id"`
	Subject     *Subject `gorm:"foreignKey:subject_id;references:id"`
	Description string   `gorm:"column:description"`
	DoneAt      string   `gorm:"column:done_at"`
}

func (Homework) TableName() string {
	return "homework"
}
