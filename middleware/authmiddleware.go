package middleware

import (
	"fishba.top/common"
	"fishba.top/model"
	"fishba.top/util"
	"github.com/gin-gonic/gin"
	"net/http"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(context *gin.Context) {
        //获取 authorizathion 头
		tokenString:=context.GetHeader("Authorization")
		if tokenString==""{
			tokenString,_=context.Cookie("sadfjwqq")
		}
		if tokenString==""{
			context.JSON(http.StatusUnauthorized,gin.H{"code":401,"msg":"权限不足"})
			context.Abort()
			return
		}
		token,claims,err:=util.ParseToken(tokenString)
		if err !=nil || !token.Valid{
			context.JSON(http.StatusUnauthorized,gin.H{"code":401,"msg":"权限不足"})
			context.Abort()
			return
		}
		userId:=claims.UserId
		db:=common.InitDb()
		var user model.User
		db.Where("id = ?", userId).First(&user)
		if user.Id==0{
			context.JSON(http.StatusUnauthorized,gin.H{"code":401,"msg":"权限不足"})
			context.Abort()
			return
		}

		context.Set("user",user)
		context.Next()
	}
}