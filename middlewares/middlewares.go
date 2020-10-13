package middlewares

import (
	"context"
	"net/http"

	"github.com/rubhiauliatirta/rublog/config"
	"github.com/rubhiauliatirta/rublog/models"
	"github.com/gorilla/mux"
	"github.com/gorilla/sessions"
)

func Authentication(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		session := r.Context().Value("Session").(*sessions.Session)
		userID, ok := session.Values["user"].(uint)

		if ok {
			ctx := context.WithValue(r.Context(), "UserID", userID)
			next.ServeHTTP(w, r.WithContext(ctx))

		} else {
			next.ServeHTTP(w, r)
		}

	})
}

func DetailAuthorization(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		session := r.Context().Value("Session").(*sessions.Session)
		userID, _ := session.Values["user"].(uint)

		isLogin := userID != uint(0)

		articleID, _ := mux.Vars(r)["id"]

		article, err := models.GetArticleById(articleID)
		// fmt.Println("isPublished: ", article.IsPublish)
		// fmt.Println("UserID: ", userID)
		// fmt.Println("article.UserID: ", article.UserID)

		if err != nil {
			http.Redirect(w, r, "/?from=404", http.StatusSeeOther)
		} else if article.IsPublish || article.UserID == userID {
			ctx := context.WithValue(r.Context(), "isLogin", isLogin)
			ctx = context.WithValue(ctx, "article", article)
			next.ServeHTTP(w, r.WithContext(ctx))

		} else {
			http.Redirect(w, r, "/?from=403", http.StatusSeeOther)
		}

	})
}

func StrictAuthentication(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		session := r.Context().Value("Session").(*sessions.Session)
		userID, _ := session.Values["user"].(uint)

		isLogin := userID != uint(0)

		if isLogin {
			ctx := context.WithValue(r.Context(), "UserID", userID)
			next.ServeHTTP(w, r.WithContext(ctx))

		} else {
			http.Redirect(w, r, "/?from=forbidden", http.StatusSeeOther)
		}

	})
}

func ArticleAuthorization(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		userID := r.Context().Value("UserID").(uint)

		articleID, _ := mux.Vars(r)["id"]

		article, err := models.GetArticleById(articleID)
		// fmt.Println("isPublished: ", article.IsPublish)
		// fmt.Println("UserID: ", userID)
		// fmt.Println("article.UserID: ", article.UserID)

		if err != nil {
			http.Redirect(w, r, "/?from=404", http.StatusSeeOther)
		} else if article.UserID == userID {
			ctx := context.WithValue(r.Context(), "article", article)
			next.ServeHTTP(w, r.WithContext(ctx))

		} else {
			http.Redirect(w, r, "/?from=403", http.StatusSeeOther)
		}

	})
}
func SessionContext(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		session, err := config.Store.Get(r, "auth")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		ctx := context.WithValue(r.Context(), "Session", session)
		next.ServeHTTP(w, r.WithContext(ctx))

	})
}
