package apis

import (
	. "background/models"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"time"
)

// UploadHomeworkApi
//	path: uploadhomework
//  Input: ID, password, classid: String, homeworkid: String, file, file_name:string
//  Output:
// 返回状态包含在code中，如未成功上传返回200之外的值
func UploadHomeworkApi(c *gin.Context) {
	id := c.PostForm("ID")
	// password := c.PostForm("password")
	classId := c.PostForm("classid")
	hwId := c.PostForm("homeworkid")
	filename := c.PostForm("file_name")

	// s := Student{Id: id, Password: password}
	hwInfo := HomeworkInfo{ClassId: classId, Pid: hwId, Name: filename}
	hw := Homework{Info: hwInfo, StudentId: id, Filename: filename}

	/*
	if hwInfo.ExistedHomeworkInfo() {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": 400,
		})
		return
	}

	 */

	file, err := c.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadGateway, gin.H{
			"code": 401,
		})
		return
	}

	basePath := "/user/upload/"
	fileName := basePath + filepath.Base(file.Filename)
	if err := c.SaveUploadedFile(file, fileName); err != nil {
		e := fmt.Sprintln(err)
		c.JSON(http.StatusRequestTimeout, gin.H{
			"code": 402,
			"err":e,
		})
		return
	}

	hw.URL = fileName
	uploaded := hw.UploadHomework()

	c.JSON(http.StatusOK, gin.H{
		"code": uploaded,
	})
}

// CreateHomeworkInfoApi
//+ path：createhomewrok
//    + 参数ID，password同前，homeworkname：string（作业名字）,ddl:String(ddl，格式为2000/01/01/00:00), classid。
//    + 返回JSON，属性result：int，1成功，0不成功。
func CreateHomeworkInfoApi(c *gin.Context) {
	//teacherId := c.Query("ID")
	//password := c.Query("password")
	classId := c.Query("classid")
	pubTime := time.Now().Format("2006/01/02/03:04")
	hwName := c.Query("homeworkname")
	ddl := c.Query("ddl")


	info := HomeworkInfo{ClassId: classId, StartTime: pubTime, Name: hwName, Deadline: ddl}
	created, err := info.CreateHomeworkInfo()

	c.JSON(http.StatusOK, gin.H{
		"result": created,
		"error":err,
		"start":info.StartTime,
		"ddl":info.Deadline,
	})
}

// DownloadHomeworkApi
//path:downloadhomework
//ID，password同上，studentid:string(要下载的作业对应的学生id，前面的id为老师id)，classid（课程id），homeworkid（作业id）。
//返回的header中包含一个属性file_name表示文件名字，文件内容包含在响应体中/*
func DownloadHomeworkApi(c *gin.Context) {
	//teacherId := c.Query("ID")
	//password := c.Query("password")
	studentId := c.Query("studentid")
	classId := c.Query("classid")
	pid := c.Query("homeworkid")

	info := HomeworkInfo{ClassId: classId, Pid: pid}
	hw := Homework{Info: info, StudentId: studentId}
	fileURL, fileName, err := hw.DownloadHomework()
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"error1":err,
		})
		return
	}

	file, _ := os.Open(fileURL)
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			c.JSON(http.StatusOK, gin.H{
				"error2":err,
			})
		}
	}(file)

	buf := make([]byte, 10 * 1024 * 1024)
	_, err = file.Read(buf)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"error3":err,
		})
		return
	}

	att := fmt.Sprintf("attachment; filename=%s", fileName)
	length := 10 * 1024 * 1024
	c.Writer.WriteHeader(http.StatusOK)
	c.Header("Content-Disposition", att)
	c.Header("Content-type", "application/octet-stream")
	c.Header("filename", fileName)
	c.Header("Accept-Length", fmt.Sprintf("%d", length))
	_, err = c.Writer.Write(buf)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"error4":err,
		})
		return
	}

/*
	c.JSON(http.StatusOK, gin.H{
		"file_name": fileName,
		"file_url":fileURL,
	})

 */
}


// DownloadAllHomeworkApi
//+ path:downloadallhomework
//    + ID,password同上,classid:string课程id，homeworkid：string作业id
//    + 响应JSON array中包含一个属性file_name, file_url/*
func DownloadAllHomeworkApi(c *gin.Context) {
	//id := c.Query("ID")
	//password := c.Query("password")
	classId := c.Query("classid")
	pid := c.Query("homeworkid")

	info := HomeworkInfo{ClassId: classId, Pid: pid}
	hwList := info.DownloadAllHomework()

	c.IndentedJSON(200, hwList)
}

// GetHomeworkListApi
//+ path:gethomeworkany
//    + 参数ID，password同前，classid(课程id)，homeworkid（作业id）
//    + 返回一个array，每一项为一个JSON属性为，name：string（学生名字）number：String（学号，学生id）grade：String（作业分数，如果还没打的话返回空字符串），status：String，作业是否提交。/*
func GetHomeworkListApi(c *gin.Context) {
	// teacherId := c.Query("ID")
	// password := c.Query("password")
	classId := c.Query("classid")
	pid := c.Query("homeworkid")

	info := HomeworkInfo{ClassId: classId, Pid: pid}
	hwList := info.GetHomeworkList()

	c.IndentedJSON(200, hwList)
}

// AddScoreApi
//+ path：mark，
//    + ID，password同上，studentid：string学生id（学号），classid：string课程id，homeworkid：作业id，score：float分数，
//    + 返回object，属性result：int，1成功，0不成功。/*
func AddScoreApi(c *gin.Context) {
	// teacherId := c.Query("ID")
	// password := c.Query("password")
	studentId := c.Query("studentid")
	classId := c.Query("classid")
	pid := c.Query("homeworkid")
	score, _ := strconv.ParseFloat(c.Query("score"), 64)

	info := HomeworkInfo{ClassId: classId, Pid: pid}
	hw := Homework{Info: info, Score: score, StudentId: studentId}
	added := hw.AddScore()

	c.JSON(http.StatusOK, gin.H{
		"result": added,
	})
}
