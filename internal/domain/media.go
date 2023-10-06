package domain

type Media struct {
	Id       uint64            `gorm:"primaryKey;column:id"`
	ParentId uint64            `gorm:"type:bigint;column:parent_id"`
	TypeId   uint16             `gorm:"column:type_id"`
	Type     *MediaContentType `gorm:"foreignKey:type_id;references:id"`
	Content  string            `gorm:"type:text"`
}

func (Media) TableName() string {
	return "media"
}

type MediaContentType struct {
	Id    uint16 `gorm:"primaryKey;column:id"`
	Title string `gorm:"type:varchar(25);column:title"`
}

func (MediaContentType) TableName() string {
	return "media_content_type"
}
