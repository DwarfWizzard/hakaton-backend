package service

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type MobileMedia struct {
	TypeId  string `json:"typeId"`
	Content string `json:"content"`
}

type MobileLectureItem struct {
	TypeId      string         `json:"typeId"`
	Title       string         `json:"title"`
	Description string         `json:"description"`
	Media       []*MobileMedia `json:"media"`
}

type MobileTaskItem struct {
	TypeId      string         `json:"typeId"`
	Title       string         `json:"title"`
	Description string         `json:"description"`
	Params      map[string]any `json:"params"`
	Media       []*MobileMedia `json:"media"`
}

type MobileListTopicBySubjectResponse struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	ThemeNumber int    `json:"themeNumber"`
	Items       []any  `json:"items"`
}

// MobileListTopicBySubject - GET /api/subject/:id/topics
func (s *Service) MobileListTopicBySubject(c echo.Context, apikey *ApiKey) error {
	user := apikey.User

	if user.StudentClass == nil {
		return errInvalidRequest
	}

	subjectIdStr := c.Param("id")
	if len(subjectIdStr) == 0 {
		return errEmptyRequiredParam
	}

	subjectId, err := strconv.ParseUint(subjectIdStr, 10, 8)
	if err != nil {
		return errInvalidParamType
	}

	topicList, err := s.repo.ListTopicBySubjectAndClassIds(uint8(subjectId), user.ClassId)
	if err != nil {
		s.logger.Error("Get topic list error", err)
		return errInternalServerError
	}

	var resp []*MobileListTopicBySubjectResponse
	for i, topic := range topicList {
		topicResp := &MobileListTopicBySubjectResponse{
			Title:       topic.Title,
			Description: topic.Description,
			ThemeNumber: i + 1,
		}

		var topicItems []any
		for _, lecture := range topic.Lectures {
			item := &MobileLectureItem{
				TypeId:      lecture.Type.Title,
				Title:       lecture.Title,
				Description: lecture.Description,
			}

			for _, media := range lecture.Media {
				item.Media = append(item.Media, &MobileMedia{
					TypeId:  media.Type.Title,
					Content: media.Content,
				})
			}

			topicItems = append(topicItems, item)
		}

		for _, task := range topic.Tasks {
			item := &MobileTaskItem{
				TypeId:      task.Type.Title,
				Title:       task.Title,
				Description: task.Description,
				Params:      task.Params,
			}

			for _, media := range task.Media {
				item.Media = append(item.Media, &MobileMedia{
					TypeId:  media.Type.Title,
					Content: media.Content,
				})
			}

			topicItems = append(topicItems, item)
		}

		topicResp.Items = topicItems

		resp = append(resp, topicResp)
	}

	return c.JSON(http.StatusOK, &Response{Result: resp})
}
