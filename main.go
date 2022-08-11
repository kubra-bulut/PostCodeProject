package main

import (
	"PostCodeProject/config"
	"PostCodeProject/libs"
	"PostCodeProject/rooter"
	"fmt"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
)

func main() {
	filePath, _ := libs.GetAppDataPath("pk_list.zip")
	unzipDirectory, _ := libs.GetAppDataPath("")
	err := config.DownloadFile(filePath, "https://postakodu.ptt.gov.tr/Dosyalar/pk_list.zip")
	if err != nil {
		log.Fatal(err.Error())
	}
	err = config.GetSourceToUnzip(filePath, unzipDirectory)
	if err != nil {
		log.Fatal(err.Error())
	}

	files, err := ioutil.ReadDir(unzipDirectory)
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {
		if strings.Contains(file.Name(), "xlsx") {
			filePath, _ = libs.GetAppDataPath(file.Name())
		}
	}

	versionFilePath, _ := libs.GetAppDataPath("version.dat")
	oldVersionFilePath := ""
	if _, err := os.Stat(versionFilePath); err == nil {
		dat, _ := os.ReadFile(versionFilePath)
		oldVersionFilePath = string(dat)
		fmt.Println(oldVersionFilePath, filePath)
	}
	if oldVersionFilePath != filePath {
		fmt.Println("version changed")
		config.ImportFromExcel(filePath)
	}
	// son indirilen dosyanın adını yazıyoruz.
	err = os.WriteFile(versionFilePath, []byte(filePath), 0644)
	if err != nil {
		log.Fatal(err)
	}
	os.Remove(filePath)

	filePath, _ = libs.GetAppDataPath("pk_list.zip")
	_ = os.RemoveAll(filePath)
	config.Init()
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, "Hi from PostCode Project")
	})
	api := r.Group("")
	rooter.PostCodesRoot(api)
	fmt.Printf("http://%s:%d\n", config.LocalIPAddress, config.Port)
	_ = r.Run(fmt.Sprintf("%s:%d", config.LocalIPAddress, config.Port))

}
