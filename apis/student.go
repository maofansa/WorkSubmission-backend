package apis

import (
	. "background/models"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// GetClassesStudentApi
// 	path: classes_student
//  Input: ID,password
//  Output(JSON array)
//  	img: String
// （标识课程的照片，如果没有照片就弄个默认的）
//  	name: String
// （课程名字）
//  	ID: String
// （课程ID）
func GetClassesStudentApi(c *gin.Context) {
	id := c.Query("ID")
	password := c.Query("password")

	s := Student{Id: id, Password: password}
	classes := s.GetClassesStudent()

	c.IndentedJSON(200, classes)
}

// JoinClassApi
// 	path: joinclass
//  Input: ID, password, classid
//  Output(JSON)
//  	result: int
//（1表示加入成功，0表示加入不成功）
func JoinClassApi(c *gin.Context) {
	id := c.Query("ID")
	password := c.Query("password")
	classId := c.Query("classid")

	s := Student{Id: id, Password: password}
	joined := s.JoinClass(classId)

	c.JSON(http.StatusOK, gin.H{
		"result": joined,
	})
}

// GetHomeworkListStudentApi
//	path:homeworks_student
//  Input: ID, password, classid
//      type: int
// （返回array的排序准则，0为按发布时间排序，最先发布的排在前面，索引小，为1的话还未提交的作业排在前面，其余的按照发布时间排序）
//  Output(JSONArray)
//  	name: string
//      pub_time: string
//      ddl: string
//      status: string
//      id(id of homework)
func GetHomeworkListStudentApi(c *gin.Context) {
	id := c.Query("ID")
	password := c.Query("password")
	classId := c.Query("classid")
	rank, _ := strconv.Atoi(c.Query("type"))

	s := Student{Id: id, Password: password}
	hwList := s.GetHomeworkListStudent(classId, rank)

	c.IndentedJSON(200, hwList)
}
