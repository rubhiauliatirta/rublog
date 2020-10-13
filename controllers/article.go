package controllers

import (
	"net/http"
	"strconv"
	"text/template"
	"time"

	"../handler"
	"../models"
)

func ArticleUpdate(w http.ResponseWriter, r *http.Request) {
	var tmp = template.Must(template.ParseFiles(
		"./templates/pages/edit.html",
		"./templates/components/_header.html",
		"./templates/components/_navbar.html",
		"./templates/components/_alert.html",
	))
	editArticle, _ := r.Context().Value("article").(models.Article)

	var data = handler.TemplateData{
		IsLogin: true,
		Title:   "Edit",
		Data:    make(map[string]interface{}),
	}
	r.ParseForm()

	editArticle.Title = r.Form.Get("title")
	editArticle.Content = r.Form.Get("content")
	editArticle.Text = r.Form.Get("text")
	editArticle.ImageURL = r.Form.Get("image_url")
	editArticle.UserID = r.Context().Value("UserID").(uint)

	isPublish := r.Form.Get("is_publish")

	if isPublish != "" {
		editArticle.IsPublish, _ = strconv.ParseBool(isPublish)
	}

	if isPublish == "true" {
		editArticle.PublishTime = time.Now()
	}

	err := models.EditArticle(&editArticle)

	if err != nil {
		data.Error = err
		handler.RenderPage("edit", w, tmp, data)
		return
	}

	http.Redirect(w, r, "/articles?from=edit", http.StatusSeeOther)

}
func ArticleCreate(w http.ResponseWriter, r *http.Request) {

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

	r.ParseForm()

	newArticle := models.Article{
		Title:    r.Form.Get("title"),
		Content:  r.Form.Get("content"),
		Text:     r.Form.Get("text"),
		ImageURL: r.Form.Get("image_url"),
		UserID:   r.Context().Value("UserID").(uint),
	}
	newArticle.IsPublish, _ = strconv.ParseBool(r.Form.Get("is_publish"))

	if newArticle.IsPublish {
		newArticle.PublishTime = time.Now()
	}

	err := models.CreateArticle(&newArticle)

	if err != nil {
		data.Error = err
		handler.RenderPage("create", w, tmp, data)
		return
	}

	http.Redirect(w, r, "/articles?from=create", http.StatusSeeOther)

}
func ArticleDelete(w http.ResponseWriter, r *http.Request) {
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

	article, _ := r.Context().Value("article").(models.Article)

	err := models.DeleteArticle(&article)

	if err != nil {
		data.Error = err
		handler.RenderPage("create", w, tmp, data)
		return
	}

	http.Redirect(w, r, "/articles?from=delete", http.StatusSeeOther)
}
func ArticleUnpublish(w http.ResponseWriter, r *http.Request) {
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

	article, _ := r.Context().Value("article").(models.Article)

	article.IsPublish = false

	err := models.EditArticle(&article)

	if err != nil {
		data.Error = err
		handler.RenderPage("article-list", w, tmp, data)
		return
	}

	http.Redirect(w, r, "/articles?from=unpublish", http.StatusSeeOther)
}
func ArticlePublish(w http.ResponseWriter, r *http.Request) {
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

	article, _ := r.Context().Value("article").(models.Article)

	article.IsPublish = true
	article.PublishTime = time.Now()

	err := models.EditArticle(&article)

	if err != nil {
		data.Error = err
		handler.RenderPage("article-list", w, tmp, data)
		return
	}

	http.Redirect(w, r, "/articles?from=publish", http.StatusSeeOther)
}

func renderErrorArticle(tmpName string, w http.ResponseWriter, tmp *template.Template, data handler.TemplateData) {
	err := tmp.ExecuteTemplate(w, tmpName, data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
