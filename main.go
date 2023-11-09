package main

import (
	"github.com/gin-gonic/gin"
	"golang.org/x/sync/errgroup"
	"log"
	"messenger/global"
)

func main() {
	authConf, err := global.GetAuthConf()
	if err != nil {
		log.Fatalln(err)
	}

	if err != nil {
		log.Fatalln(err)
	}
	appConf, err := global.GetAppConf()
	if err != nil {
		log.Fatalln(err)
	}

	gin.SetMode(gin.ReleaseMode)

	r := gin.Default()

	log.Println("start successfully...")
	if err := eg.Wait(); err != nil {
		log.Println(err)
	}
}
