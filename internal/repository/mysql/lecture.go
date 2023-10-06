package mysql

import (
	"github.com/DwarfWizzard/hakaton-backend/internal/domain"
	"gorm.io/gorm"
)

type LectureRepo struct {
	dbClient *gorm.DB
}

func NewLectureRepo(dbClient *gorm.DB) *LectureRepo {
	return &LectureRepo{
		dbClient: dbClient,
	}
}

func (r *LectureRepo) GetLectureById(lectureId uint64) (*domain.Lecture, error) {
	var lecture *domain.Lecture
	result := r.dbClient.Table("lecture").
		Where("lecture.id = ?", lectureId).
		Preload("Type").
		Preload("Media").
		First(&lecture)
	if result.Error != nil {
		return nil, result.Error
	}

	return lecture, nil
}

func (r *TaskRepo) AddLectureToUser(lectureId uint64, userId uint64) error {
	result := r.dbClient.Create(&domain.StudentToLecture{
		StudentId: userId,
		LectureId: lectureId,
	})

	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (r *LectureRepo) ListCompletedLectureByUserAndTopicIds(userId, topicId uint64) ([]*domain.Lecture, error) {
	var lecture []*domain.Lecture
	result := r.dbClient.Table("lecture").Select(
		`lecture.id, 
		lecture.topic_id, 
		lecture.type_id, 
		lecture.title, 
		lecture.description`,
		).
		Joins("inner join student2lecture on lecture.id=student2lecture.lecture_id").
		Where("student2lecture.student_id = ? AND lecture.topic_id = ?", userId, topicId).
		Preload("Type").
		Find(&lecture)
	if result.Error != nil {
		return nil, result.Error
	}

	return lecture, nil
}