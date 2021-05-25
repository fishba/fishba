package util

import (

	"log"
	"math/rand"
	"time"
)

func CheckError(e error) {
	if e != nil {
		log.Println(e)
		return
	}
}

func RandomString(n int) string  {
	var letters=[]byte("qwertyuiopasdfghjklzxcvbnmqweQWERTYUIOPASDFGHJKLZXCVBNM")
	result:=make([]byte,n)
	rand.Seed(time.Now().Unix())
	for i := range result {
		result[i]=letters[rand.Intn(len(letters))]
	}
	return string(result)
}
