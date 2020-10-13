package controllers

import (
	"fmt"
	"log"
	"net/http"
	"net/smtp"
	"os"
	"text/template"

	"../handler"
)

// smtpServer data to smtp server
type smtpServer struct {
	host string
	port string
} // Address URI to smtp server
func (s *smtpServer) Address() string {
	return s.host + ":" + s.port
}

func SendEmail(w http.ResponseWriter, r *http.Request) {

	var tmp = template.Must(template.ParseFiles(
		"./templates/pages/contact.html",
		"./templates/components/_header.html",
		"./templates/components/_navbar.html",
		"./templates/components/_alert.html",
	))
	userID := r.Context().Value("UserID")
	isLogin := userID != uint(0) && userID != nil

	var data = handler.TemplateData{
		IsLogin: isLogin,
		Title:   "Contact",
		Data:    make(map[string]interface{}),
	}

	r.ParseForm()
	name := r.Form.Get("name")
	email := r.Form.Get("email")
	content := r.Form.Get("content")

	username := os.Getenv("MAILTRAP_USERNAME")
	password := os.Getenv("MAILTRAP_PASSWORD")
	auth := smtp.PlainAuth("", username, password, "smtp.mailtrap.io")

	to := []string{"mail@rublog.com"}
	msg := []byte("From: " + email + " \r\n" +
		"To: mail@rublog.com\r\n" +
		"Subject: Message From RuBlog User!\r\n" +
		"\r\nHi, this is message from " + name + "\r\n" +
		content + "\r\n")
	err := smtp.SendMail("smtp.mailtrap.io:25", auth, email, to, msg)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Email Sent!")

	data.Data["Message"] = "Thankyou for reaching us, we will reply your message directly to " + email
	handler.RenderPage("contact", w, tmp, data)

}
