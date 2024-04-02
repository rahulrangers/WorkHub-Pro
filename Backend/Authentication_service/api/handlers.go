package api

import (
	"errors"
	"fmt"
	"math/rand"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/sriganeshres/WorkHub-Pro/Backend/Authentication_service/utils"
	"github.com/sriganeshres/WorkHub-Pro/Backend/models"
	"gopkg.in/gomail.v2"
)

func (app *Config) Login(ctx echo.Context) error {
	var userData models.LoginUser
	err := ctx.Bind(&userData)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, err.Error())
	}
	email := userData.Email
	password := userData.Password
	user, err := app.Db.GetUserByEmail(email)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, err.Error())
	}
	if whether, erro := utils.VerifyPassword(password, user.Password); erro == nil {
		if whether {
			return ctx.JSON(http.StatusOK, user.Username)
		} else {
			return ctx.JSON(http.StatusBadRequest, "Invalid password")
		}
	} else {
		return ctx.JSON(http.StatusUnauthorized, "Invalid credentials")
	}
}

func (app *Config) Signup(ctx echo.Context) error {

	var userData models.UserData
	err := ctx.Bind(&userData)

	if err != nil {
		return ctx.JSON(http.StatusBadRequest, err.Error())
	}
	password := userData.Password
	if password == "" {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return errors.New("no Password was sent")
	} else if len(password) <= 8 {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return errors.New("password must be at least 8 characters long")
	}
	if user, er := app.Db.GetUserByEmail(userData.Email); user != nil {
		ctx.JSON(http.StatusBadRequest, "Already the email is in use.")
		return er
	}
	userData.Password = utils.HashString(password)
	if errorer := app.Db.CreateUser(&userData); errorer != nil {
		ctx.JSON(http.StatusBadRequest, errorer)
		return errorer
	}
	ctx.JSON(http.StatusCreated, userData.Username)
	return nil
}
func (app *Config) WorkHub(ctx echo.Context) error {
	var WorkHub models.WorkHub
	err := ctx.Bind(&WorkHub)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, err.Error())
	}
	if errorer := app.Db.CreateWorkhub(&WorkHub); errorer != nil {
		ctx.JSON(http.StatusBadRequest, errorer)
		return errorer
	}
	var Code = generateRandomCode()
	ctx.JSON(http.StatusCreated, Code)

	return nil
}
func generateRandomCode() int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(9000) + 1000
}
func (app *Config) SendEmailHandler(ctx echo.Context) error {
	var requestData struct {
		Email string `json:"email"`
		Code  int    `json:"code"`
	}

	if err := ctx.Bind(&requestData); err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request"})
	}
	email := requestData.Email
	code := requestData.Code

	var err = SendEmail(email, code)
	if err != nil {
		return ctx.String(http.StatusInternalServerError, err.Error())
	}
	return ctx.String(http.StatusOK, "Email sent successfully")
}
func SendEmail(to string, code int) error {
	// Load SMTP configuration from environment variables
	smtpHost := "smtp.gmail.com"
	smtpPort := 587
	smtpUsername := "rahulreddypurmani123@gmail.com"
	smtpPassword := "tozm liuy rdcy iahd"

	if smtpHost == "" || smtpUsername == "" || smtpPassword == "" {
		return fmt.Errorf("SMTP configuration not set")
	}

	// Initialize SMTP dialer
	dialer := gomail.NewDialer(smtpHost, smtpPort, smtpUsername, smtpPassword)

	// Compose email message
	msg := gomail.NewMessage()
	msg.SetHeader("From", "rahulreddypurmani123@gmail.com")
	msg.SetHeader("To", to)
	msg.SetHeader("Subject", "Your Privacy Code")
	msg.SetBody("text/plain", fmt.Sprintf("Your privacy code is: %d", code))

	// Send email
	err := dialer.DialAndSend(msg)
	if err != nil {
		return err
	}
	fmt.Println("Email sent to", to)
	return nil
}
func (app *Config) Verify(ctx echo.Context) error {
	var requestData struct {
		Code int `json:"code"`
	}
	if err := ctx.Bind(&requestData); err != nil {
		return err
	}
	code := requestData.Code

	workhub, err := app.Db.FindWorkHub(code)
	if err != nil {
		return err
	}
	ctx.JSON(http.StatusCreated, workhub)
	return nil

}
