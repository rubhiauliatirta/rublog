package main

import (
	"flag"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/mux"
	"github.com/rubhiauliatirta/rublog/config"
	"github.com/rubhiauliatirta/rublog/controllers"
	"github.com/rubhiauliatirta/rublog/middlewares"
	"github.com/rubhiauliatirta/rublog/models"
)

var PORT string

func init() {
	value := os.Getenv("PORT")
	if value == "" {
		PORT = "9090"
	} else {
		PORT = value
	}
}

func main() {

	router := setRouter()

	config.Db.AutoMigrate(&models.User{})
	config.Db.AutoMigrate(&models.Article{})

	server := &http.Server{
		Handler: middlewares.SessionContext(router),
		Addr:    "127.0.0.1:" + PORT,

		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Fatal(server.ListenAndServe())
}

func setRouter() *mux.Router {
	var dir string
	flag.StringVar(&dir, "dir", "./static", "the directory to serve files from. Defaults to the current dir")
	flag.Parse()

	router := mux.NewRouter()
	router.PathPrefix("/static").Handler(http.StripPrefix("/static", http.FileServer(http.Dir(dir))))

	router.HandleFunc("/login", controllers.LoginPage).Methods("GET")
	router.HandleFunc("/login", controllers.Login).Methods("POST")
	router.HandleFunc("/logout", controllers.Logout).Methods("GET")

	router.HandleFunc("/register", controllers.RegisterPage).Methods("GET")
	router.HandleFunc("/register", controllers.Register).Methods("POST")

	router.Handle("/articles/delete/{id}", middlewares.StrictAuthentication(middlewares.ArticleAuthorization(http.HandlerFunc(controllers.ArticleDelete)))).Methods("GET")

	router.Handle("/articles/edit/{id}", middlewares.StrictAuthentication(middlewares.ArticleAuthorization(http.HandlerFunc(controllers.ArticleEditPage)))).Methods("GET")
	router.Handle("/articles/edit/{id}", middlewares.StrictAuthentication(middlewares.ArticleAuthorization(http.HandlerFunc(controllers.ArticleUpdate)))).Methods("POST")

	router.Handle("/articles/publish/{id}", middlewares.StrictAuthentication(middlewares.ArticleAuthorization(http.HandlerFunc(controllers.ArticlePublish)))).Methods("GET")
	router.Handle("/articles/unpublish/{id}", middlewares.StrictAuthentication(middlewares.ArticleAuthorization(http.HandlerFunc(controllers.ArticleUnpublish)))).Methods("GET")

	router.Handle("/articles/create", middlewares.StrictAuthentication(http.HandlerFunc(controllers.ArticleCreatePage))).Methods("GET")
	router.Handle("/articles/create", middlewares.StrictAuthentication(http.HandlerFunc(controllers.ArticleCreate))).Methods("POST")

	router.Handle("/articles/detail/{id}", middlewares.DetailAuthorization(http.HandlerFunc(controllers.ArticleViewPage))).Methods("GET")
	router.Handle("/articles", middlewares.StrictAuthentication(http.HandlerFunc(controllers.ArticleListPage))).Methods("GET")

	router.Handle("/about", middlewares.Authentication(http.HandlerFunc(controllers.AboutPage))).Methods("GET")
	router.Handle("/contact", middlewares.Authentication(http.HandlerFunc(controllers.SendEmail))).Methods("POST")
	router.Handle("/contact", middlewares.Authentication(http.HandlerFunc(controllers.ContactPage))).Methods("GET")
	router.Handle("/", middlewares.Authentication(http.HandlerFunc(controllers.HomePage))).Methods("GET")

	return router
}
