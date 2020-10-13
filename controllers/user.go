package controllers

import (
	"errors"
	"net/http"
	"text/template"

	"github.com/rubhiauliatirta/rublog/handler"
	"github.com/rubhiauliatirta/rublog/helpers"
	"github.com/rubhiauliatirta/rublog/models"
	"github.com/gorilla/sessions"
	"github.com/jinzhu/gorm"
)

func Login(w http.ResponseWriter, r *http.Request) {

	var tmp = template.Must(template.ParseFiles(
		"./templates/pages/auth.html",
		"./templates/components/_header.html",
		"./templates/components/_navbar.html",
		"./templates/components/_alert.html",
	))

	var data = handler.TemplateData{
		IsLogin: false,
		Title:   "Login",
		Data:    make(map[string]interface{}),
	}

	r.ParseForm()

	email := r.Form.Get("email")
	password := r.Form.Get("password")

	user, err := models.GetUser(email)

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			data.Error = errors.New("Invalid Email/Password")
			handler.RenderPage("auth", w, tmp, data)
		} else {
			data.Error = errors.New("Something error. Please try again")
			handler.RenderPage("auth", w, tmp, data)
		}
		return
	}

	if !helpers.CheckPasswordHash(password, user.Password) {
		data.Error = errors.New("Invalid Email/Password")
		handler.RenderPage("auth", w, tmp, data)
		return
	}
	session := r.Context().Value("Session").(*sessions.Session)
	session.Values["user"] = user.ID
	err = session.Save(r, w)

	if err != nil {
		data.Error = errors.New("Something error. Please try again")
		handler.RenderPage("auth", w, tmp, data)
		return
	}

	http.Redirect(w, r, "/?from=login", http.StatusSeeOther)

}

func Register(w http.ResponseWriter, r *http.Request) {

	var tmp = template.Must(template.ParseFiles(
		"./templates/pages/auth.html",
		"./templates/components/_header.html",
		"./templates/components/_navbar.html",
		"./templates/components/_alert.html",
	))

	var data = handler.TemplateData{
		IsLogin: false,
		Title:   "Register",
		Data:    make(map[string]interface{}),
	}

	r.ParseForm()
	name := r.Form.Get("name")
	email := r.Form.Get("email")
	password := r.Form.Get("password")
	confirmPassword := r.Form.Get("confirm_password")

	//tmp, _ := template.ParseFiles("./templates/auth.html")

	if password != confirmPassword {
		data.Error = errors.New("Password and Confirm Password should be same")
		data.ErrorMessage = "Password and Confirm Password should be same"
		handler.RenderPage("auth", w, tmp, data)
		return
	}

	newUser := models.User{
		Name:     name,
		Email:    email,
		Password: password,
	}

	err := newUser.IsValid()
	if err != nil {
		data.Error = err
		handler.RenderPage("auth", w, tmp, data)
		return
	}

	err = models.CreateUser(&newUser)

	if err != nil {
		data.Error = err
		handler.RenderPage("auth", w, tmp, data)
		return
	}

	//handler.RespondJSON(w, 201, newUser)
	http.Redirect(w, r, "/login?register=success", http.StatusSeeOther)
}

func Logout(w http.ResponseWriter, r *http.Request) {
	session := r.Context().Value("Session").(*sessions.Session)
	session.Values["user"] = nil
	session.Options.MaxAge = -1

	err := session.Save(r, w)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	http.Redirect(w, r, "/?from=logout", http.StatusSeeOther)
}

// func handler.RenderPage("auth", w http.ResponseWriter, tmp *template.Template, data handler.TemplateData) {
// 	err := tmp.ExecuteTemplate(w, "auth", data)
// 	if err != nil {
// 		http.Error(w, err.Error(), http.StatusInternalServerError)
// 	}
// }
