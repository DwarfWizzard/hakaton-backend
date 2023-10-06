package mysql

import (
	"github.com/DwarfWizzard/hakaton-backend/internal/domain"
	"gorm.io/gorm"
)

type TaskRepo struct {
	dbClient *gorm.DB
}

func NewTaskRepo(dbClient *gorm.DB) *TaskRepo {
	return &TaskRepo{
		dbClient: dbClient,
	}
}

func (r *TaskRepo) GetTaskById(taskId uint64) (*domain.Task, error) {
	var task *domain.Task
	result := r.dbClient.Table("task").
		Where("task.id = ?", taskId).
		Preload("Type").
		Preload("Media").
		First(&task)
	if result.Error != nil {
		return nil, result.Error
	}

	return task, nil
}

func (r *TaskRepo) AddTaskToUser(taskId uint64, userId uint64) error {
	result := r.dbClient.Create(&domain.StudentToTask{
		StudentId: userId,
		TaskId:    taskId,
	})

	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (r *TaskRepo) SaveTaskAnswer(taskId uint64, userId uint64, answer bool) error {
	var s2tId uint64
	result := r.dbClient.Table("student2task").
		Select("student2task.id").
		Where("student2task.student_id = ? AND student2task.task_id", userId, taskId).
		Scan(&s2tId)

	if result.Error != nil {
		return result.Error
	}

	result = r.dbClient.Save(&domain.StudentToTask{
		Id:     s2tId,
		IsDone: &answer,
	})

	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (r *TaskRepo) ListCompletedTaskByUserAndTopicIds(userId, topicId uint64) ([]*domain.Task, error) {
	var listTask []*domain.Task
	result := r.dbClient.Table("task").Select(
		`task.id, 
			task.topic_id, 
			task.type_id, 
			task.title, 
			task.description,
			task.done_at,
			task.task_params_json`,
	).
		Joins("inner join student2task on task.id=student2task.task_id").
		Where("student2task.is_done = 1 AND student2task.student_id = ? AND task.topic_id = ?", userId, topicId).
		Preload("Type").
		Find(&listTask)
	if result.Error != nil {
		return nil, result.Error
	}

	return listTask, nil
}
