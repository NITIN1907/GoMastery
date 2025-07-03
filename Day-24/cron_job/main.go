package main

import (
	"fmt"
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
	cron "github.com/robfig/cron/v3"
)

func main() {
	RunCron()
	initiateGin()
}
func RunCron() {
	c := cron.New(cron.WithSeconds())

	c.AddFunc("*/10 * * * * *", SentHello)
	c.Start()
}
func SentHello() {
	res, err := http.Get("http://localhost:8000/sent")
	if err != nil {
		fmt.Println("err:", err)
		return
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println("err from read all in cron", err)
		return
	}
	fmt.Println(string(body))
}
func initiateGin() {
	r := gin.Default()
	r.GET("/sent", SentHelloMessage)
	r.Run(":8000")
}
func SentHelloMessage(c *gin.Context) {
	fmt.Println("handled /sent route")
	c.JSON(http.StatusOK, gin.H{"message": "hello"})
}
