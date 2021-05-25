package controller

import (
	"fishba.top/common"
	"fishba.top/model"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"log"
	"net/http"
	"strconv"
	"time"
)

var db=common.InitDb()
func init() {
	db.LogMode(true)
	db.AutoMigrate(&model.JobLog{},&model.User{})
}



func ShowJob(context *gin.Context){
	//获取上下文用户信息
	user ,_:= context.Get("user")
	centeruser:=user.(model.User).CenterName

	size:=context.DefaultQuery("size","10")
	page:=context.DefaultQuery("page","1")
    sizeint,_:=strconv.Atoi(size)
    pageint,_:=strconv.Atoi(page)

	var jobs []model.JobLog
	err:=db.Limit(size).Offset((pageint-1)*sizeint).Find(&jobs,"center_name LIKE ?",centeruser).Error
	if gorm.IsRecordNotFoundError(err) {
		fmt.Println("查询不到数据")
	} else if err != nil {
		//如果err不等于record not found错误，又不等于nil，那说明sql执行失败了。
		fmt.Println("查询失败", err)
	}
<<<<<<< HEAD

	Response(context,http.StatusOK,200,gin.H{"jobs":jobs,"user":model.GetDotUser(user.(model.User))},"jobs")
=======
	user ,_:= context.Get("user")
	Response(context,http.StatusOK,200,gin.H{"jobs":jobs,"user":GetUserInfo(user.(model.User))},"显示日志成功")
>>>>>>> 20210524
}

// localhost:8080/job/?size=50&index=0
func Getjob(context *gin.Context) {

	name:=context.Query("name")
	job:=model.JobLog{}
	isNotFound := db.Where("username = ?", name).First(&job).RecordNotFound()
	if isNotFound {
		context.JSON(200,gin.H{
			"msg":"查不到对应用户.",
		})
		return
	}
	context.JSON(200,gin.H{
		"msg":job,
	})
}


// localhost:8080/job/add
func AddJob(context *gin.Context) {
	job:=model.JobLog{}
	context.ShouldBind(&job)
	job.Date=time.Now().Unix()
	job.LastDate=job.Date
	if context.PostForm("partsname")==""{
		Response(context,http.StatusOK,200,gin.H{"data":job},"job")
		return
	}
	tx:=db.Begin()  //开户事务
	err:=tx.Create(&job).Error
		if err != nil {
			log.Println(err)
			tx.Rollback()
			context.JSON(200,gin.H{"msg":job})
			return
		} else {
			tx.Commit()
			context.JSON(200,gin.H{"msg":job})
		}

}

//localhost:8080/job/update/?id=5
func UpdateJob(context *gin.Context) {
	//获取参数
	job:=model.JobLog{}

	q:=context.BindQuery(&job)

	db.Model(&job).Where("id=?",job.Id).Updates(q)
	//返回数据
	context.JSON(200,gin.H{

		"msg":job,
	})



}

//localhost:8080/job/del/?id=5
func DelJob(context *gin.Context) {
	//获取参数
id:=context.Query("id")
	//数据验证
var job model.JobLog
db.Where("id=?",id).Take(&job)
db.Delete(&job)
	//返回数据


	Response(context,http.StatusOK,200,gin.H{"data":job},"job")
}