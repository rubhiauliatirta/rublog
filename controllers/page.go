package controllers

import (
	"errors"
	"net/http"
	"text/template"

	"../handler"
	"../models"
)

func LoginPage(w http.ResponseWriter, r *http.Request) {
	var data = handler.TemplateData{
		IsLogin: false,
		Title:   "Login",
		Data:    make(map[string]interface{}),
	}

	var tmp = template.Must(template.ParseFiles(
		"./templates/pages/auth.html",
		"./templates/components/_header.html",
		"./templates/components/_navbar.html",
		"./templates/components/_alert.html",
	))

	registerSuccess := r.URL.Query().Get("register")

	if registerSuccess == "success" {
		data.Data["Message"] = "Register Success"
	}

	handler.RenderPage("auth", w, tmp, data)
}

func RegisterPage(w http.ResponseWriter, r *http.Request) {
	var data = handler.TemplateData{
		IsLogin: false,
		Title:   "Register",
		Data:    make(map[string]interface{}),
	}

	var tmp = template.Must(template.ParseFiles(
		"./templates/pages/auth.html",
		"./templates/components/_header.html",
		"./templates/components/_navbar.html",
		"./templates/components/_alert.html",
	))
	handler.RenderPage("auth", w, tmp, data)
}

func HomePage(w http.ResponseWriter, r *http.Request) {
	var tmp = template.Must(template.ParseFiles(
		"./templates/pages/home.html",
		"./templates/components/_header.html",
		"./templates/components/_navbar.html",
		"./templates/components/_alert.html",
	))
	userID := r.Context().Value("UserID")
	isLogin := userID != uint(0) && userID != nil

	var data = handler.TemplateData{
		IsLogin: isLogin,
		Title:   "Home",
		Data:    make(map[string]interface{}),
	}

	from := r.URL.Query().Get("from")

	if from == "login" {
		data.Data["Message"] = "Welcome to RuBlog!"
	} else if from == "logout" {
		data.Data["Message"] = "Bye, see you again!"
	} else if from == "403" {
		data.Error = errors.New("Forbidden Access! Please login first")
	} else if from == "404" {
		data.Error = errors.New("Article not found!")
	}

	articles, err := models.GetPublishedArticles()

	if err != nil {
		panic(err)
	}

	data.Data["Articles"] = articles
	handler.RenderPage("home", w, tmp, data)

}

func AboutPage(w http.ResponseWriter, r *http.Request) {
	var tmp = template.Must(template.ParseFiles(
		"./templates/pages/about.html",
		"./templates/components/_header.html",
		"./templates/components/_navbar.html",
	))
	userID := r.Context().Value("UserID")
	isLogin := userID != uint(0) && userID != nil

	var data = handler.TemplateData{
		IsLogin: isLogin,
		Title:   "About",
	}

	handler.RenderPage("about", w, tmp, data)

}

func ContactPage(w http.ResponseWriter, r *http.Request) {
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
	}

	handler.RenderPage("contact", w, tmp, data)
}

func ArticleEditPage(w http.ResponseWriter, r *http.Request) {
	var tmp = template.Must(template.ParseFiles(
		"./templates/pages/edit.html",
		"./templates/components/_header.html",
		"./templates/components/_navbar.html",
		"./templates/components/_alert.html",
	))
	article, _ := r.Context().Value("article").(models.Article)

	var data = handler.TemplateData{
		IsLogin: true,
		Title:   "Edit",
		Data:    make(map[string]interface{}),
	}

	data.Data["Article"] = article
	handler.RenderPage("edit", w, tmp, data)
}
func ArticleCreatePage(w http.ResponseWriter, r *http.Request) {
	var tmp = template.Must(template.ParseFiles(
		"./templates/pages/create.html",
		"./templates/components/_header.html",
		"./templates/components/_navbar.html",
		"./templates/components/_alert.html",
	))

	var data = handler.TemplateData{
		IsLogin: true,
		Title:   "Create",
		Data:    make(map[string]interface{}),
	}

	handler.RenderPage("create", w, tmp, data)

}
func ArticleViewPage(w http.ResponseWriter, r *http.Request) {
	var tmp = template.Must(template.ParseFiles(
		"./templates/pages/detail.html",
		"./templates/components/_header.html",
		"./templates/components/_navbar.html",
	))

	isLogin := r.Context().Value("isLogin").(bool)
	article := r.Context().Value("article").(models.Article)

	//	var tmpContent = template.Must(template.New("article-content").Parse(`{{define "content"}}` + article.Content + `{{end}}`))

	template.Must(tmp.New("").Parse(`{{define "content"}}` + article.Content + `{{end}}`))

	var data = handler.TemplateData{
		IsLogin: isLogin,
		Title:   article.Title,
		Data:    make(map[string]interface{}),
	}

	data.Data["Article"] = article
	data.Data["PublishDate"] = article.FormattedDate()
	handler.RenderPage("detail", w, tmp, data)

}

func ArticleListPage(w http.ResponseWriter, r *http.Request) {
	var tmp = template.Must(template.ParseFiles(
		"./templates/pages/article-list.html",
		"./templates/components/_header.html",
		"./templates/components/_navbar.html",
		"./templates/components/_alert.html",
	))

	var data = handler.TemplateData{
		IsLogin: true,
		Title:   "My Article",
		Data:    make(map[string]interface{}),
	}

	from := r.URL.Query().Get("from")

	if from == "create" {
		data.Data["Message"] = "Create article success"
	} else if from == "edit" {
		data.Data["Message"] = "Edit article success"
	} else if from == "delete" {
		data.Data["Message"] = "Delete article success"
	} else if from == "publish" {
		data.Data["Message"] = "Publish article success"
	} else if from == "unpublish" {
		data.Data["Message"] = "Unpublish article success"
	}

	userID := r.Context().Value("UserID").(uint)

	articles, err := models.GetUserArticles(userID)
	if err != nil {

		panic(err)
	}
	data.Data["Articles"] = articles
	handler.RenderPage("article-list", w, tmp, data)

}
