package service

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type MobileUserInfoResponse struct {
	Id        uint64 `json:"id"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	ClassName string `json:"class"`
}

// MobileUserInfo - GET /api/user
func (s *Service) MobileUserInfo(c echo.Context, apikey *ApiKey) error {
	user := apikey.User

	resp := &MobileUserInfoResponse{
		Id:        user.Id,
		FirstName: user.FirstName,
		LastName:  user.SecondName,
	}

	if user.StudentClass != nil {
		resp.ClassName = user.StudentClass.Title
	}

	return c.JSON(http.StatusOK, &Response{
		Result: resp,
	})
}

type MobileSubjectPoints struct {
	CurrentPoints int `json:"currentPoints"`
	MaxPoints     int `json:"maxPoints"`
}

type MobileListSubjectByUserResponse struct {
	SubjectId    uint8                `json:"id"`
	SubjectTitle string               `json:"title"`
	TopicTitle   *string              `json:"topic"`
	Color        string               `json:"color"`
	Points       *MobileSubjectPoints `json:"points"`
}

// MobileListSubjectByUser - GET /api/user/subjects
func (s *Service) MobileListSubjectByUser(c echo.Context, apikey *ApiKey) error {
	user := apikey.User

	if user.StudentClass == nil {
		return errInvalidRequest
	}

	listSubject, err := s.repo.ListSubjectByClassId(user.ClassId)
	if err != nil {
		s.logger.Error("Get list subject error", err)
		return errInternalServerError
	}

	s.logger.Info("1")
	var resp []*MobileListSubjectByUserResponse
	for _, subject := range listSubject {

		subjectResp := &MobileListSubjectByUserResponse{
			SubjectId:    subject.Id,
			SubjectTitle: subject.Title,
			Color:        subject.Color,
		}

		listTopics, err := s.repo.ListTopicBySubjectAndClassIds(subject.Id, user.ClassId)
		if err != nil {
			s.logger.Error("Get list topics error", err)
		}

		if len(listTopics) == 0 {
			resp = append(resp, subjectResp)
			continue
		}

		currentTopic := listTopics[0]

		subjectResp.TopicTitle = &currentTopic.Title

		completedLectures, err := s.repo.ListCompletedLectureByUserAndTopicIds(user.Id, currentTopic.Id)
		if err != nil {
			s.logger.Error("Get list lecture error", err)
			return errInternalServerError
		}

		completedTasks, err := s.repo.ListCompletedTaskByUserAndTopicIds(user.Id, currentTopic.Id)
		if err != nil {
			s.logger.Error("Get list task error", err)
			return errInternalServerError
		}

		subjectResp.Points = &MobileSubjectPoints{
			CurrentPoints: len(completedLectures) + len(completedTasks),
			MaxPoints:     len(currentTopic.Lectures) + len(currentTopic.Tasks),
		}

		resp = append(resp, subjectResp)
	}

	return c.JSON(http.StatusOK, &Response{Result: resp})
}

type MobileListHomeworkByUserResponse struct {
	SubjectId    uint8  `json:"id"`
	SubjectTitle string `json:"title"`
	HwDesc       string `json:"description"`
	Date         string `json:"date"`
	Color        string `json:"homework"`
}

// MobileListHomeworkByUser - GET /api/user/homework
func (s *Service) MobileListHomeworkByUser(c echo.Context, apikey *ApiKey) error {
	user := apikey.User

	if user.StudentClass == nil {
		return errInvalidRequest
	}

	listHw, err := s.repo.ListHomeworkByClassIds(user.ClassId)
	if err != nil {
		s.logger.Error("Get homework list error", err)
		return errInternalServerError
	}

	var resp []*MobileListHomeworkByUserResponse
	for _, homework := range listHw {
		if homework.Subject == nil {
			s.logger.Warn("Empty subject field")
			continue
		}
		resp = append(resp, &MobileListHomeworkByUserResponse{
			SubjectId:    homework.SubjectId,
			SubjectTitle: homework.Subject.Title,
			Color:        homework.Subject.Color,
			HwDesc:       homework.Description,
			Date:         homework.DoneAt,
		})
	}

	return c.JSON(http.StatusOK, &Response{Result: resp})
}
