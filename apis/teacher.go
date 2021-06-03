package apis

import (
	. "background/models"
	"github.com/gin-gonic/gin"
)

// GetClassesTeacherApi
// 	path: classes_teacher
//  Input: ID,password
//  Output(JSON array)
//  	img: String
// （标识课程的照片，如果没有照片就弄个默认的）
//  	name: String
// （课程名字）
//  	ID: String
// （课程ID）
func GetClassesTeacherApi(c *gin.Context) {
	id := c.Query("ID")
	password := c.Query("password")

	t := Teacher{Id: id, Password: password}
	classes := t.GetClassesTeacher()

	c.IndentedJSON(200, classes)
}

