package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	"github.com/DwarfWizzard/hakaton-backend/internal/repository"
	"github.com/DwarfWizzard/hakaton-backend/internal/service"
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	dsn  = "root:1234@tcp(localhost:3306)/hakaton_db"
	port = "8080"
)

func main() {
	logger := NewLogger(logrus.DebugLevel)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		logger.Panic(err)
	}

	repo := repository.New(db)

	svc := service.New(logger, repo)

	router := echo.New()
	router.HTTPErrorHandler = svc.HttpErrorHandler
	router.HideBanner = true
	router.HidePort = true

	rtAuth := router.Group("/auth")
	rtAuth.POST("/login", svc.BindApi((*service.Service).MobileLogin))
	rtAuth.GET("/refresh", svc.BindApi((*service.Service).MobileRefreshToken))

	rtUser := router.Group("/api/user")
	rtUser.GET("", svc.BindApiWithAuth((*service.Service).MobileUserInfo))
	rtUser.GET("/subjects", svc.BindApiWithAuth((*service.Service).MobileListSubjectByUser))
	rtUser.GET("/homework", svc.BindApiWithAuth((*service.Service).MobileListHomeworkByUser))

	rtSubject := router.Group("/api/subject")
	rtSubject.GET("/:id/topics", svc.BindApiWithAuth((*service.Service).MobileListTopicBySubject))

	// WIP //
	
	// rtTask := router.Group("/api/task")
	// rtTask.GET("/:id", svc.BindApiWithAuth((*service.Service).MobileTaskById))
	// rtTask.POST("/:id", svc.BindApiWithAuth((*service.Service).MobileAddTaskToUser))
	// rtTask.PUT("/:id", svc.BindApiWithAuth((*service.Service).MobileTaskAnswer))

	// rtLecture := router.Group("/api/lecture")
	// rtLecture.GET("/:id", svc.BindApiWithAuth((*service.Service).MobileLectureById))
	// rtLecture.POST("/:id", svc.BindApiWithAuth((*service.Service).MobileAddLectureToUser))

	go func() {
		if err := router.Start(":" + port); err != nil {
			logger.Warn("API web server stopped with error.", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT, syscall.SIGQUIT)
	<-quit

	logger.Info("Shutdown server")
	if err := router.Shutdown(context.Background()); err != nil {
		logger.Error(err)
	}
}
