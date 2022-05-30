package main

import (
	"github.com/gin-gonic/gin"
	"github.com/golang/protobuf/proto"
	"gohttpserver/model"
	"io/ioutil"
	"log"
	"os"
	"os/signal"
)

func main(){

	router := gin.Default()

	Handler(router)

	listenAddr := "0.0.0.0:8888"

	go func() {
		router.Run(listenAddr)
	}()

	log.Printf("Listening port is: %v \n", listenAddr)

	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit
	log.Println("end function")

}

func Handler(r *gin.Engine) {
	r.POST("/tencent.htm", func(c *gin.Context) {
		defer func(c  *gin.Context) {
			recover()
			c.JSON(204, "")
			return
		}(c)
		var reqTencentBody=&model.Request{}
		reqBody,_:=ioutil.ReadAll(c.Request.Body)
		proto.Unmarshal(reqBody,reqTencentBody)
		log.Printf("%+v\n",reqTencentBody)
		c.JSON(204, "")
	})
}