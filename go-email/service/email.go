package service

import (
	"os"
	"strconv"

	"github.com/emersonluiz/go-email/logger"
	"github.com/emersonluiz/go-email/models"
	gomail "gopkg.in/mail.v2"
)

func SendEmail(user *models.User) {
	mail := gomail.NewMessage()

	mail.SetHeader("From", os.Getenv("EMAIL_FROM"))
	mail.SetHeader("To", os.Getenv("EMAIL_TO"))
	mail.SetHeader("Subject", "User Register")
	mail.SetBody("text/plain", "User: "+user.Name+" and email: "+user.Email+" registered with sucess")

	port, error := strconv.Atoi(os.Getenv("EMAIL_PORT"))
	if error != nil {
		panic(error)
	}
	rtn := gomail.NewDialer(os.Getenv("EMAIL_HOST"), port, os.Getenv("EMAIL_FROM"), os.Getenv("EMAIL_PASS"))
	rtn.StartTLSPolicy = gomail.MandatoryStartTLS

	if err := rtn.DialAndSend(mail); err != nil {
		logger.SetLog(err.Error())
		panic(err)
	}

	logger.SetLog("Email Sent Successfully!")
}
