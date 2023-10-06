package domain

type Topic struct {
	Id          uint64      `gorm:"primaryKey;column:id;not null"`
	SubjectId   uint16      `gorm:"column:subject_id"`
	Subject     *Subject    `gorm:"foreignKey:subject_id;references:id"`
	ClassId     uint16      `gorm:"column:class_id"`
	Class       *Class      `gorm:"foreignKey:class_id;references:id"`
	ThemeId     uint8       `gorm:"column:theme_id"`
	Theme       *TopicTheme `gorm:"foreignKey:theme_id;references:id"`
	Title       string      `gorm:"type:varchar(60);column:title"`
	Description string      `gorm:"type:text;colum:title"`
	CreatedAt   string

	Lectures []*Lecture `gorm:"foreignKey:topic_id;references:id"`
	Tasks    []*Task    `gorm:"foreignKey:topic_id;references:id"`
}

func (Topic) TableName() string {
	return "topic"
}

type TopicTheme struct {
	Id    uint16 `gorm:"primaryKey;column:id"`
	Title string `gorm:"type:varchar(25);column:title"`
}

func (TopicTheme) TableName() string {
	return "topic_theme"
}
