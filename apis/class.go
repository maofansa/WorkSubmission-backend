package apis

import (
	. "background/models"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"path/filepath"
	"strconv"
)

// CreateClassApi
//path：createclass
//    + 参数ID，password同前，classname：string（课程名字），请求体中包含了图像数据。
//    + 返回状态包含在code中，未成功返回200之外的值。/*
func CreateClassApi(c *gin.Context) {
	teacherId := c.PostForm("ID")
	// password := c.PostForm("password")
	className := c.PostForm("classname")

	file, err := c.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": 400,
		})
		return
	}

	basePath := "/user/class/"
	fileName := basePath + filepath.Base(file.Filename)
	if err := c.SaveUploadedFile(file, fileName); err != nil {
		e := fmt.Sprintln(err)
		c.JSON(http.StatusRequestTimeout, gin.H{
			"code": 402,
			"err":e,
		})
		return
	}

	class := Class{TeacherId: teacherId, Name: className, PicURL: fileName}
	created := class.CreateClass()

	c.JSON(http.StatusOK, gin.H{
		"code": created,
	})
}

// GetHomeworkListTeacherApi
//+ path:homeworks_teacher
//    + 参数ID:string,password:string,classid:string(课程id号)，type：int
//      （返回array的排序准则，0为按发布时间排序，最先发布的排在前面，索引小，为1的话还未提交的作业排在前面，其余的按照发布时间排序）
//    + 返回JSONArray，其中每个元素为一个json，其中属性为name:string(作业名称)，pub_time:string（作业发布时间），
//   ddl：string（deadline），pub：string（提交了作业的人数)，unpub：string（未提交的人数），id(作业的id)。/*
func GetHomeworkListTeacherApi(c *gin.Context) {
	//teacherId := c.Query("ID")
	//password := c.Query("password")
	classId := c.Query("classid")
	rank, _ := strconv.Atoi(c.Query("type"))

	class := Class{Id: classId}
	hwList := class.GetHomeworkListTeacher(rank)

	c.IndentedJSON(200, hwList)
}

// GetStudentListApi
//+ path:getstudentclass
//    + 参数ID，password同前，classid：string（课程号id）
//    + 返回array，每一个项为json，属性name：string学生姓名，id：string（学生id）/*
func GetStudentListApi(c *gin.Context) {
	// teacherId := c.Query("ID")
	// password := c.Query("password")
	classId := c.Query("classid")

	class := Class{Id: classId}
	studentList := class.GetStudentList()

	c.IndentedJSON(200, studentList)
}

// DeleteStudentApi
//+ path:firestudent
//    + 参数ID，password同前，classid：string，studentid：string（学生id）
//    + 返回json,属性result：int（1成功0失败），老师将学生从该门课程中踢出。/*
func DeleteStudentApi(c *gin.Context) {
	// teacherId := c.Query("ID")
	// password := c.Query("password")
	classId := c.Query("classid")
	studentId := c.Query("studentid")

	class := Class{Id: classId}
	deleted := class.DeleteStudent(studentId)

	c.JSON(http.StatusOK, gin.H{
		"result":deleted,
	})
}

/*
func DeleteClassApi(c *gin.Context) {
	// teacherId := c.Query("ID")
	// password := c.Query("password")
	classId := c.Query("classid")

	class := Class{Id: classId}
	deleted := class.DeleteClass()

	c.JSON(http.StatusOK, gin.H{
		"result":deleted,
	})
}

 */
