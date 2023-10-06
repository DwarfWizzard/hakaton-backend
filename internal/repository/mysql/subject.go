package mysql

import (
	"github.com/DwarfWizzard/hakaton-backend/internal/domain"
	"gorm.io/gorm"
)

type SubjectRepo struct {
	dbClient *gorm.DB
}

func NewSubjectRepo(dbClient *gorm.DB) *SubjectRepo {
	return &SubjectRepo{
		dbClient: dbClient,
	}
}

func (r *SubjectRepo) ListSubjectByClassId(classId uint16) ([]*domain.Subject, error) {
	var subject []*domain.Subject
	result := r.dbClient.Table("subject").Select(
		`subject.id, 
		subject.title, 
		subject.color`,
	).
		Joins("inner join class2subject on class2subject.subject_id=subject.id").
		Where("class2subject.class_id=?", classId).
		Find(&subject)
	if result.Error != nil {
		return nil, result.Error
	}

	return subject, nil
}
