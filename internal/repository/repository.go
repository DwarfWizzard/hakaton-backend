package repository

import (
	"github.com/DwarfWizzard/hakaton-backend/internal/domain"
	"github.com/DwarfWizzard/hakaton-backend/internal/repository/mysql"
	"gorm.io/gorm"
)

type User interface {
	GetUserByLogin(login string) (*domain.User, error)
	GetUserById(userId uint64) (*domain.User, error)
}

type Subject interface {
	ListSubjectByClassId(classId uint16) ([]*domain.Subject, error)
}

type Topic interface {
	ListTopicBySubjectAndClassIds(subjectId uint8, classId uint16) ([]*domain.Topic, error)
}

type Lecture interface {
	ListCompletedLectureByUserAndTopicIds(userId, topicId uint64) ([]*domain.Lecture, error)
}

type Task interface {
	GetTaskById(taskId uint64) (*domain.Task, error)
	ListCompletedTaskByUserAndTopicIds(userId, topicId uint64) ([]*domain.Task, error)
}

type Homework interface {
	ListHomeworkByClassIds(classId uint16) ([]*domain.Homework, error)
}

type Repository struct {
	User
	Subject
	Topic
	Lecture
	Task
	Homework
}

func New(dbClient *gorm.DB) *Repository {
	return &Repository{
		User:    mysql.NewUserRepo(dbClient),
		Subject: mysql.NewSubjectRepo(dbClient),
		Topic:   mysql.NewTopicRepo(dbClient),
		Lecture: mysql.NewLectureRepo(dbClient),
		Task:    mysql.NewTaskRepo(dbClient),
		Homework: mysql.NewHomeworkRepo(dbClient),
	}
}
