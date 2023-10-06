package domain

type Subject struct {
	Id    uint8  `gorm:"primaryKey;column:id"`
	Title string `gorm:"type:varchar(25);column:title"`
	Color string `gorm:"type:varchar(7);column:color"`
}

func (Subject) TableName() string {
	return "subject"
}
