package apis

import (
	. "background/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func IndexApi(c *gin.Context) {
	c.String(http.StatusOK, "It works")
}

// LoginApi
// 	Path:login
// 	Input: ID, password
// 	OutPut(JSON):
//      result:boolean
// (表示账户密码是否输入正确)
//  	type:int
// (表示该用户是否为学生，1为学生，0为老师)
func LoginApi(c *gin.Context) {
	id := c.Query("ID")
	password := c.Query("password")

	p := Person{Id: id, Password: password}
	identity := p.GetPeople()

	existed := false
	if identity != -1 {
		existed = true
	}

	c.JSON(http.StatusOK, gin.H{
		"result": existed,
		"type": identity,
	})
}
