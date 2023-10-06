package domain

type Class struct {
	Id    uint16 `gorm:"primaryKey;column:id"`
	Title string `gorm:"type:varchar(255);column:title"`

	Teachers []*User `gorm:"many2many:teacher2class"`
}

func (Class) TableName() string {
	return "class"
}
