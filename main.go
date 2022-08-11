package main

import (
	"PostCodeProject/config"
	"PostCodeProject/rooter"
	"github.com/gin-gonic/gin"
)

func main() {
	config.Init()
	r := gin.Default()
	api := r.Group("")
	rooter.PostCodesRoot(api)
	r.Run()
}
