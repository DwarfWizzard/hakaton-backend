package domain

type Lecture struct {
	Id          uint64       `gorm:"primaryKey;column:id"`
	TopicId     uint64       `gorm:"column:topic_id"`
	TypeId      uint16       `gorm:"column:type_id"`
	Type        *LectureType `gorm:"foreignKey:type_id;references:id"`
	Title       string       `gorm:"type:varchar(60);column:title"`
	Description string       `gorm:"type:text;colum:title"`

	Media []*Media `gorm:"many2many:lecture2media"`
}

func (Lecture) TableName() string {
	return "lecture"
}

type LectureType struct {
	Id    uint16 `gorm:"primaryKey;column:id"`
	Title string `gorm:"type:varchar(10);column:title"`
}

func (LectureType) TableName() string {
	return "lecture_type"
}
