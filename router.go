package main

import (
	. "background/apis"
	"github.com/gin-gonic/gin"
)

func initRouter() *gin.Engine {
	router := gin.Default()

	router.GET("/", IndexApi)

	router.GET("/login", LoginApi)

	router.GET("/classes_student", GetClassesStudentApi)

	router.GET("/classes_teacher", GetClassesTeacherApi)

	router.GET("/joinclass", JoinClassApi)

	router.GET("/homeworks_student", GetHomeworkListStudentApi)

	router.POST("/uploadhomework", UploadHomeworkApi)

	router.GET("createhomework", CreateHomeworkInfoApi)

	router.POST("createclass", CreateClassApi)

	router.GET("downloadhomework", DownloadHomeworkApi)

	router.GET("downloadallhomework", DownloadAllHomeworkApi)

	router.GET("homeworks_teacher", GetHomeworkListTeacherApi)

	router.GET("getstudentclass", GetStudentListApi)

	router.GET("firestudent", DeleteStudentApi)

	router.GET("gethomeworkany", GetHomeworkListApi)

	router.GET("mark", AddScoreApi)

	return router
}
