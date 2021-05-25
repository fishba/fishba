package controller

import (
	"fishba.top/model"
	"github.com/gin-gonic/gin"
	"net/http"
)

//根据request头.来返回对应格式
func Response(c *gin.Context, httpStatus int ,code int,data gin.H, msg string) {
	switch c.Request.Header.Get("Accept") {
	case "application/json":
		// Respond with JSON
		c.JSON(httpStatus, gin.H{"code":code,"data":data,"msg":msg})
	case "application/xml":
		// Respond with XML
		c.XML(httpStatus, gin.H{"code":code,"data":data,"msg":msg})
	default:
		c.JSON(httpStatus, gin.H{"code":code,"data":data,"msg":msg})
		// Respond with HTML
		//c.HTML(http.StatusOK, templateName, data)
	}
}

func Success(ctx *gin.Context,data gin.H,msg string)  {
	Response(ctx,http.StatusOK,200,data,msg)
}

func Fail(ctx *gin.Context,data gin.H,msg string)  {
	Response(ctx,http.StatusOK,400,data,msg)
}

type UserInfo struct {
	Username string `json:"username" form:"username" binding:"required" gorm:"column:username"`
	CenterName string `json:"center_name" form:"center_name" gorm:"column:center_name"`
}

func GetUserInfo(user model.User) UserInfo{
	return UserInfo{
		Username :user.Username,
		CenterName: user.CenterName,
	}
}