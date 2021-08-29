package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"io"
	"log"
	"os"
	"fmt"
)

func test() {
	engine := gin.Default()
	engine.GET("/",func(c *gin.Context){
		c.JSON(http.StatusOK,gin.H{
			"message":"hello world",
		})
	})
	engine.Run(":3000")
}

func test2() {
	engine := gin.Default()
	//	htmlのディレクトリを指定
	engine.LoadHTMLGlob("templates/*")
	engine.GET("/",func(c *gin.Context) {
		// c.JSON(http.StatusOK,gin.H{
		// 	"message":	"hello world",
		// })
		c.HTML(http.StatusOK,"index.html",gin.H{
			//	htmlに渡す変数を定義
			"message": "ほげやな",
		})
	})
	engine.Run(":3000")
}

func test3() {
	engine := gin.Default()
	engine.LoadHTMLGlob("templates/*")
    engine.POST("/upload", func(c *gin.Context) {
        file,header, err :=  c.Request.FormFile("image")
        if err != nil {
            c.String(http.StatusBadRequest, "Bad request")
            return
        }
        fileName := header.Filename
        dir, _ := os.Getwd()
        out, err := os.Create(dir+"\\images\\"+fileName)
        if err != nil {
            log.Fatal(err)
        }
        defer out.Close()
        _, err = io.Copy(out, file)
        if err != nil {
            log.Fatal(err)
        }
        c.JSON(http.StatusOK, gin.H{
            "status": "ok",
        })
    })
    engine.Run(":3000")
}

func main() {
	engine := gin.Default()
	engine.Static("js/","templates/js/")
	engine.LoadHTMLGlob("templates/*.html")

	engine.GET("/",func(c *gin.Context) {
		c.HTML(http.StatusOK,"index.html",gin.H{
			
		})
	})

	engine.POST("/upload",func(c *gin.Context){
		file,header, err :=  c.Request.FormFile("image")
        if err != nil {
			fmt.Println(err)
            c.String(http.StatusBadRequest, "Bad request")
            return
        }

		fileName := header.Filename
		dir,_ := os.Getwd()
		fmt.Println("dir:"+dir)
		fmt.Println("fileName:"+fileName)
		fmt.Println(dir+"/images/"+fileName)
		out,err := os.Create("images/"+fileName)
		if err != nil {
			fmt.Println(err)
			log.Fatal(err)
		}
		defer out.Close()

		_,err = io.Copy(out,file)
		if err != nil {
			fmt.Println(err)
			log.Fatal(err)
		}

		c.JSON(http.StatusOK, gin.H{
			"status":"ok",
		})
	})
	engine.Run(":3000")
}