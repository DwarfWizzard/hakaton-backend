package mysql

import (
	"github.com/DwarfWizzard/hakaton-backend/internal/domain"
	"gorm.io/gorm"
)

type HomeworkRepo struct {
	dbClient *gorm.DB
}

func NewHomeworkRepo(dbClient *gorm.DB) *HomeworkRepo {
	return &HomeworkRepo{
		dbClient: dbClient,
	}
}

func (r *HomeworkRepo) ListHomeworkByClassIds(classId uint16) ([]*domain.Homework, error) {
	var listHomeWork []*domain.Homework
	result := r.dbClient.
		Table("homework").
		Where("homework.class_id = ?", classId).
		Preload("Subject").
		Preload("Teacher").
		Preload("Class").
		Find(&listHomeWork)
	if result.Error != nil {
		return nil, result.Error
	}

	return listHomeWork, nil
}