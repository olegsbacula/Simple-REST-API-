// routes.go
package router

import (
    "net/http"
    "github.com/go-chi/chi"
    "yourapp/handler"
)

func UserRouter(h *handler.UserHandler) http.Handler {
    r := chi.NewRouter()
    r.Get("/users", h.ListUsers)
    r.Post("/users", h.CreateUser)
    r.Get("/users/{id}", h.GetUser)
    r.Put("/users/{id}", h.UpdateUser)
    r.Delete("/users/{id}", h.DeleteUser)
    return r
}

func ArticleRouter(h *handler.ArticleHandler) http.Handler {
    a := chi.NewRouter()
    a.Get("/article", h.ListArticles)
    a.Post("/article", h.CreateArticle)
    a.Get("/article/{id}", h.GetArticle)
    a.Put("/article/{id}", h.UpdateArticle)
    a.Delete("/article/{id}", h.DeleteArticle)
    return a
}

func JournalistRouter(h *handler.JournalistHandler) http.Handler {
    j := chi.NewRouter()
    j.Get("/journalist", h.ListJournalist)
    j.Post("/article", h.CreateJournalist)
    j.Get("/article/{id}", h.GetJournalist)
    j.Put("/article/{id}", h.UpdateJournalist)
    j.Delete("/article/{id}", h.DeleteJournalist)
    return j
}