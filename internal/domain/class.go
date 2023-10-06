package domain

type Class struct {
	Id    uint16 `gorm:"primaryKey;column:id"`
	Title string `gorm:"type:varchar(255);column:title"`

	Teachers []*User    `gorm:"many2many:teacher2class;joinForeignKey:class_id;joinReferences:id"`
	Subject  []*Subject `gorm:"many2many:class2subject"`
}

func (Class) TableName() string {
	return "class"
}
