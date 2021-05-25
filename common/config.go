package common

import (
	"github.com/spf13/viper"
	"log"
	"os"
)

func init()  {
	InitConfig()
}

func InitConfig() {
	workDir,_:=os.Getwd()
	viper.SetConfigName("config")
	viper.SetConfigType("ini")
	viper.AddConfigPath(workDir+"/config")
	err:=viper.ReadInConfig()
	if err !=nil{
		log.Panicln(err)
	}
}

