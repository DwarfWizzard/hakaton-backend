package mysql

import (
	"github.com/DwarfWizzard/hakaton-backend/internal/domain"
	"gorm.io/gorm"
)

type TopicRepo struct {
	dbClient *gorm.DB
}

func NewTopicRepo(dbClient *gorm.DB) *TopicRepo {
	return &TopicRepo{
		dbClient: dbClient,
	}
}

func (r *TopicRepo) ListTopicBySubjectAndClassIds(subjectId uint8, classId uint16) ([]*domain.Topic, error) {
	var listTopic []*domain.Topic
	result := r.dbClient.
		Table("topic").
		Where("topic.subject_id = ? AND topic.class_id = ?", subjectId, classId).
		Order("topic.created_at DESC").
		Preload("Lectures").
		Preload("Tasks").
		Preload("Lectures.Type").
		Preload("Tasks.Type").
		Preload("Lectures.Media").
		Preload("Tasks.Media").
		Preload("Lectures.Media.Type").
		Preload("Tasks.Media.Type").
		Find(&listTopic)
	if result.Error != nil {
		return nil, result.Error
	}
	return listTopic, nil
}
