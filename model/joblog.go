package model

type JobLog struct {
	Id  uint `json:"id" form:"id" gorm:"column:id;PRIMARY_KEY"`
	//创建日期
	Date int64 `json:"date" form:"date" binding:"required" gorm:"column:date"`
	//问题类型
	ProblemType string `json:"problemtype" form:"problemtype" binding:"required" gorm:"column:problem_type"`
	//配件名称
	PartsName  string `json:"partsname" form:"partsname" gorm:"column:parts_name"`
    //问题描述
	ProblemDescription  string `json:"description" form:"description" binding:"required" gorm:"column:description"`
	//位置
	Location  string `json:"location" form:"location" binding:"required" gorm:"column:location"`
	//处理人
	UserName string `json:"username" form:"username" binding:"required" gorm:"column:username"`
	//中心名称
	CenterName string `json:"centername" form:"centername" binding:"required" gorm:"column:center_name"`
	//最后更新时间
	LastDate int64 `json:"lastdate" form:"lastdate" binding:"required" gorm:"column:last_date"`
}
func (j *JobLog)TableName() string {
	return "jobs"
}

