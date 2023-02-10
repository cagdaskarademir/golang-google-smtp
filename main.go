package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"log"
	"net/smtp"
	"os"
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Printf("No .env file found")
	}
}

// this is an application that you can send a email through
func main() {

	r := gin.Default()
	r.GET("/send", func(c *gin.Context) {
		from := os.Getenv("SMTP_FROM")
		password := os.Getenv("SMTP_PASSWORD")

		to := []string{
			os.Getenv("SMTP_TO"),
		}

		host := os.Getenv("SMTP_HOST")
		port := os.Getenv("SMTP_PORT")

		// Message.
		message := []byte("This is a test email message.")

		auth := smtp.PlainAuth("", from, password, host)

		err := smtp.SendMail(host+":"+port, auth, from, to, message)
		if err != nil {
			fmt.Println(err)
			c.JSON(500, gin.H{
				"message": err.Error(),
			})
		}

		c.JSON(200, gin.H{
			"message": "Email Delivered Successfully!",
		})
	})

	r.Run()
}
