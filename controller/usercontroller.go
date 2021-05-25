package controller

import (
	"fishba.top/model"
	"fishba.top/util"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"log"
	"net/http"
	"time"
)


func Register(context *gin.Context) {
	//获取参数
	user:=model.User{}
	context.ShouldBind(&user)

	//数据验证
	if user.Username==""{
		Fail(context,gin.H{},"请输入用户名.")
		return
	}

	if user.Password==""{
		Fail(context,gin.H{},"请输入密码.")
		return
	}
	if user.CenterName==""{
		Fail(context,gin.H{},"请选择中心名称.")
		return
	}
	user.CreateTime=time.Now().Unix()
	user.LastTime=user.CreateTime

	var err error
	hasedPassword,err:=bcrypt.GenerateFromPassword([]byte(user.Password),bcrypt.DefaultCost)
	user.Password=string(hasedPassword)
	if err!=nil{
		Fail(context,gin.H{},"加密错误.")
		return
	}


	//写入数据库
	tx:=db.Begin()  //开户事务
	err=tx.Create(&user).Error
	if err != nil {
		log.Println(err)
		tx.Rollback()
		Fail(context,gin.H{},"新建用户失败.")
		return
	} else {
		tx.Commit()
		db.First(&user,user.Username)
		token,err:=util.CreateToken(user)
		if err !=nil{
			log.Println(err)
		}
		context.Header("Authorization",token)
		Response(context,http.StatusOK,200,gin.H{"user":user,"token":token},"用户注册成功")
	}
}

func Login(context *gin.Context) {
	//获取参数
	var  user model.User
	username:=context.PostForm("user")
	password:=context.PostForm("password")
	//数据验证
	db.Where("username = ?",username).First(&user)
	if user.Id==0{
		Response(context,http.StatusUnprocessableEntity,422,gin.H{},"用户不存在")
		return
	}
	//判断密码
	if err:=bcrypt.CompareHashAndPassword([]byte(user.Password),[]byte(password));err!=nil{
		Response(context,http.StatusBadRequest,400,gin.H{},"密码不正确")
		return
	}

	//发放token
	token,err:=util.CreateToken(user)
	if err !=nil{
		Response(context,http.StatusInternalServerError,500,gin.H{},"系统异常")
		log.Println(err)
		return
	}

	context.SetCookie("sadfjwqq",token,0,"","",false,true)
	//返回结果
	Success(context,gin.H{"token":token},"登入成功")
}

func Info(context *gin.Context)  {
	user ,_:= context.Get("user")

	Response(context,http.StatusOK,200,gin.H{"user":GetUserInfo(user.(model.User))},"用户信息")
}