package model

type User struct {
	Id  uint `json:"id" form:"id" gorm:"column:id;PRIMARY_KEY"`
<<<<<<< HEAD
	//用户名
	Username string `json:"user" form:"user" binding:"required" gorm:"column:username"`
	//密码
=======
	Username string `json:"username" form:"username" binding:"required" gorm:"column:username"`
>>>>>>> 20210524
	Password string `json:"password" form:"password" binding:"required" gorm:"column:password"`
	//中心名称
	CenterName string `json:"center_name" form:"center_name" gorm:"column:center_name"`
	//创建时间
	CreateTime int64 `gorm:"column:create_time"`
	//最后登入时间
	LastTime int64  `gorm:"column:Last_time"`
}

func (u *User)TableName() string {
	return "users"
}

type DotUser struct {
	//用户名
	Username string `json:"user" form:"user" binding:"required"`
	//中心名称
	CenterName string `json:"center_name" form:"center_name"`
}

func GetDotUser(user User) DotUser {
	return DotUser{
		Username: user.Username,
		CenterName: user.CenterName,
	}
}