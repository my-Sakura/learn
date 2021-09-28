package main

import (
	"path/filepath"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.Static("/", "./pulbic")
	r.POST("/upload", func(c *gin.Context) {
		//one
		name := c.PostForm("name")
		email := c.PostForm("email")

		//two
		form, _ := c.MultipartForm()
		files := form.File["files"]
		for _, file := range files {
			filepath := filepath.Base(file.Filename)
			//three
			if err = c.SaveUploadedFile(file, filepath); err != nil {
				c.String()
			}
		}
		c.String(200, "name: %s, email: %s", name, email)
	})

	r.Run(":8080")
}
